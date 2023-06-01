package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
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
	res, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(w, errors.New("cannot get the reservation from session"))
		return
	}
	room, err := m.DB.GetRoomByID(res.Room.ID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	res.Room.RoomName = room.RoomName
	m.App.Session.Put(r.Context(), "reservation", res)
	sd := res.StartDate.Format("2006-01-02")
	ed := res.EndDate.Format("2006-01-02")
	stringMap := make(map[string]string)
	stringMap["start_date"] = sd
	stringMap["end_date"] = ed

	// var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = res

	render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form:      forms.New(nil),
		Data:      data,
		StringMap: stringMap,
	})
}

// handles the post of the reservation page and display form
func (m *Repository) HandlePostMakeReservation(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		//log.Println(err)
		helpers.ServerError(w, err)
		return
	}

	// get the start and end date from the form
	sd := r.Form.Get("start_date")
	ed := r.Form.Get("end_date")
	// date we get 2021-01-01 format so parsing the date format
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, sd)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	endDate, err := time.Parse(layout, ed)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	roomID, err := strconv.Atoi(r.Form.Get("room_id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	reservation := &models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
		StartDate: startDate,
		EndDate:   endDate,
		Room:      models.Room{ID: roomID},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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
	newReservationID, err := m.DB.InsertReservation(reservation)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	restriction := &models.RoomRestriction{
		StartDate:     reservation.StartDate,
		EndDate:       reservation.EndDate,
		RoomID:        reservation.Room.ID,
		ReservationID: newReservationID,
		RestrictionID: 1,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := m.DB.InsertRoomRestrictions(restriction); err != nil {
		helpers.ServerError(w, err)
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
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, start)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	endDate, err := time.Parse(layout, end)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	rooms, err := m.DB.SearchAvailabilityForAllRooms(startDate, endDate)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	for _, i := range rooms {
		m.App.InfoLog.Println("Room: ", i.ID, i.RoomName)
	}
	if len(rooms) == 0 {
		// no availability
		m.App.Session.Put(r.Context(), "error", "No availablity")
		http.Redirect(w, r, "/search-avaliable", http.StatusSeeOther)
		return
	}
	data := make(map[string]interface{})
	data["rooms"] = rooms
	res := models.Reservation{
		StartDate: startDate,
		EndDate:   endDate,
	}
	m.App.Session.Put(r.Context(), "reservation", res)

	render.Template(w, r, "choose-room.page.tmpl", &models.TemplateData{
		Data: data,
	})
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

func (m *Repository) HandleChooseRoom(w http.ResponseWriter, r *http.Request) {
	roomID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	// get the reservation data from the session
	res, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(w, err)
		return
	}
	// assign roomID from url to the reservation model
	res.Room.ID = roomID
	// put back the reservation data into session that will consist extra roomID
	m.App.Session.Put(r.Context(), "reservation", res)

	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)
}
