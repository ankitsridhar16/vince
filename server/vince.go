package server

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/caddyserver/certmagic"
	"github.com/vinceanalytics/vince/alerts"
	"github.com/vinceanalytics/vince/assets"
	"github.com/vinceanalytics/vince/caches"
	"github.com/vinceanalytics/vince/config"
	"github.com/vinceanalytics/vince/email"
	"github.com/vinceanalytics/vince/health"
	"github.com/vinceanalytics/vince/models"
	"github.com/vinceanalytics/vince/pkg/group"
	"github.com/vinceanalytics/vince/pkg/log"
	"github.com/vinceanalytics/vince/plug"
	"github.com/vinceanalytics/vince/router"
	"github.com/vinceanalytics/vince/sessions"
	"github.com/vinceanalytics/vince/timeseries"
	"github.com/vinceanalytics/vince/userid"
	"github.com/vinceanalytics/vince/worker"
	"golang.org/x/sync/errgroup"
)

func Serve(o *config.Options) error {
	ctx, err := config.Load(o)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	return HTTP(ctx, o)
}

type resourceList []io.Closer

type resourceFunc func() error

func (r resourceFunc) Close() error {
	return r()
}

func (r resourceList) Close() error {
	e := make([]error, len(r))
	for i, f := range r {
		e[i] = f.Close()
	}
	return errors.Join(e...)
}

func HTTP(ctx context.Context, o *config.Options) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	var resources resourceList

	// we start listeners early to make sure we can actually bind to the network.
	// This saves us managing all long running goroutines we start in this process.
	httpListener, err := net.Listen("tcp", o.Listen)
	if err != nil {
		return fmt.Errorf("failed to bind to a network address %v", err)
	}
	resources = append(resources, httpListener)
	var httpsListener net.Listener
	var magic *certmagic.Config
	if o.TLS.Enabled {
		if o.TLS.Address == "" {
			resources.Close()
			return errors.New("tls-address is required")
		}
		if o.TLS.Key == "" || o.TLS.Cert == "" {
			if !o.Acme.Enabled {
				resources.Close()
				return errors.New("tls-key and tls-cert  are required")
			}
		}
		if o.Acme.Enabled {
			if o.Acme.Email == "" || o.Acme.Domain == "" {
				resources.Close()
				return errors.New("acme-email and acme-domain  are required")
			}
			magic = certmagic.NewDefault()
			// we use file storage for certs
			certsPath := filepath.Join(o.DataPath, "certs")
			os.MkdirAll(certsPath, 0755)
			magic.Storage = &certmagic.FileStorage{Path: certsPath}
			myACME := certmagic.NewACMEIssuer(magic, certmagic.ACMEIssuer{
				CA:     certmagic.LetsEncryptStagingCA,
				Email:  o.Acme.Email,
				Agreed: true,
			})
			magic.Issuers = append(magic.Issuers, myACME)
			err = magic.ManageSync(ctx, []string{o.Acme.Domain})
			if err != nil {
				resources.Close()
				return fmt.Errorf("failed to sync acme domain %v", err)
			}
			httpsListener, err = net.Listen("tcp", o.TLS.Address)
			if err != nil {
				resources.Close()
				return fmt.Errorf("failed to bind to https socket %v", err)
			}
		} else {
			cert, err := tls.LoadX509KeyPair(o.TLS.Cert, o.TLS.Key)
			if err != nil {
				resources.Close()
				return fmt.Errorf("failed to load https certificate %v", err)
			}
			config := tls.Config{}
			config.Certificates = append(config.Certificates, cert)
			httpsListener, err = tls.Listen("tcp", o.TLS.Address, &config)
			if err != nil {
				resources.Close()
				return fmt.Errorf("failed to bind https socket %v", err)
			}
			resources = append(resources, httpsListener)
		}
	}

	var g errgroup.Group
	ctx = group.Set(ctx, &g)

	ctx = userid.Open(ctx)

	sqlDb, err := models.Open(models.Database(o))
	if err != nil {
		resources.Close()
		return err
	}
	resources = append(resources, resourceFunc(func() error {
		return models.CloseDB(sqlDb)
	}))

	ctx = models.Set(ctx, sqlDb)

	if o.Bootstrap.Enabled {
		log.Get().Debug().Msg("bootstrapping user")
		if o.Bootstrap.Name == "" ||
			o.Bootstrap.Email == "" ||
			o.Bootstrap.Password == "" ||
			o.Bootstrap.Key == "" {
			return errors.New("bootstrap-name, bootstrap-email, bootstrap-password, and bootstrap-key, are required")
		}
		models.Bootstrap(ctx,
			o.Bootstrap.Name, o.Bootstrap.Email, o.Bootstrap.Password, o.Bootstrap.Key,
		)
	}
	if o.Alerts.Enabled {
		log.Get().Debug().Msg("setup alerts")
		ctx = alerts.Setup(ctx, o)
	}
	if o.EnableEmail {
		log.Get().Debug().Msg("setup mailer")
		mailer, err := email.FromConfig(o)
		if err != nil {
			log.Get().Err(err).Msg("failed creating mailer")
			resources.Close()
			return err
		}
		resources = append(resources, mailer)
		ctx = email.Set(ctx, mailer)
	}
	ctx, ts, err := timeseries.Open(ctx, o)
	if err != nil {
		resources.Close()
		return err
	}
	resources = append(resources, ts)
	ctx, err = caches.Open(ctx)
	if err != nil {
		log.Get().Err(err).Msg("failed to open caches")
		resources.Close()
		return err
	}
	resources = append(resources, resourceFunc(func() error {
		return caches.Close(ctx)
	}))

	session := sessions.NewSession("_vince")
	ctx = sessions.Set(ctx, session)
	var h health.Health
	addHealth := func(x *health.Ping) {
		h = append(h, x)
	}
	h = append(h, health.Base{
		Key:       "database",
		CheckFunc: models.Check,
	})
	{
		// register and start workers
		g.Go(worker.UpdateCacheSites(ctx, addHealth))
		g.Go(worker.SaveTimeseries(ctx, addHealth))
	}

	resources = append(resources, h)
	ctx = health.Set(ctx, h)

	svr := buildServer(ctx, &g, httpListener, httpsListener, magic, Handle(ctx))
	// We start by shutting down the server before shutting everything else. So we
	// prepend svr for it to be called first.
	resources = append(resourceList{svr}, resources...)

	g.Go(func() error {
		// Ensure we close the server.
		<-ctx.Done()
		log.Get().Debug().Msg("shutting down gracefully ")
		return resources.Close()
	})
	log.Get().Debug().Str("address", httpListener.Addr().String()).Msg("started serving  http traffic")
	if httpsListener != nil {
		log.Get().Debug().Str("address", httpsListener.Addr().String()).Msg("started serving  https traffic")
	}
	g.Go(func() error {
		abort := make(chan os.Signal, 1)
		signal.Notify(abort, os.Interrupt)
		sig := <-abort
		log.Get().Info().Msgf("received signal %s shutting down the server", sig)
		cancel()
		return nil
	})
	return g.Wait()
}

func buildServer(
	ctx context.Context,
	g *errgroup.Group,
	httpListener, httpsListener net.Listener,
	magic *certmagic.Config,
	h http.Handler,
) (r resourceList) {
	httpSvr := &http.Server{
		Handler:           h,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}
	if httpsListener != nil {
		httpSvr.Handler = redirect(httpsListener.Addr().String())
		if magic != nil {
			// We are using tls with auto tls
			httpSvr.Handler = magic.Issuers[0].(*certmagic.ACMEIssuer).HTTPChallengeHandler(
				redirect(httpsListener.Addr().String()),
			)
		}
	}
	g.Go(func() error {
		return httpSvr.Serve(httpListener)
	})
	r = append(r, httpSvr)
	if httpsListener != nil {
		httpsSvr := &http.Server{
			Handler:           h,
			ReadHeaderTimeout: 10 * time.Second,
			ReadTimeout:       30 * time.Second,
			WriteTimeout:      2 * time.Minute,
			IdleTimeout:       5 * time.Minute,
			BaseContext: func(l net.Listener) context.Context {
				return ctx
			},
		}
		if magic != nil {
			// httpsListener is not wrapped with tls yet. We use certmagic to obtain
			// tls Config and properly wrap it.
			tlsConfig := magic.TLSConfig()
			tlsConfig.NextProtos = append([]string{"h2", "http/1.1"}, tlsConfig.NextProtos...)
			httpsListener = tls.NewListener(httpsListener, tlsConfig)
		}
		g.Go(func() error {
			return httpsSvr.Serve(httpsListener)
		})
		r = append(r, httpsSvr)
	}
	return
}

func Handle(ctx context.Context) http.Handler {
	pipe := append(
		plug.Pipeline{
			plug.Track(),
			plug.Favicon(plug.DefaultClient),
			assets.Plug(),
			plug.RequestID,
			plug.CORS,
		},
		router.Pipe(ctx)...,
	)
	h := pipe.Pass(plug.NOOP)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}

func redirect(addr string) http.Handler {
	_, port, _ := net.SplitHostPort(addr)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		toURL := "https://"
		requestHost := hostOnly(r.Host)
		toURL += requestHost + ":" + port
		toURL += r.URL.RequestURI()
		w.Header().Set("Connection", "close")
		http.Redirect(w, r, toURL, http.StatusMovedPermanently)
	})
}

// hostOnly returns only the host portion of hostport.
// If there is no port or if there is an error splitting
// the port off, the whole input string is returned.
func hostOnly(hostport string) string {
	host, _, err := net.SplitHostPort(hostport)
	if err != nil {
		return hostport // OK; probably had no port to begin with
	}
	return host
}
