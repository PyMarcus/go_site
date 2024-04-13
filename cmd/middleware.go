package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf is a middleware that block CSRF attack
func NoSurf(next http.Handler) http.Handler {

	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.Production, // https
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad saves the session provides middleware which automatically
// loads and saves session data for the current request, and communicates
// the session token to and from the client in a cookie.
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
