package main

import (
	"net/http"

	"github.com/PyMarcus/go_site/config"
	h "github.com/PyMarcus/go_site/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func router(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", http.HandlerFunc(h.Repo.Index))
	mux.Get("/sobre", http.HandlerFunc(h.Repo.About))

	// enable static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
