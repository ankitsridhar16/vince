package site

import (
	"net/http"

	"github.com/gernest/vince/render"
)

func ImportFromGoogle(w http.ResponseWriter, r *http.Request) {
	render.ERROR(r.Context(), w, http.StatusNotImplemented)
}
