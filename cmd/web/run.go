package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ishanshre/Booking-App/internal/handler"
	"github.com/ishanshre/Booking-App/internal/helpers"
	"github.com/ishanshre/Booking-App/internal/models"
	"github.com/ishanshre/Booking-App/internal/render"
)

var infoLog *log.Logger
var errorLog *log.Logger

func run() error {
	// store values in the sessions
	gob.Register(models.Reservation{})

	// set InProduction to true in production
	app.InProduction = false

	// setup logger
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// initiating a session
	session = scs.New()               // creating a new session
	session.Lifetime = 24 * time.Hour // for more securty we can set it to 30 minutes
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	// initiate a template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Printf("Cannot create a template cache %s\n", err)
		return err
	}
	// store the template in global template cache
	app.TemplateCache = tc
	app.UseCache = false // when false it will read from disk always

	// pass global config to render
	render.NewTemplate(&app)

	// pass global config to render
	repo := handler.NewRepo(&app)
	// pass the repo to the handler
	handler.NewHandler(repo)
	helpers.NewHelpers(&app)
	return nil
}
