package main

import (
	"log"
	"time"

	"github.com/ishanshre/Booking-App/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

// fires a go routinse
func listenForMail() {
	go func() {
		for {
			msg := <-app.MailChan
			sendMsg(msg)
		}
	}()
}

func sendMsg(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second
	// for production extra config needed such as username, password and encryption
	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
	}
	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	email.SetBody(mail.TextHTML, "Hello, <strong>world</strong>!")
	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("working")
	}
}
