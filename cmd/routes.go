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
	mux.Get("/quartos/quartel", http.HandlerFunc(h.Repo.RoomQuartel))
	mux.Get("/quartos/suite", http.HandlerFunc(h.Repo.RoomSuite))
	mux.Get("/contato", http.HandlerFunc(h.Repo.Contact))
	mux.Get("/disponivel", http.HandlerFunc(h.Repo.Disponible))
	mux.Get("/fazer-reserva", http.HandlerFunc(h.Repo.DoReservation))

	// not found
	mux.NotFound(h.Repo.NotFound)

	// enable static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
