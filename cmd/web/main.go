package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/alexedwards/scs/v2"
	"github.com/ishanshre/Booking-App/internal/config"
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

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

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
