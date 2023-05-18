package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ishanshre/Booking-App/pkg/config"
	"github.com/ishanshre/Booking-App/pkg/handler"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)      // csrf middleware
	mux.Use(SessionLoad) // session load middleware
	mux.Get("/", handler.Repo.HandleHome)
	mux.Get("/about", handler.Repo.HandleAbout)
	mux.Get("/contact", handler.Repo.HandleContact)
	mux.Get("/generals", handler.Repo.HandleGenerals)
	mux.Get("/majors", handler.Repo.HandleMajors)
	mux.Get("/make-reservation", handler.Repo.HandleMakeReservation)
	mux.Get("/reservation-summary", handler.Repo.HandleReservSummary)
	mux.Get("/search-avaliable", handler.Repo.HandleSearchAvaliable)
	mux.Post("/search-avaliable", handler.Repo.HandlePostSearchAvaliable)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
