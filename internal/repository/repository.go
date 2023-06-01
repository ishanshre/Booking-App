package repository

import "github.com/ishanshre/Booking-App/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(*models.Reservation) (int, error)
	InsertRoomRestrictions(*models.RoomRestriction) error
}
