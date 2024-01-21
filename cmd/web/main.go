package main

import (
	"fmt"
	"github.com/stellyes/go-web-routes/pkg/config"
	"github.com/stellyes/go-web-routes/pkg/handlers"
	"github.com/stellyes/go-web-routes/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":4000"

// Lowercase function names make them private
// to this specific module
func main() {
	var app config.AppConfig

	// Templates are HTML fragments, so to speak
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	_, _ = fmt.Println(fmt.Sprintf("Starting applicatiton on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
