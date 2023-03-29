package templates

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"net/url"
	"strings"

	"github.com/belak/octicon"
	"github.com/gernest/vince/config"
	"github.com/gernest/vince/flash"
	"github.com/gernest/vince/models"
)

//go:embed layout pages site auth error email
var files embed.FS

var LoginForm = template.Must(
	template.ParseFS(files,
		"layout/focus.html",
		"layout/flash.html",
		"layout/csrf.html",
		"layout/captcha.html",
		"auth/login_form.html",
	),
)

var RegisterForm = template.Must(
	template.ParseFS(files,
		"layout/focus.html",
		"layout/flash.html",
		"layout/csrf.html",
		"layout/captcha.html",
		"auth/register_form.html",
	),
)

var Error = template.Must(template.ParseFS(files,
	"error/error.html",
))

var ActivationEmail = template.Must(
	template.ParseFS(files,
		"layout/focus.html",
		"email/activation_code.html",
	),
)

var Activate = template.Must(
	template.ParseFS(files,
		"layout/focus.html",
		"layout/flash.html",
		"layout/csrf.html",
		"auth/activate.html",
	),
)

var Home = template.Must(
	template.ParseFS(files,
		"layout/app.html",
		"layout/header.html",
		"layout/flash.html",
		"layout/notice.html",
		"layout/footer.html",
	),
)

var Sites = template.Must(
	template.ParseFS(files,
		"layout/app.html",
		"layout/header.html",
		"layout/flash.html",
		"layout/notice.html",
		"site/index.html",
		"layout/footer.html",
	),
)

var SiteNew = template.Must(
	template.ParseFS(files,
		"layout/app.html",
		"layout/csrf.html",
		"layout/header.html",
		"layout/flash.html",
		"layout/notice.html",
		"layout/footer.html",
		"site/new.html",
	),
)

var Pricing = template.Must(
	template.ParseFS(files,
		"layout/app.html",
		"layout/csrf.html",
		"layout/header.html",
		"layout/flash.html",
		"layout/notice.html",
		"layout/footer.html",
		"pages/pricing.html",
	),
)

type NewSite struct {
	IsFirstSite bool
	IsAtLimit   bool
	SiteLimit   int
}

type Errors struct {
	Status     int
	StatusText string
}

type Context struct {
	Title         string
	CurrentUser   *models.User
	Data          map[string]any
	CSRF          template.HTML
	Captcha       template.HTMLAttr
	Errors        map[string]string
	Form          url.Values
	Code          uint64
	Config        *config.Config
	HasInvitation bool
	HasPin        bool
	Flash         *flash.Flash
	NewSite       *NewSite
	Error         *Errors
	Position      int
	Page          string
}

func New(ctx context.Context, f ...func(c *Context)) *Context {
	c := &Context{
		Data:        make(map[string]any),
		CSRF:        getCsrf(ctx),
		Captcha:     getCaptcha(ctx),
		CurrentUser: models.GetCurrentUser(ctx),
		Config:      config.Get(ctx),
		Code:        GetActivationCode(ctx),
		Flash:       flash.Get(ctx),
	}
	if len(f) > 0 {
		f[0](c)
	}
	return c
}

type csrfTokenCtxKey struct{}

func getCsrf(ctx context.Context) template.HTML {
	if c := ctx.Value(csrfTokenCtxKey{}); c != nil {
		return c.(template.HTML)
	}
	return template.HTML("")
}

type captchaTokenKey struct{}

func SetCaptcha(ctx context.Context, x template.HTMLAttr) context.Context {
	return context.WithValue(ctx, captchaTokenKey{}, x)
}

func SetCSRF(ctx context.Context, x template.HTML) context.Context {
	return context.WithValue(ctx, csrfTokenCtxKey{}, x)
}

func getCaptcha(ctx context.Context) template.HTMLAttr {
	if c := ctx.Value(captchaTokenKey{}); c != nil {
		return c.(template.HTMLAttr)
	}
	return template.HTMLAttr("")
}

func (t *Context) VinceURL() template.HTML {
	return template.HTML("http://localhost:8080")
}

func (t *Context) SetPosition(pos int) *Context {
	t.Position = pos
	return t
}

func (t *Context) OnboardSteps() []string {
	return []string{
		"Register", "Activate account", "Add site info", "Install snippet",
	}
}

func (t *Context) Validate(name string) template.HTML {
	if t.Errors != nil {
		o, _ := octicon.Icon("alert-fill", 12)
		if v, ok := t.Errors[name]; ok {
			return template.HTML(fmt.Sprintf(`
<div class="FormControl-inlineValidation">
    %s
    <span>%s</span>
</div>
		`, o, v))
		}
	}
	return template.HTML("")
}

func (t *Context) Icon(name string, height int, class ...string) (template.HTML, error) {
	return octicon.IconTemplateFunc(name, height, class...)
}

func (t *Context) InputField(name string) template.HTMLAttr {
	var s strings.Builder
	if t.Errors != nil && t.Errors[name] != "" {
		s.WriteString(`invalid="true"`)
	}
	if t.Form != nil && t.Form.Get(name) != "" {
		s.WriteString(fmt.Sprintf("value=%q", t.Form.Get(name)))
	}
	return template.HTMLAttr(s.String())
}

var steps = []string{
	"Register", "Activate account", "Add site info", "Install snippet",
}

func (t *Context) CurrentStep() string {
	return steps[t.Position]
}

func (t *Context) StepsBefore() []string {
	return steps[0:t.Position]
}

func (t *Context) StepsAfter() []string {
	return steps[t.Position:]
}

type activationCodeKey struct{}

func SetActivationCode(ctx context.Context, code uint64) context.Context {
	return context.WithValue(ctx, activationCodeKey{}, code)
}

func GetActivationCode(ctx context.Context) uint64 {
	if v := ctx.Value(activationCodeKey{}); v != nil {
		return v.(uint64)
	}
	return 0
}
