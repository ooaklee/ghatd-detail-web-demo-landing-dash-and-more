package webapp

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"

	"go.uber.org/zap"

	webapphelpers "github.com/ooaklee/ghatd-detail-web-demo-landing-dash-and-more/external/helpers"
	"github.com/ooaklee/ghatd/external/common"
	"github.com/ooaklee/ghatd/external/items/policy"
	"github.com/ooaklee/ghatd/external/logger"
)

// NewWebAppHandlerRequest is the request needed to create a web app handler
type NewWebAppHandlerRequest struct {
	EmbeddedContent fs.FS
	// EmbeddedContentFilePathPrefix the prefix used to access the embedded files
	EmbeddedContentFilePathPrefix string
	TermsOfServicePolicy          *policy.WebAppPolicy
	PrivacyPolicy                 *policy.WebAppPolicy
	CookiePolicy                  *policy.WebAppPolicy
}

// NewWebAppHandler creates a new instance of a web app handler
func NewWebAppHandler(r *NewWebAppHandlerRequest) *Handler {
	return &Handler{
		embeddedFileSystem:            r.EmbeddedContent,
		embeddedContentFilePathPrefix: r.EmbeddedContentFilePathPrefix,
		termsOfServicePolicy:          r.TermsOfServicePolicy,
		privacyPolicy:                 r.PrivacyPolicy,
		cookiePolicy:                  r.CookiePolicy,
	}
}

// Handler manages request for webapp
type Handler struct {
	embeddedFileSystem            fs.FS
	embeddedContentFilePathPrefix string
	termsOfServicePolicy          *policy.WebAppPolicy
	privacyPolicy                 *policy.WebAppPolicy
	cookiePolicy                  *policy.WebAppPolicy
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {

	logger := logger.AcquireFrom(r.Context())

	// If the path is not exactly "/"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// list of template files to parse, must be in order of inheritence
	templateFilesToParse := []string{
		fmt.Sprintf("%sui/html/base.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/pages/home.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/tailwind-dash-script.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/header.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/footer.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/footer-nav-links-info.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/social-links.tmpl.html", h.embeddedContentFilePathPrefix),
	}

	// Parse template
	parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
	if err != nil {
		logger.Error("Unable to parse referenced template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write template to response
	err = parsedTemplates.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error("Unable to execute parsed template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Terms(w http.ResponseWriter, r *http.Request) {

	logger := logger.AcquireFrom(r.Context())

	// list of template files to parse, must be in order of inheritence
	templateFilesToParse := []string{
		fmt.Sprintf("%sui/html/base.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/pages/base-policy.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/tailwind-dash-script.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/header.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/footer.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/footer-nav-links-info.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/social-links.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/policy/policy-holder.tmpl.html", h.embeddedContentFilePathPrefix),
	}

	// Parse template
	parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
	if err != nil {
		logger.Error("Unable to parse referenced template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write template to response
	err = parsedTemplates.ExecuteTemplate(w, "base", h.termsOfServicePolicy)
	if err != nil {
		logger.Error("Unable to execute parsed template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Privacy(w http.ResponseWriter, r *http.Request) {

	logger := logger.AcquireFrom(r.Context())

	// list of template files to parse, must be in order of inheritence
	templateFilesToParse := []string{
		fmt.Sprintf("%sui/html/base.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/pages/base-policy.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/tailwind-dash-script.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/header.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/footer.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/footer-nav-links-info.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/social-links.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/policy/policy-holder.tmpl.html", h.embeddedContentFilePathPrefix),
	}

	// Parse template
	parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
	if err != nil {
		logger.Error("Unable to parse referenced template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write template to response
	err = parsedTemplates.ExecuteTemplate(w, "base", h.privacyPolicy)
	if err != nil {
		logger.Error("Unable to execute parsed template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func (h *Handler) Cookie(w http.ResponseWriter, r *http.Request) {

	logger := logger.AcquireFrom(r.Context())

	// list of template files to parse, must be in order of inheritence
	templateFilesToParse := []string{
		fmt.Sprintf("%sui/html/base.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/pages/base-policy.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/tailwind-dash-script.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/header.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/footer.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/footer-nav-links-info.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/social-links.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/policy/policy-holder.tmpl.html", h.embeddedContentFilePathPrefix),
	}

	// Parse template
	parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
	if err != nil {
		logger.Error("Unable to parse referenced template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write template to response
	err = parsedTemplates.ExecuteTemplate(w, "base", h.cookiePolicy)
	if err != nil {
		logger.Error("Unable to execute parsed template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func (h *Handler) AuthLogin(w http.ResponseWriter, r *http.Request) {

	logger := logger.AcquireFrom(r.Context())

	// list of template files to parse, must be in order of inheritence
	templateFilesToParse := []string{
		fmt.Sprintf("%sui/html/base.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/pages/base-auth.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/auth/auth-login.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/tailwind-dash-script.tmpl.html", h.embeddedContentFilePathPrefix),
	}

	// Parse template
	parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
	if err != nil {
		logger.Error("Unable to parse referenced template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write template to response
	err = parsedTemplates.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error("Unable to execute parsed template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) AuthSignup(w http.ResponseWriter, r *http.Request) {

	logger := logger.AcquireFrom(r.Context())

	// list of template files to parse, must be in order of inheritence
	templateFilesToParse := []string{
		fmt.Sprintf("%sui/html/base.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/pages/base-auth.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/auth/auth-signup.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/tailwind-dash-script.tmpl.html", h.embeddedContentFilePathPrefix),
	}

	// Parse template
	parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
	if err != nil {
		logger.Error("Unable to parse referenced template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write template to response
	err = parsedTemplates.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error("Unable to execute parsed template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) AuthResetPassword(w http.ResponseWriter, r *http.Request) {

	logger := logger.AcquireFrom(r.Context())

	// list of template files to parse, must be in order of inheritence
	templateFilesToParse := []string{
		fmt.Sprintf("%sui/html/base.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/pages/base-auth.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/auth/auth-reset-password.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/tailwind-dash-script.tmpl.html", h.embeddedContentFilePathPrefix),
	}

	// Parse template
	parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
	if err != nil {
		logger.Error("Unable to parse referenced template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write template to response
	err = parsedTemplates.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error("Unable to execute parsed template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) ComingSoon(w http.ResponseWriter, r *http.Request) {

	logger := logger.AcquireFrom(r.Context())

	// list of template files to parse, must be in order of inheritence
	templateFilesToParse := []string{
		fmt.Sprintf("%sui/html/base.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/pages/coming-soon.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/tailwind-dash-script.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/preloader.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/coming-soon/countdown-timer.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/social-links.tmpl.html", h.embeddedContentFilePathPrefix),
	}

	// Parse template
	parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
	if err != nil {
		logger.Error("Unable to parse referenced templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write template to response
	err = parsedTemplates.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error("Unable to execute parsed templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Dash(w http.ResponseWriter, r *http.Request) {

	logger := logger.AcquireFrom(r.Context())

	if r.Header.Get(common.WebPartialHttpRequestHeader) == "true" {

		w.Header().Add(common.CacheSkipHttpResponseHeader, "true")

		// list of template files to parse, must be in order of inheritence
		templateFilesToParse := []string{
			fmt.Sprintf("%sui/html/partials/dash/dash-ecommerce.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/shared/chart-area.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/shared/chart-bar.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/shared/chart-donut.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/dash/map-01.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/dash/table-01.tmpl.html", h.embeddedContentFilePathPrefix),
		}

		// Parse template
		parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
		if err != nil {
			logger.Error("Unable to parse referenced template", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Write template to response
		err = parsedTemplates.ExecuteTemplate(w, "dash-main", webapphelpers.UpdateSiteTitleHelper{
			EnableTitleUpdate: true,
			NewTitle:          "E Commerce",
		})
		if err != nil {
			logger.Error("Unable to execute parsed templates", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		return
	}

	// list of template files to parse, must be in order of inheritence
	templateFilesToParse := []string{
		fmt.Sprintf("%sui/html/base.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/pages/dash.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-ecommerce.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/tailwind-dash-script.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/preloader.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-sidebar.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-header.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/chart-area.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/chart-bar.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/chart-donut.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/map-01.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/table-01.tmpl.html", h.embeddedContentFilePathPrefix),
	}

	// Parse template
	parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
	if err != nil {
		logger.Error("Unable to parse referenced templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write template to response
	err = parsedTemplates.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error("Unable to execute parsed templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DashCalendar(w http.ResponseWriter, r *http.Request) {

	logger := logger.AcquireFrom(r.Context())

	if r.Header.Get(common.WebPartialHttpRequestHeader) == "true" {

		w.Header().Add(common.CacheSkipHttpResponseHeader, "true")

		// Parse template
		parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, fmt.Sprintf("%sui/html/partials/dash/dash-calendar.tmpl.html", h.embeddedContentFilePathPrefix))
		if err != nil {
			logger.Error("Unable to parse referenced template", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Write template to response
		err = parsedTemplates.ExecuteTemplate(w, "dash-main", webapphelpers.UpdateSiteTitleHelper{
			EnableTitleUpdate: true,
			NewTitle:          "Calendar",
		})
		if err != nil {
			logger.Error("Unable to execute parsed templates", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		return
	}

	// list of template files to parse, must be in order of inheritence
	templateFilesToParse := []string{
		fmt.Sprintf("%sui/html/base.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/pages/dash.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-sidebar.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-calendar.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/tailwind-dash-script.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/preloader.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-header.tmpl.html", h.embeddedContentFilePathPrefix),
	}

	// Parse template
	parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
	if err != nil {
		logger.Error("Unable to parse referenced templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write template to response
	err = parsedTemplates.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error("Unable to execute parsed templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DashProfile(w http.ResponseWriter, r *http.Request) {

	logger := logger.AcquireFrom(r.Context())

	if r.Header.Get(common.WebPartialHttpRequestHeader) == "true" {

		w.Header().Add(common.CacheSkipHttpResponseHeader, "true")

		// Parse template
		parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, fmt.Sprintf("%sui/html/partials/dash/dash-profile.tmpl.html", h.embeddedContentFilePathPrefix))
		if err != nil {
			logger.Error("Unable to parse referenced template", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Write template to response
		err = parsedTemplates.ExecuteTemplate(w, "dash-main", webapphelpers.UpdateSiteTitleHelper{
			EnableTitleUpdate: true,
			NewTitle:          "Profile",
		})
		if err != nil {
			logger.Error("Unable to execute parsed templates", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		return
	}

	// list of template files to parse, must be in order of inheritence
	templateFilesToParse := []string{
		fmt.Sprintf("%sui/html/base.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/pages/dash.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-sidebar.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-profile.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/tailwind-dash-script.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/preloader.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-header.tmpl.html", h.embeddedContentFilePathPrefix),
	}

	// Parse template
	parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
	if err != nil {
		logger.Error("Unable to parse referenced templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write template to response
	err = parsedTemplates.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error("Unable to execute parsed templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DashBlank(w http.ResponseWriter, r *http.Request) {

	logger := logger.AcquireFrom(r.Context())

	if r.Header.Get(common.WebPartialHttpRequestHeader) == "true" {

		w.Header().Add(common.CacheSkipHttpResponseHeader, "true")

		// Parse template
		parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, fmt.Sprintf("%sui/html/partials/dash/dash-blank.tmpl.html", h.embeddedContentFilePathPrefix))
		if err != nil {
			logger.Error("Unable to parse referenced template", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Write template to response
		err = parsedTemplates.ExecuteTemplate(w, "dash-main", webapphelpers.UpdateSiteTitleHelper{
			EnableTitleUpdate: true,
			NewTitle:          "Blank Page",
		})
		if err != nil {
			logger.Error("Unable to execute parsed templates", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		return
	}

	// list of template files to parse, must be in order of inheritence
	templateFilesToParse := []string{
		fmt.Sprintf("%sui/html/base.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/pages/dash.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-sidebar.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-blank.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/tailwind-dash-script.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/preloader.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-header.tmpl.html", h.embeddedContentFilePathPrefix),
	}

	// Parse template
	parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
	if err != nil {
		logger.Error("Unable to parse referenced templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write template to response
	err = parsedTemplates.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error("Unable to execute parsed templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DashFormElements(w http.ResponseWriter, r *http.Request) {

	logger := logger.AcquireFrom(r.Context())

	if r.Header.Get(common.WebPartialHttpRequestHeader) == "true" {

		w.Header().Add(common.CacheSkipHttpResponseHeader, "true")

		// list of template files to parse, must be in order of inheritence
		templateFilesToParse := []string{
			fmt.Sprintf("%sui/html/partials/dash/dash-form-elements.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/dash/form-elements-checkbox-radio.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/dash/form-elements-file-upload.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/dash/form-elements-input-fields.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/dash/form-elements-select-input.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/dash/form-elements-switch-input.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/dash/form-elements-textarea-fields.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/dash/form-elements-time-date.tmpl.html", h.embeddedContentFilePathPrefix),
		}
		// Parse template
		parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
		if err != nil {
			logger.Error("Unable to parse referenced template", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Write template to response
		err = parsedTemplates.ExecuteTemplate(w, "dash-main", webapphelpers.UpdateSiteTitleHelper{
			EnableTitleUpdate: true,
			NewTitle:          "Form Elements",
		})
		if err != nil {
			logger.Error("Unable to execute parsed templates", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		return
	}

	// list of template files to parse, must be in order of inheritence
	templateFilesToParse := []string{
		fmt.Sprintf("%sui/html/base.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/pages/dash.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-sidebar.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-form-elements.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/form-elements-checkbox-radio.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/form-elements-file-upload.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/form-elements-input-fields.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/form-elements-select-input.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/form-elements-switch-input.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/form-elements-textarea-fields.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/form-elements-time-date.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/tailwind-dash-script.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/preloader.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-header.tmpl.html", h.embeddedContentFilePathPrefix),
	}

	// Parse template
	parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
	if err != nil {
		logger.Error("Unable to parse referenced templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write template to response
	err = parsedTemplates.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error("Unable to execute parsed templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DashFormLayout(w http.ResponseWriter, r *http.Request) {

	logger := logger.AcquireFrom(r.Context())

	if r.Header.Get(common.WebPartialHttpRequestHeader) == "true" {

		w.Header().Add(common.CacheSkipHttpResponseHeader, "true")

		// list of template files to parse, must be in order of inheritence
		templateFilesToParse := []string{
			fmt.Sprintf("%sui/html/partials/dash/dash-form-layout.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/dash/form-layout-contact-form.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/dash/form-layout-sign-up-form.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/dash/form-layout-sign-in-form.tmpl.html", h.embeddedContentFilePathPrefix),
		}
		// Parse template
		parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
		if err != nil {
			logger.Error("Unable to parse referenced template", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Write template to response
		err = parsedTemplates.ExecuteTemplate(w, "dash-main", webapphelpers.UpdateSiteTitleHelper{
			EnableTitleUpdate: true,
			NewTitle:          "Form Layout",
		})
		if err != nil {
			logger.Error("Unable to execute parsed templates", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		return
	}

	// list of template files to parse, must be in order of inheritence
	templateFilesToParse := []string{
		fmt.Sprintf("%sui/html/base.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/pages/dash.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-sidebar.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-form-layout.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/form-layout-contact-form.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/form-layout-sign-up-form.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/form-layout-sign-in-form.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/tailwind-dash-script.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/preloader.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-header.tmpl.html", h.embeddedContentFilePathPrefix),
	}

	// Parse template
	parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
	if err != nil {
		logger.Error("Unable to parse referenced templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write template to response
	err = parsedTemplates.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error("Unable to execute parsed templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DashTables(w http.ResponseWriter, r *http.Request) {

	logger := logger.AcquireFrom(r.Context())

	if r.Header.Get(common.WebPartialHttpRequestHeader) == "true" {

		w.Header().Add(common.CacheSkipHttpResponseHeader, "true")

		// list of template files to parse, must be in order of inheritence
		templateFilesToParse := []string{
			fmt.Sprintf("%sui/html/partials/dash/dash-tables.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/dash/table-01.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/dash/table-02.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/dash/table-03.tmpl.html", h.embeddedContentFilePathPrefix),
		}
		// Parse template
		parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
		if err != nil {
			logger.Error("Unable to parse referenced template", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Write template to response
		err = parsedTemplates.ExecuteTemplate(w, "dash-main", webapphelpers.UpdateSiteTitleHelper{
			EnableTitleUpdate: true,
			NewTitle:          "Tables",
		})
		if err != nil {
			logger.Error("Unable to execute parsed templates", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		return
	}

	// list of template files to parse, must be in order of inheritence
	templateFilesToParse := []string{
		fmt.Sprintf("%sui/html/base.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/pages/dash.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-sidebar.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-tables.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/table-01.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/table-02.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/table-03.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/tailwind-dash-script.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/preloader.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-header.tmpl.html", h.embeddedContentFilePathPrefix),
	}

	// Parse template
	parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
	if err != nil {
		logger.Error("Unable to parse referenced templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write template to response
	err = parsedTemplates.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error("Unable to execute parsed templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DashSettings(w http.ResponseWriter, r *http.Request) {

	logger := logger.AcquireFrom(r.Context())

	if r.Header.Get(common.WebPartialHttpRequestHeader) == "true" {

		w.Header().Add(common.CacheSkipHttpResponseHeader, "true")

		// list of template files to parse, must be in order of inheritence
		templateFilesToParse := []string{
			fmt.Sprintf("%sui/html/partials/dash/dash-settings.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/dash/settings-profile.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/dash/settings-photo.tmpl.html", h.embeddedContentFilePathPrefix),
		}
		// Parse template
		parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
		if err != nil {
			logger.Error("Unable to parse referenced template", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Write template to response
		err = parsedTemplates.ExecuteTemplate(w, "dash-main", webapphelpers.UpdateSiteTitleHelper{
			EnableTitleUpdate: true,
			NewTitle:          "Settings",
		})
		if err != nil {
			logger.Error("Unable to execute parsed templates", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		return
	}

	// list of template files to parse, must be in order of inheritence
	templateFilesToParse := []string{
		fmt.Sprintf("%sui/html/base.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/pages/dash.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-sidebar.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-settings.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/settings-profile.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/settings-photo.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/tailwind-dash-script.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/preloader.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-header.tmpl.html", h.embeddedContentFilePathPrefix),
	}

	// Parse template
	parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
	if err != nil {
		logger.Error("Unable to parse referenced templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write template to response
	err = parsedTemplates.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error("Unable to execute parsed templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DashCharts(w http.ResponseWriter, r *http.Request) {

	logger := logger.AcquireFrom(r.Context())

	if r.Header.Get(common.WebPartialHttpRequestHeader) == "true" {

		w.Header().Add(common.CacheSkipHttpResponseHeader, "true")

		// list of template files to parse, must be in order of inheritence
		templateFilesToParse := []string{
			fmt.Sprintf("%sui/html/partials/dash/dash-charts.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/shared/chart-area.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/shared/chart-bar.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/shared/chart-donut.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/shared/chart-bar-2.tmpl.html", h.embeddedContentFilePathPrefix),
		}
		// Parse template
		parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
		if err != nil {
			logger.Error("Unable to parse referenced template", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Write template to response
		err = parsedTemplates.ExecuteTemplate(w, "dash-main", webapphelpers.UpdateSiteTitleHelper{
			EnableTitleUpdate: true,
			NewTitle:          "Charts",
		})
		if err != nil {
			logger.Error("Unable to execute parsed templates", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		return
	}

	// list of template files to parse, must be in order of inheritence
	templateFilesToParse := []string{
		fmt.Sprintf("%sui/html/base.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/pages/dash.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-sidebar.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-charts.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/chart-area.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/chart-bar.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/chart-donut.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/chart-bar-2.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/tailwind-dash-script.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/preloader.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-header.tmpl.html", h.embeddedContentFilePathPrefix),
	}

	// Parse template
	parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
	if err != nil {
		logger.Error("Unable to parse referenced templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write template to response
	err = parsedTemplates.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error("Unable to execute parsed templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DashAlerts(w http.ResponseWriter, r *http.Request) {

	logger := logger.AcquireFrom(r.Context())

	if r.Header.Get(common.WebPartialHttpRequestHeader) == "true" {

		w.Header().Add(common.CacheSkipHttpResponseHeader, "true")

		// list of template files to parse, must be in order of inheritence
		templateFilesToParse := []string{
			fmt.Sprintf("%sui/html/partials/dash/dash-alerts.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/shared/alert-error.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/shared/alert-success.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/shared/alert-warning.tmpl.html", h.embeddedContentFilePathPrefix),
		}
		// Parse template
		parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
		if err != nil {
			logger.Error("Unable to parse referenced template", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Write template to response
		err = parsedTemplates.ExecuteTemplate(w, "dash-main", webapphelpers.UpdateSiteTitleHelper{
			EnableTitleUpdate: true,
			NewTitle:          "Alerts",
		})
		if err != nil {
			logger.Error("Unable to execute parsed templates", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		return
	}

	// list of template files to parse, must be in order of inheritence
	templateFilesToParse := []string{
		fmt.Sprintf("%sui/html/base.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/pages/dash.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-sidebar.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-alerts.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/alert-error.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/alert-success.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/alert-warning.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/tailwind-dash-script.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/preloader.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-header.tmpl.html", h.embeddedContentFilePathPrefix),
	}

	// Parse template
	parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
	if err != nil {
		logger.Error("Unable to parse referenced templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write template to response
	err = parsedTemplates.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error("Unable to execute parsed templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DashButtons(w http.ResponseWriter, r *http.Request) {

	logger := logger.AcquireFrom(r.Context())

	if r.Header.Get(common.WebPartialHttpRequestHeader) == "true" {

		w.Header().Add(common.CacheSkipHttpResponseHeader, "true")

		// list of template files to parse, must be in order of inheritence
		templateFilesToParse := []string{
			fmt.Sprintf("%sui/html/partials/dash/dash-buttons.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/dash/dash-buttons.tmpl.html", h.embeddedContentFilePathPrefix),
			fmt.Sprintf("%sui/html/partials/dash/dash-calendar.tmpl.html", h.embeddedContentFilePathPrefix),
		}
		// Parse template
		parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
		if err != nil {
			logger.Error("Unable to parse referenced template", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Write template to response
		err = parsedTemplates.ExecuteTemplate(w, "dash-main", webapphelpers.UpdateSiteTitleHelper{
			EnableTitleUpdate: true,
			NewTitle:          "Buttons",
		})
		if err != nil {
			logger.Error("Unable to execute parsed templates", zap.Error(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		return
	}

	// list of template files to parse, must be in order of inheritence
	templateFilesToParse := []string{
		fmt.Sprintf("%sui/html/base.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/pages/dash.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-sidebar.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-buttons.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-buttons.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-calendar.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/tailwind-dash-script.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/shared/preloader.tmpl.html", h.embeddedContentFilePathPrefix),
		fmt.Sprintf("%sui/html/partials/dash/dash-header.tmpl.html", h.embeddedContentFilePathPrefix),
	}

	// Parse template
	parsedTemplates, err := template.ParseFS(h.embeddedFileSystem, templateFilesToParse...)
	if err != nil {
		logger.Error("Unable to parse referenced templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write template to response
	err = parsedTemplates.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error("Unable to execute parsed templates", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
