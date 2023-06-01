package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ishanshre/Booking-App/internal/config"
	"github.com/ishanshre/Booking-App/internal/handler"
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
	mux.Post("/make-reservation", handler.Repo.HandlePostMakeReservation)

	mux.Get("/reservation-summary", handler.Repo.HandleReservSummary)

	mux.Get("/search-avaliable", handler.Repo.HandleSearchAvaliable)
	mux.Post("/search-avaliable", handler.Repo.HandlePostSearchAvaliable)
	mux.Post("/search-avaliable-json", handler.Repo.HandleSearchAvaliableJson)

	mux.Get("/choose-room/{id}", handler.Repo.HandleChooseRoom)

	mux.Get("/user/login", handler.Repo.HandleLogin)
	mux.Post("/user/login", handler.Repo.HandlePostLogin)
	mux.Get("/user/logout", handler.Repo.HandleLogout)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
