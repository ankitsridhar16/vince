package auth

import (
	"net/http"

	"github.com/vinceanalytics/vince/render"
)

func PasswordReset(w http.ResponseWriter, r *http.Request) {
	render.ERROR(r.Context(), w, http.StatusNotImplemented)
}
