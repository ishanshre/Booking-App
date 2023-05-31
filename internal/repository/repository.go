package repository

import "github.com/ishanshre/Booking-App/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res *models.Reservation) error
}
