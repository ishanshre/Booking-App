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
	return mux
}
