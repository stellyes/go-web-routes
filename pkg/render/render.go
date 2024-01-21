package render

import (
	"bytes"
	"github.com/stellyes/go-web-routes/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// get the template cache from app config
	tc := app.TemplateCache

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get tmpl from template cache")
	}

	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{} // Map of strings pointing to templates

	// get all files with name *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	// Sift through files with *.page.tmpl extension
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		// ... attach reference to base layout for each page template
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return myCache, err
			}
		}

		// Add page to template cache
		myCache[name] = ts
	}

	// Return final cache and /nil/ error
	return myCache, nil
}

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	// Check to see if template is already in the template cache (tc)
// 	_, inMap := tc[t]

// 	if !inMap {
// 		//load template into map
// 		log.Println("Creating template and adding to cache")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		//template exists
// 		log.Println("Using cached template")
// 	}

// 	tmpl = tc[t]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	// Create slice with template string and base layout
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}

// 	tmpl, err := template.ParseFiles(templates...)

// 	if err != nil {
// 		return err
// 	}

// 	// Add template to cache
// 	tc[t] = tmpl

// 	return nil
// }
