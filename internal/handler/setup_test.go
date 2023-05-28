package handler

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ishanshre/Booking-App/internal/config"
	"github.com/ishanshre/Booking-App/internal/models"
	"github.com/ishanshre/Booking-App/internal/render"
	"github.com/justinas/nosurf"
)

var app config.AppConfig
var session *scs.SessionManager
var pathToTemplates string = "./../../templates"

// var functions = template.FuncMap{}

func getRoutes() http.Handler {
	gob.Register(models.Reservation{})

	// set InProduction to true in production
	app.InProduction = false

	// initiating a session
	session = scs.New()               // creating a new session
	session.Lifetime = 24 * time.Hour // for more securty we can set it to 30 minutes
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	// initiate a template cache
	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Printf("Cannot create a template cache %s\n", err)
	}
	// store the template in global template cache
	app.TemplateCache = tc
	app.UseCache = true // when false it will read from disk always

	// pass global config to render

	// pass global config to render
	repo := NewRepo(&app)
	// pass the repo to the handler
	NewHandler(repo)
	render.NewTemplate(&app)
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	//mux.Use(NoSurf)      // csrf middleware
	mux.Use(SessionLoad) // session load middleware
	mux.Get("/", Repo.HandleHome)
	mux.Get("/about", Repo.HandleAbout)
	mux.Get("/contact", Repo.HandleContact)
	mux.Get("/generals", Repo.HandleGenerals)
	mux.Get("/majors", Repo.HandleMajors)

	mux.Get("/make-reservation", Repo.HandleMakeReservation)
	mux.Post("/make-reservation", Repo.HandlePostMakeReservation)

	mux.Get("/reservation-summary", Repo.HandleReservSummary)

	mux.Get("/search-avaliable", Repo.HandleSearchAvaliable)
	mux.Post("/search-avaliable", Repo.HandlePostSearchAvaliable)
	mux.Post("/search-avaliable-json", Repo.HandleSearchAvaliableJson)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}

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

func CreateTestTemplateCache() (map[string]*template.Template, error) {

	// path patterns for layout and pages templates
	pathLayoutPattern := filepath.Join(pathToTemplates, "layout", "*.layout.tmpl")
	pathPagePattern := filepath.Join(pathToTemplates, "pages", "*.page.tmpl")
	myCache := map[string]*template.Template{}

	// filepath.Glob return a slice with all *.page.tmpl in template/pages folder
	pages, err := filepath.Glob(pathPagePattern)
	if err != nil {
		return myCache, err
	}
	// loop through the pages and add base templates
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, fmt.Errorf("error in parsing template: %s", err)
		}

		// find the base layout file and
		matches, err := filepath.Glob(pathLayoutPattern)
		if err != nil {
			return myCache, err
		}
		// if found
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(pathLayoutPattern)
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
