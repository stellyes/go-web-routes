package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tc = make(map[string]*template.Template)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// Check to see if template is already in the template cache (tc)
	_, inMap := tc[t]

	if !inMap {
		//load template into map
		log.Println("Creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		//template exists
		log.Println("Using cached template")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	// Create slice with template string and base layout
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	// Add template to cache
	tc[t] = tmpl

	return nil
}
