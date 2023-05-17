package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf implement csrf token
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next) // create a handler
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode, // lax mode allows cokkies to be sent in cross site
	})
	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
