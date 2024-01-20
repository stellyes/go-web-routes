package main

import (
	"errors"
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
	fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

// Uppercase function names make them public
// to other packages beyond this one
func AddValue(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("%s", err))
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f / %f = %f", 100.0, 10.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by zero!")
		return 0, err
	}
	result := x / y
	return result, nil
}

// Lowercase function names make them private
// to this specific module
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	_, _ = fmt.Println(fmt.Sprintf("Starting applicatiton on port %s", portNumber))

	// Start web server (port, and handler)
	_ = http.ListenAndServe(portNumber, nil)
}
