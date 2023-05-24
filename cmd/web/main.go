package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ishanshre/Booking-App/internal/config"
	"github.com/ishanshre/Booking-App/internal/handler"
	"github.com/ishanshre/Booking-App/internal/models"
	"github.com/ishanshre/Booking-App/internal/render"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usuage: command <port-number>")
	}
	p, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Println("Usuage: command <port-number>")
		log.Fatalln("port number must be an integer: 1024 to 65535")
	}

	// store values in the sessions
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
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Printf("Cannot create a template cache %s\n", err)
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

	// create the http server
	port := fmt.Sprintf(":%v", p)
	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	log.Println("starting the server at port :", p)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}

}
