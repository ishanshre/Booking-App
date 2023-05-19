package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ishanshre/Booking-App/pkg/config"
	"github.com/ishanshre/Booking-App/pkg/models"
	"github.com/ishanshre/Booking-App/pkg/render"
)

// create a Repository type to get use global config
type Repository struct {
	App *config.AppConfig
}

// Repo to use by handlers
var Repo *Repository

// NewRepo create a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) HandleHome(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) HandleAbout(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) HandleContact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) HandleGenerals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

func (m *Repository) HandleMajors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

func (m *Repository) HandleMakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

func (m *Repository) HandleReservSummary(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "reservation-summary.page.tmpl", &models.TemplateData{})
}

func (m *Repository) HandleSearchAvaliable(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-avaliable.page.tmpl", &models.TemplateData{})
}

func (m *Repository) HandlePostSearchAvaliable(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start-date")
	end := r.Form.Get("end-date")
	w.Write([]byte(fmt.Sprintf("Post method to search avaliablility %s and %s", start, end)))
}

func (m *Repository) HandleSearchAvaliableJson(w http.ResponseWriter, r *http.Request) {
	log.Println("working")
	resp := models.JsonResponse{
		Ok:      true,
		Message: "Avaliable",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
