package auth

import (
	"net/http"
	"strings"

	"github.com/vinceanalytics/vince/models"
	"github.com/vinceanalytics/vince/pkg/log"
	"github.com/vinceanalytics/vince/render"
	"github.com/vinceanalytics/vince/sessions"
	"github.com/vinceanalytics/vince/templates"
)

func Register(w http.ResponseWriter, r *http.Request) {
	session, r := sessions.Load(r)
	r.ParseForm()
	u := new(models.User)
	m, err := models.NewUser(u, r)
	if err != nil {
		log.Get().Err(err).Msg("Failed decoding new user from")
		render.ERROR(r.Context(), w, http.StatusInternalServerError)
		return
	}

	validCaptcha := session.VerifyCaptchaSolution(r)
	if len(m) > 0 || !validCaptcha {
		r = sessions.SaveCsrf(w, r)
		r = sessions.SaveCaptcha(w, r)
		render.HTML(r.Context(), w, templates.RegisterForm, http.StatusOK, func(ctx *templates.Context) {
			for k, v := range m {
				ctx.Errors[k] = v
			}
			if !validCaptcha {
				ctx.Errors["captcha"] = "Please complete the captcha to register"
			}
			ctx.Form = r.Form
		})
		return
	}
	if err := models.Get(r.Context()).Save(u).Error; err != nil {
		log.Get().Err(err).Msg("failed saving new user")
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			r = sessions.SaveCsrf(w, r)
			r = sessions.SaveCaptcha(w, r)
			render.HTML(r.Context(), w, templates.RegisterForm, http.StatusOK, func(ctx *templates.Context) {
				ctx.Errors["email"] = "already exists"
				ctx.Form = r.Form
			})
			return
		}
		render.ERROR(r.Context(), w, http.StatusInternalServerError)
		return
	}
	ctx := models.SetUser(r.Context(), u)
	session.Data.CurrentUserID = u.ID
	session.Data.LoggedIn = true
	session.Save(ctx, w)
	if u.EmailVerified {
		http.Redirect(w, r, "/new", http.StatusFound)
	} else {
		err := SendVerificationEmail(ctx, u)
		if err != nil {
			log.Get().Err(err).Msg("failed sending email message")
		}
		http.Redirect(w, r, "/activate", http.StatusFound)
	}
}