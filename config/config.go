package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v3"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

//go:generate protoc -I=. --go_out=paths=source_relative:. config.proto

type configKey struct{}

func Get(ctx context.Context) *Config {
	return ctx.Value(configKey{}).(*Config)
}

func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Category: "core",
			Name:     "config",
			Usage:    "configuration file in json format",
			Value:    "vince.json",
			EnvVars:  []string{"VINCE_CONFIG"},
		},
		// server
		&cli.StringFlag{
			Category: "core",
			Name:     "listen",
			Usage:    "http address to listen to",
			Value:    ":8080",
			EnvVars:  []string{"VINCE_LISTEN"},
		},
		&cli.BoolFlag{
			Category: "tls",
			Name:     "enable-tls",
			Usage:    "Enables serving https traffick.",
			EnvVars:  []string{"VINCE_ENABLE_TLS"},
		},
		&cli.StringFlag{
			Category: "tls",
			Name:     "tls-address",
			Usage:    "https address to listen to. You must provide tls-key and tls-cert or configure auto-tls",
			Value:    ":8443",
			EnvVars:  []string{"VINCE_LISTEN_TLS"},
		},
		&cli.StringFlag{
			Category: "tls",
			Name:     "tls-key",
			Usage:    "Path to key file used for https",
			EnvVars:  []string{"VINCE_TLS_KEY"},
		},
		&cli.StringFlag{
			Category: "tls",
			Name:     "tls-cert",
			Usage:    "Path to certificate file used for https",
			EnvVars:  []string{"VINCE_TLS_CERT"},
		},
		&cli.StringFlag{
			Category: "core",
			Name:     "data",
			Usage:    "path to data directory",
			Value:    ".vince",
			EnvVars:  []string{"VINCE_DATA"},
		},
		&cli.StringFlag{
			Category: "core",
			Name:     "env",
			Usage:    "environment on which vince is run (dev,staging,production)",
			Value:    "dev",
			EnvVars:  []string{"VINCE_ENV"},
		},
		&cli.StringFlag{
			Category: "core",
			Name:     "url",
			Usage:    "url for the server on which vince is hosted(it shows up on emails)",
			Value:    "http://localhost:8080",
			EnvVars:  []string{"VINCE_URL"},
		},
		&cli.BoolFlag{
			Category: "email",
			Name:     "enable-email",
			Usage:    "allows sending emails",
			EnvVars:  []string{"VINCE_ENABLE_EMAIL"},
		},
		&cli.StringFlag{
			Category: "core",
			Name:     "log",
			Usage:    "level of logging",
			Value:    "info",
			EnvVars:  []string{"VINCE_LOG_LEVEL"},
		},
		&cli.StringFlag{
			Category: "core",
			Name:     "backup-dir",
			Usage:    "directory where backups are stored",
			EnvVars:  []string{"VINCE_BACKUP_DIR"},
		},
		&cli.StringFlag{
			Category: "email",
			Name:     "mailer-address",
			Usage:    "email address used for the sender of outgoing emails ",
			Value:    "vince@mailhog.example",
			EnvVars:  []string{"VINCE_MAILER_ADDRESS"},
		},
		&cli.StringFlag{
			Category: "email",
			Name:     "mailer-address-name",
			Usage:    "email address name  used for the sender of outgoing emails ",
			Value:    "gernest from vince analytics",
			EnvVars:  []string{"VINCE_MAILER_ADDRESS_NAME"},
		},
		&cli.StringFlag{
			Category: "email smtp",
			Name:     "mailer-smtp-host",
			Usage:    "host address of the smtp server used for outgoing emails",
			Value:    "localhost",
			EnvVars:  []string{"VINCE_MAILER_SMTP_HOST"},
		},
		&cli.IntFlag{
			Category: "email smtp",
			Name:     "mailer-smtp-port",
			Usage:    "port address of the smtp server used for outgoing emails",
			Value:    1025,
			EnvVars:  []string{"VINCE_MAILER_SMTP_PORT"},
		},
		&cli.StringFlag{
			Category: "email smtp anonymous",
			Name:     "mailer-smtp-anonymous",
			Usage:    "trace value for anonymous smtp auth",
			EnvVars:  []string{"VINCE_MAILER_SMTP_ANONYMOUS"},
		},
		&cli.StringFlag{
			Category: "email smtp plain",
			Name:     "mailer-smtp-plain-identity",
			Usage:    "identity value for plain smtp auth",
			EnvVars:  []string{"VINCE_MAILER_SMTP_PLAIN_IDENTITY"},
		},
		&cli.StringFlag{
			Category: "email smtp plain",
			Name:     "mailer-smtp-plain-username",
			Usage:    "username value for plain smtp auth",
			EnvVars:  []string{"VINCE_MAILER_SMTP_PLAIN_USERNAME"},
		},
		&cli.StringFlag{
			Category: "email smtp plain",
			Name:     "mailer-smtp-plain-password",
			Usage:    "password value for plain smtp auth",
			EnvVars:  []string{"VINCE_MAILER_SMTP_PLAIN_PASSWORD"},
		},
		&cli.StringFlag{
			Category: "email smtp oauth",
			Name:     "mailer-smtp-oauth-username",
			Usage:    "username value for oauth bearer smtp auth",
			EnvVars:  []string{"VINCE_MAILER_SMTP_OAUTH_USERNAME"},
		},
		&cli.StringFlag{
			Category: "email smtp oauth",
			Name:     "mailer-smtp-oauth-token",
			Usage:    "token value for oauth bearer smtp auth",
			EnvVars:  []string{"VINCE_MAILER_SMTP_OAUTH_TOKEN"},
		},
		&cli.StringFlag{
			Category: "email smtp oauth",
			Name:     "mailer-smtp-oauth-host",
			Usage:    "host value for oauth bearer smtp auth",
			EnvVars:  []string{"VINCE_MAILER_SMTP_OAUTH_HOST"},
		},
		&cli.IntFlag{
			Category: "email smtp oauth",
			Name:     "mailer-smtp-oauth-port",
			Usage:    "port value for oauth bearer smtp auth",
			EnvVars:  []string{"VINCE_MAILER_SMTP_OAUTH_PORT"},
		},
		&cli.DurationFlag{
			Category: "core",
			Name:     "cache-refresh",
			Usage:    "window for refreshing sites cache",
			Value:    15 * time.Minute,
			EnvVars:  []string{"VINCE_SITE_CACHE_REFRESH_INTERVAL"},
		},
		&cli.DurationFlag{
			Category: "core",
			Name:     "rotation-check",
			Usage:    "window for checking log rotation",
			Value:    time.Hour,
			EnvVars:  []string{"VINCE_LOG_ROTATION_CHECK_INTERVAL"},
		},
		&cli.DurationFlag{
			Category: "core",
			Name:     "ts-buffer",
			Usage:    "window for buffering timeseries in memory before savin them",
			// This seems reasonable to avoid users to wait for a long time between
			// creating the site and seeing something on the dashboard. A bigger
			// duration is better though, to reduce pressure on our kv store
			Value:   time.Minute,
			EnvVars: []string{"VINCE_TS_BUFFER_INTERVAL"},
		},
		// secrets
		&cli.StringFlag{
			Category: "secrets",
			Name:     "secret",
			Usage:    "path to a file with  ed25519 private key",
			EnvVars:  []string{"VINCE_SECRET"},
		},
		&cli.StringFlag{
			Category: "secrets",
			Name:     "secret-age",
			Usage:    "path to file with age.X25519Identity",
			EnvVars:  []string{"VINCE_SECRET_AGE"},
		},
		&cli.BoolFlag{
			Category: "tls",
			Name:     "enable-auto-tls",
			Usage:    "Enables using acme for automatic https.",
			EnvVars:  []string{"VINCE_AUTO_TLS"},
		},
		&cli.StringFlag{
			Category: "tls",
			Name:     "acme-email",
			Usage:    "Email address to use with letsencrypt.",
			EnvVars:  []string{"VINCE_ACME_EMAIL"},
		},
		&cli.StringFlag{
			Category: "tls",
			Name:     "acme-domain",
			Usage:    "Domain to use with letsencrypt.",
			EnvVars:  []string{"VINCE_ACME_DOMAIN"},
		},
		&cli.BoolFlag{
			Category: "bootstrap",
			Name:     "enable-bootstrap",
			Usage:    "allows creating a user and api key on startup.",
			EnvVars:  []string{"VINCE_ENABLE_BOOTSTRAP"},
		},
		&cli.StringFlag{
			Category: "bootstrap",
			Name:     "bootstrap-name",
			Usage:    "Full name of the user to bootstrap.",
			EnvVars:  []string{"VINCE_BOOTSTRAP_NAME"},
		},
		&cli.StringFlag{
			Category: "bootstrap",
			Name:     "bootstrap-email",
			Usage:    "Email address of the user to bootstrap.",
			EnvVars:  []string{"VINCE_BOOTSTRAP_EMAIL"},
		},
		&cli.StringFlag{
			Category: "bootstrap",
			Name:     "bootstrap-password",
			Usage:    "Password of the user to bootstrap.",
			EnvVars:  []string{"VINCE_BOOTSTRAP_PASSWORD"},
		},
		&cli.StringFlag{
			Category: "bootstrap",
			Name:     "bootstrap-key",
			Usage:    "API Key of the user to bootstrap.",
			EnvVars:  []string{"VINCE_BOOTSTRAP_KEY"},
		},
		&cli.BoolFlag{
			Category: "core",
			Name:     "enable-profile",
			Usage:    "Expose /debug/pprof endpoint",
			EnvVars:  []string{"VINCE_ENABLE_PROFILE"},
		},
		&cli.BoolFlag{
			Category: "core",
			Name:     "enable-alerts",
			Usage:    "allows loading and executing alerts",
			EnvVars:  []string{"VINCE_ENABLE_ALERTS"},
		},
	}
}

func Load(ctx *cli.Context) (*Config, context.Context, error) {
	base := fromCli(ctx)
	conf, err := fromFile(ctx)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, nil, err
		}
	} else {
		proto.Merge(base, conf)
	}
	sec, a, err := setupSecrets(base)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to setup secrets %v", err)
	}
	baseCtx := context.WithValue(context.Background(), configKey{}, base)
	baseCtx = context.WithValue(baseCtx, securityKey{}, sec)
	baseCtx = context.WithValue(baseCtx, ageKey{}, a)
	return base, baseCtx, nil
}

func fromFile(ctx *cli.Context) (*Config, error) {
	path := ctx.String("config")
	if path == "" {
		return &Config{}, nil
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var c Config
	err = protojson.Unmarshal(b, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func fromCli(ctx *cli.Context) *Config {
	x := &Config{
		Listen:    ctx.String("listen"),
		EnableTls: ctx.Bool("enable-tls"),
		Tls: &TLS{
			Address: ctx.String("tls-address"),
			Key:     ctx.String("tls-key"),
			Cert:    ctx.String("tls-cert"),
		},
		Url:           ctx.String("url"),
		DataPath:      ctx.String("data"),
		EnableEmail:   ctx.Bool("enable-email"),
		BackupDir:     ctx.String("backup-dir"),
		EnableAutoTls: ctx.Bool("enable-auto-tls"),
		Acme: &ACME{
			Email:  ctx.String("acme-email"),
			Domain: ctx.String("acme-domain"),
		},
		EnableBootstrap: ctx.Bool("enable-bootstrap"),
		Bootstrap: &Bootstrap{
			Name:     ctx.String("bootstrap-name"),
			Email:    ctx.String("bootstrap-email"),
			Password: ctx.String("bootstrap-password"),
			Key:      ctx.String("bootstrap-key"),
		},
		Secrets: &Secrets{
			Secret: ctx.String("secret"),
			Age:    ctx.String("secret-age"),
		},
		Intervals: &Intervals{
			SitesByDomainCacheRefreshInterval: durationpb.New(ctx.Duration("cache-refresh")),
			LogRotationCheckInterval:          durationpb.New(ctx.Duration("rotation-check")),
			SaveTimeseriesBufferInterval:      durationpb.New(ctx.Duration("ts-buffer")),
		},
		Mailer: &Config_Mailer{
			Address: &Config_Address{
				Name:  ctx.String("mailer-address-name"),
				Email: ctx.String("mailer-address"),
			},
			Smtp: &Config_Mailer_Smtp{
				Host: ctx.String("mailer-smtp-host"),
				Port: int32(ctx.Int("mailer-smtp-port")),
			},
		},
		Cors: &Cors{
			Origin:      "*",
			Credentials: true,
			MaxAge:      1_728_000,
			Headers: []string{
				"Authorization",
				"Content-Type",
				"Accept",
				"Origin",
				"User-Agent",
				"DNT",
				"Cache-Control",
				"X-Mx-ReqToken",
				"Keep-Alive",
				"X-Requested-With",
				"If-Modified-Since",
				"X-CSRF-Token",
			},
			Methods:               []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			SendPreflightResponse: true,
		},
		EnableProfile: ctx.Bool("enable-profile"),
		EnableAlerts:  ctx.Bool("enable-alerts"),
	}
	var anon *Config_Mailer_Anonymous
	var plain *Config_Mailer_Plain
	var oauth *Config_Mailer_OauthBearer
	if v := ctx.String("mailer-smtp-anonymous"); v != "" {
		anon = &Config_Mailer_Anonymous{Trace: v}
	}

	if v := ctx.String("mailer-smtp-plain-identity"); v != "" {
		if plain == nil {
			plain = &Config_Mailer_Plain{}
		}
		plain.Identity = v
	}
	if v := ctx.String("mailer-smtp-plain-username"); v != "" {
		if plain == nil {
			plain = &Config_Mailer_Plain{}
		}
		plain.Username = v
	}
	if v := ctx.String("mailer-smtp-plain-password"); v != "" {
		if plain == nil {
			plain = &Config_Mailer_Plain{}
		}
		plain.Password = v
	}
	if v := ctx.String("mailer-smtp-oauth-username"); v != "" {
		if oauth == nil {
			oauth = &Config_Mailer_OauthBearer{}
		}
		oauth.Username = v
	}
	if v := ctx.String("mailer-smtp-oauth-token"); v != "" {
		if oauth == nil {
			oauth = &Config_Mailer_OauthBearer{}
		}
		oauth.Token = v
	}
	if v := ctx.String("mailer-smtp-oauth-host"); v != "" {
		if oauth == nil {
			oauth = &Config_Mailer_OauthBearer{}
		}
		oauth.Host = v
	}
	if v := ctx.Int("mailer-smtp-oauth-port"); v != 0 {
		if oauth == nil {
			oauth = &Config_Mailer_OauthBearer{}
		}
		oauth.Port = int32(v)
	}
	switch {
	case anon != nil:
		x.Mailer.Smtp.Auth = &Config_Mailer_Smtp_Anonymous{Anonymous: anon}
	case plain != nil:
		x.Mailer.Smtp.Auth = &Config_Mailer_Smtp_Plain{Plain: plain}
	case oauth != nil:
		x.Mailer.Smtp.Auth = &Config_Mailer_Smtp_OauthBearer{OauthBearer: oauth}
	}
	return x
}

func (c *Config) Scrub() *Config {
	n := proto.Clone(c).(*Config)
	n.SuperUserId = nil
	n.Secrets = nil
	return n
}

func (c *Config) IsSuperUser(id uint64) bool {
	for _, u := range c.SuperUserId {
		if u == id {
			return true
		}
	}
	return false
}
