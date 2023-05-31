package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ishanshre/Booking-App/internal/config"
	"github.com/ishanshre/Booking-App/internal/driver"
	"github.com/ishanshre/Booking-App/internal/forms"
	"github.com/ishanshre/Booking-App/internal/helpers"
	"github.com/ishanshre/Booking-App/internal/models"
	"github.com/ishanshre/Booking-App/internal/render"
	"github.com/ishanshre/Booking-App/internal/repository"
	"github.com/ishanshre/Booking-App/internal/repository/dbrepo"
)

// create a Repository type to get use global config
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// Repo to use by handlers
var Repo *Repository

// NewRepo create a new Repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) HandleHome(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) HandleAbout(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "about.page.tmpl", &models.TemplateData{})
}

func (m *Repository) HandleContact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) HandleGenerals(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "generals.page.tmpl", &models.TemplateData{})
}

func (m *Repository) HandleMajors(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// HandlePostMakeReservation handle the post method
func (m *Repository) HandleMakeReservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// handles the post of the reservation page and display form
func (m *Repository) HandlePostMakeReservation(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		//log.Println(err)
		helpers.ServerError(w, err)
		return
	}
	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}
	form := forms.New(r.PostForm)
	form.Required("first_name", "last_name", "email")
	if reservation.Phone != "" {
		form.MinLength("phone", 10)
	}
	form.IsEmail("email")
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (m *Repository) HandleReservSummary(w http.ResponseWriter, r *http.Request) {
	// get reservation from session and type cast into models.Reservation
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("Cannot get error from the session")
		m.App.Session.Put(r.Context(), "error", "Cannot get the reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation
	render.Template(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

func (m *Repository) HandleSearchAvaliable(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "search-avaliable.page.tmpl", &models.TemplateData{})
}

func (m *Repository) HandlePostSearchAvaliable(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start-date")
	end := r.Form.Get("end-date")
	w.Write([]byte(fmt.Sprintf("Post method to search avaliablility %s and %s", start, end)))
}

func (m *Repository) HandleSearchAvaliableJson(w http.ResponseWriter, r *http.Request) {
	resp := models.JsonResponse{
		Ok:      true,
		Message: "Avaliable",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
