package plug

import (
	"net/http"
	"strings"

	"github.com/vinceanalytics/vince/internal/tracker"
)

type Plug func(http.Handler) http.Handler

type Pipeline []Plug

func (p Pipeline) Pass(h http.HandlerFunc) http.Handler {
	x := http.Handler(h)
	for i := range p {
		x = p[len(p)-1-i](x)
	}
	return x
}

func (p Pipeline) And(n ...Plug) Pipeline {
	return append(p, n...)
}

func NOOP(w http.ResponseWriter, r *http.Request) {}

func Track() Plug {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/js/vince") {
				w.Header().Set("x-content-type-options", "nosniff")
				w.Header().Set("cross-origin-resource-policy", "cross-origin")
				w.Header().Set("access-control-allow-origin", "*")
				w.Header().Set("cache-control", "public, max-age=86400, must-revalidate")
				tracker.Serve(w, r)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func PutSecureBrowserHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("x-frame-options", "SAMEORIGIN")
		w.Header().Set("x-xss-protection", "1; mode=block")
		w.Header().Set("x-content-type-options", "nosniff")
		w.Header().Set("x-download-options", "noopen")
		w.Header().Set("x-permitted-cross-domain-policies", "none")
		w.Header().Set("cross-origin-window-policy", "deny")
		h.ServeHTTP(w, r)
	})
}
