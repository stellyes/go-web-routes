package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	// Returns anonymous function, casts to handler function,
	// serving our middlware to client
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("We hit the page!")
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",   // Referencing the "entire" site
		Secure:   false, // For HTTPS, which I'm not using rn
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}
