package sites

import (
	"net/http"

	"github.com/vinceanalytics/vince/internal/models"
	"github.com/vinceanalytics/vince/internal/params"
	"github.com/vinceanalytics/vince/internal/render"
)

func DeleteGoal(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	site := models.SiteFor(ctx,
		models.GetUser(ctx).ID,
		params.Get(ctx)["site_id"],
	)
	if site == nil {
		render.JSON(w, http.StatusNotFound, map[string]any{
			"error": http.StatusText(http.StatusNotFound),
		})
		return
	}
	models.DeleteGoal(ctx, params.Get(ctx)["goal_id"], site.Domain)
	render.JSON(w, http.StatusOK, map[string]any{
		"deleted": true,
	})
}
