package models

import "github.com/ishanshre/Booking-App/internal/forms"

type TemplateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	FlashMessage    string
	Warning         string
	Error           string
	Form            *forms.Form
	IsAuthenticated int
}
