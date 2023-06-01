package repository

import (
	"time"

	"github.com/ishanshre/Booking-App/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(*models.Reservation) (int, error)
	InsertRoomRestrictions(*models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(time.Time, time.Time, int) (bool, error)
	SearchAvailabilityForAllRooms(time.Time, time.Time) ([]*models.Room, error)
	GetRoomByID(int) (*models.Room, error)
}
