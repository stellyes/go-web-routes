package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":4000"

// Lowercase function names make them private
// to this specific module
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_, _ = fmt.Println(fmt.Sprintf("Starting applicatiton on port %s", portNumber))

	// Start web server (port, and handler)
	_ = http.ListenAndServe(portNumber, nil)
}
