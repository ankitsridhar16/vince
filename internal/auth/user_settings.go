package auth

import (
	"encoding/base64"
	"net/http"

	"github.com/vinceanalytics/vince/models"
	"github.com/vinceanalytics/vince/pkg/secrets"
	"github.com/vinceanalytics/vince/render"
	"github.com/vinceanalytics/vince/templates"
)

func UserSettings(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	usr := models.GetUser(ctx)
	models.PreloadUser(ctx, usr, "APIKeys")
	render.HTML(ctx, w, templates.UserSettings, http.StatusOK, func(ctx *templates.Context) {
		ctx.Key = base64.StdEncoding.EncodeToString(secrets.APIKey())
		ctx.CurrentUser = usr
	})
}