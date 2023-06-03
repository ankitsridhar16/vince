package auth

import (
	"net/http"

	"github.com/vinceanalytics/vince/render"
	"github.com/vinceanalytics/vince/templates"
)

func LoginForm(w http.ResponseWriter, r *http.Request) {
	render.HTML(r.Context(), w, templates.LoginForm, http.StatusOK)
}