package site

import (
	"net/http"
	"time"

	"github.com/vinceanalytics/vince/internal/models"
	"github.com/vinceanalytics/vince/internal/render"
	"github.com/vinceanalytics/vince/internal/sessions"
	"github.com/vinceanalytics/vince/internal/templates"
	"github.com/vinceanalytics/vince/internal/timeseries"
)

func Stats(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	site := models.GetSite(ctx)
	role := models.GetRole(ctx)
	owner := models.SiteOwner(ctx, site.ID)
	var canSeeStats bool
	switch role {
	case "super_admin":
	default:
		canSeeStats = true
	}
	if canSeeStats {
		w.Header().Set("x-robots-tag", "noindex")
		var offer bool
		session, _ := sessions.Load(r)
		if session.Data.EmailReport != nil {
			offer = session.Data.EmailReport[site.Domain]
		}
		hasGoals := models.SiteHasGoals(ctx, site.Domain)
		ts := time.Now()
		timeseries.Query(ctx, timeseries.QueryRequest{
			UserID: owner.ID,
			SiteID: site.ID,
			BaseQuery: timeseries.BaseQuery{
				Start:  ts,
				Offset: 24 * time.Hour,
			},
		})
		render.HTML(ctx, w, templates.Stats, http.StatusOK, func(ctx *templates.Context) {
			ctx.Site = site
			ctx.Title = "Vince Analytics  · " + site.Domain
			ctx.EmailReport = offer
			ctx.HasGoals = hasGoals
		})
		return
	}
	render.ERROR(ctx, w, http.StatusUnauthorized)
}
