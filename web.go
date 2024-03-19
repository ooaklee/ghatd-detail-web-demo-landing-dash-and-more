package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ooaklee/ghatd/external/response"
	"github.com/ooaklee/ghatd/external/router"

	//>ghatd {{ define "WebDetailImports" }}
	webapp "github.com/ooaklee/ghatd-detail-web-demo-landing-dash-and-more/external"
	saasPolicy "github.com/ooaklee/ghatd-detail-web-demo-landing-dash-and-more/external/policy/saas"
	//>ghatd {{ end }}
)

// content holds our static web server content.
//
//go:embed external/ui/html/*.tmpl.html external/ui/static/* external/ui/html/pages/*.tmpl.html external/ui/html/partials/**/*.tmpl.html external/ui/html/partials/*.tmpl.html
var content embed.FS

const serverPort = ":4044"

func main() {

	tempRouterMiddlewares := []mux.MiddlewareFunc{}

	// Initialise router
	httpRouter := router.NewRouter(response.GetResourceNotFoundError, response.GetDefault200Response, tempRouterMiddlewares...)

	// Prep web detail
	NewWebDetail(httpRouter, "ghat(d)", "https://ghatd.com", "example@ghatd.com", "ghatd ltd", content, "external/")

	// Define server
	srv := &http.Server{
		Addr:    serverPort,
		Handler: httpRouter.GetRouter(),
	}

	// Start listening
	fmt.Printf("\nServer is listening on port - %s\n", serverPort)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Default().Panicf("server/unable-to-start-server - %v", err)
	}
}

func NewWebDetail(httpRouter *router.Router, serviceExternalName, serviceExternalWebsite, serviceExternalEmail, serviceLegalBusinessName string, embeddedContent fs.FS, embeddedContentFilePathPrefix string) {

	//>ghatd {{ define "WebDetailInit" }}

	// Generate policies for web app
	termsOfServicePolicy := saasPolicy.NewGeneratedTermsPolicy(&saasPolicy.NewGeneratedTermsPolicyRequest{
		ServiceName:       serviceExternalName,
		ServiceWebsite:    serviceExternalWebsite,
		ServiceEmail:      serviceExternalEmail,
		LegalBusinessName: serviceLegalBusinessName,
	})

	privacyPolicy := saasPolicy.NewGeneratedPrivacyPolicy(&saasPolicy.NewGeneratedPrivacyPolicyRequest{
		ServiceName:       serviceExternalName,
		ServiceWebsite:    serviceExternalWebsite,
		ServiceEmail:      serviceExternalEmail,
		LegalBusinessName: serviceLegalBusinessName,
	})

	cookiePolicy := saasPolicy.NewGeneratedCookiePolicy(&saasPolicy.NewGeneratedCookiePolicyRequest{
		ServiceName:       serviceExternalName,
		ServiceWebsite:    serviceExternalWebsite,
		ServiceEmail:      serviceExternalEmail,
		LegalBusinessName: serviceLegalBusinessName,
	})

	// Initialise handler for web app
	webAppHandler := webapp.NewWebAppHandler(&webapp.NewWebAppHandlerRequest{
		EmbeddedContent:               embeddedContent,
		EmbeddedContentFilePathPrefix: embeddedContentFilePathPrefix,
		TermsOfServicePolicy:          termsOfServicePolicy,
		PrivacyPolicy:                 privacyPolicy,
		CookiePolicy:                  cookiePolicy,
	})

	// Attach routes
	webapp.AttachRoutes(&webapp.AttachRoutesRequest{
		Router:                        httpRouter,
		Handler:                       webAppHandler,
		WebAppFileSystem:              embeddedContent,
		EmbeddedContentFilePathPrefix: embeddedContentFilePathPrefix,
	})

	//>ghatd {{ end }}
}
