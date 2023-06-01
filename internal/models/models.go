package models

import "time"

// User is the user model
type User struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	AccessLevel int       `json:"access_level"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	LastLogin   time.Time `json:"last_login"`
}

// Room is the rooms model
type Room struct {
	ID        int       `json:"id"`
	RoomName  string    `json:"room_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Reservation is the reservation model
type Reservation struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"password"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Room      Room
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Restriction holds restrictions model
type Restriction struct {
	ID              int       `json:"id"`
	RestrictionName string    `json:"restriction_name"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// RoomRestriction holds room_restrctions table model
type RoomRestriction struct {
	ID            int       `json:"id"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	RestrictionID int       `json:"restriction_id"`
	ReservationID int       `json:"reservation_id"`
	RoomID        int       `json:"room_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
