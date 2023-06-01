package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/ishanshre/Booking-App/internal/models"
)

type AppConfig struct {
	UseCache      bool
	InProduction  bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	Session       *scs.SessionManager
	MailChan      chan models.MailData
}
