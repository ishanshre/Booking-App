package main

import (
	"testing"

	"github.com/go-chi/chi"
	"github.com/ishanshre/Booking-App/internal/config"
)

func TestRouter(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)
	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing
	default:
		t.Errorf("%T; type is not chi.MUX", v)
	}
}
