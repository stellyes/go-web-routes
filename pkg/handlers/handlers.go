package handlers

import (
	"github.com/stellyes/go-web-routes/pkg/render"
	"net/http"
)

// Every function that interracts with the web browser
// needs to take in an http response writer and pointer
// to a request

// Uppercase function names make them public
// to other packages beyond this one
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
