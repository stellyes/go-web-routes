package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/stellyes/go-web-routes/pkg/config"
	"github.com/stellyes/go-web-routes/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// Creates http handler
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
