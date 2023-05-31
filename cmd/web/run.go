package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ishanshre/Booking-App/internal/driver"
	"github.com/ishanshre/Booking-App/internal/handler"
	"github.com/ishanshre/Booking-App/internal/helpers"
	"github.com/ishanshre/Booking-App/internal/models"
	"github.com/ishanshre/Booking-App/internal/render"
)

var infoLog *log.Logger
var errorLog *log.Logger

func run() (*driver.DB, error) {
	// store values in the sessions
	gob.Register(models.User{})
	gob.Register(models.Reservation{})
	gob.Register(models.Restriction{})
	gob.Register(models.Room{})
	gob.Register(models.RoomRestriction{})

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

	//connect to database
	log.Println("connecting to database")
	db, err := driver.ConnectSQL(os.Getenv("postgres"))
	if err != nil {
		log.Fatalln("Error in connecting to database", err)
		return nil, err
	}
	log.Println("connected to database")

	// initiate a template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Printf("Cannot create a template cache %s\n", err)
		return nil, err
	}
	// store the template in global template cache
	app.TemplateCache = tc
	app.UseCache = false // when false it will read from disk always

	// pass global config to render
	render.NewRenderer(&app)

	// pass global config to render
	repo := handler.NewRepo(&app, db)
	// pass the repo to the handler
	handler.NewHandler(repo)
	helpers.NewHelpers(&app)
	return db, nil
}
