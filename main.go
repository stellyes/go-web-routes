package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":4000"

// Every function that interracts with the web browser
// needs to take in an http response writer and pointer
// to a request

// Uppercase function names make them public
// to other packages beyond this one
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// Lowercase function names make them private
// to this specific module
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_, _ = fmt.Println(fmt.Sprintf("Starting applicatiton on port %s", portNumber))

	// Start web server (port, and handler)
	_ = http.ListenAndServe(portNumber, nil)
}
