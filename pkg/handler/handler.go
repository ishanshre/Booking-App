package handler

import (
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
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) HandleAbout(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
