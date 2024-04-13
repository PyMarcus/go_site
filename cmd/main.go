package main

import (
	"log"
	"net/http"
	"time"

	"github.com/PyMarcus/go_site/config"
	h "github.com/PyMarcus/go_site/handlers"
	"github.com/alexedwards/scs/v2"
)

var app *config.AppConfig
var session *scs.SessionManager

func main() {
	const addr string = ":8080"

	// change to true in production (https)
	app = &config.AppConfig{Production: false}
	repo := h.CreateNewRepository(app)
	h.NewHandlers(repo)
	setSession()

	srv := &http.Server{
		Addr:    addr,
		Handler: router(app),
	}

	log.Println("Running at", addr)
	log.Fatal(srv.ListenAndServe())
}

// setSession configurations
func setSession() {
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Production // https
	app.Session = session
}
