package k8s

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"log/slog"

	"github.com/cenkalti/backoff/v4"
	"github.com/vinceanalytics/vince/internal/must"
	"github.com/vinceanalytics/vince/internal/secrets"
	"github.com/vinceanalytics/vince/internal/v8s/gen/client/vince/clientset/versioned"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const ResyncPeriod = 5 * time.Minute

type Client interface {
	Kube() kubernetes.Interface
	Vince() versioned.Interface
	Site() SiteAPI
}

type SiteAPI interface {
	Create(ctx context.Context, secret *v1.Secret, domain string, public bool) error
	Delete(ctx context.Context, secret *v1.Secret, domain string) error
}

func New(masterURL, kubeConfig string) Client {
	return build(masterURL, kubeConfig)
}

func config(masterURL, kubeConfig string) *rest.Config {
	if os.Getenv("KUBERNETES_SERVICE_HOST") != "" && os.Getenv("KUBERNETES_SERVICE_PORT") != "" {
		// If these env vars are set, we can build an in-cluster config.
		slog.Debug("creating in-cluster client")
		c, err := rest.InClusterConfig()
		if err != nil {
			slog.Error("failed to create k8s client", "err", err)
			os.Exit(1)
		}
		return c
	}
	if masterURL != "" || kubeConfig != "" {
		slog.Info("creating external k8s client ",
			"master_url", masterURL,
			"kube_config", kubeConfig,
		)
		c, err := clientcmd.BuildConfigFromFlags(masterURL, kubeConfig)
		if err != nil {
			slog.Error("failed creating external k8s client ",
				"master_url", masterURL,
				"kube_config", kubeConfig,
			)
			os.Exit(1)
			return nil
		}
		return c
	}
	slog.Error("missing masterURL or kubeConfig")
	os.Exit(1)
	return nil
}

var _ Client = (*wrap)(nil)

type wrap struct {
	k8s   *kubernetes.Clientset
	vince *versioned.Clientset
	site  siteAPI
}

func (w *wrap) Kube() kubernetes.Interface {
	return w.k8s
}

func (w *wrap) Vince() versioned.Interface {
	return w.vince
}

func (w *wrap) Site() SiteAPI {
	return &w.site
}

func build(masterURL, kubeConfig string) *wrap {
	r := config(masterURL, kubeConfig)
	if r == nil {
		return nil
	}
	k8s := must.Must(kubernetes.NewForConfig(r))(
		"failed to build k8s client",
	)
	site := must.Must(versioned.NewForConfig(r))(
		"failed to build site client",
	)
	return &wrap{k8s: k8s, vince: site}
}

type siteAPI struct {
	client http.Client
}

func (s *siteAPI) request(secret *v1.Secret, method, path string, body io.Reader) (*http.Request, error) {
	uri := fmt.Sprintf("http://%s.%s.svc.cluster.local:80/api/v1/sites%s", secret.Name, secret.Namespace, path)
	r, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}
	bearer := base64.StdEncoding.EncodeToString(secret.Data[secrets.API_KEY])
	r.Header.Set("Authorization", "Bearer "+bearer)
	return r, nil
}

func (s *siteAPI) Create(ctx context.Context, secret *v1.Secret, domain string, public bool) error {
	b, _ := json.Marshal(map[string]any{
		"domain": domain,
		"public": public,
	})
	r, err := s.request(secret, http.MethodPost, "", bytes.NewReader(b))
	if err != nil {
		return err
	}
	return s.doRetry(r)
}

func (s *siteAPI) Delete(ctx context.Context, secret *v1.Secret, domain string) error {
	domain = url.PathEscape(domain)
	r, err := s.request(secret, http.MethodDelete, "/"+domain, nil)
	if err != nil {
		return err
	}
	return s.doRetry(r)
}

func (s *siteAPI) do(r *http.Request) error {
	res, err := s.client.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(http.StatusText(res.StatusCode))
	}
	return nil
}

func (s *siteAPI) doRetry(r *http.Request) error {
	return backoff.Retry(func() error {
		return s.do(r.Clone(r.Context()))
	}, backoff.NewExponentialBackOff())
}