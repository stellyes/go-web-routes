package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":4000"

// Every function that interracts with the web browser
// needs to take in an http response writer and pointer
// to a request
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValue(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

// Uppercase function names make them public
// to other packages beyond this one
func AddValue(x, y int) int {
	return x + y
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
