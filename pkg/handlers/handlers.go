package handlers

import (
	"net/http"

	"github.com/stellyes/go-web-routes/pkg/config"
	"github.com/stellyes/go-web-routes/pkg/models"
	"github.com/stellyes/go-web-routes/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// Creates new repository~
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Every function that interracts with the web browser
// needs to take in an http response writer and pointer
// to a request

// Uppercase function names make them public
// to other packages beyond this one
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// m becomes the reciever of everything made available by
// Repository struct
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "This feature is very similar to handlebars."

	// Send data to template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
