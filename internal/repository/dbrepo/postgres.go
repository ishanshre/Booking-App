package dbrepo

import (
	"context"
	"time"

	"github.com/ishanshre/Booking-App/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

func (m *postgresDBRepo) InsertReservation(res *models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var newID int
	stmt := `
		INSERT INTO reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING id
	`
	err := m.DB.QueryRowContext(
		ctx,
		stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		res.CreatedAt,
		res.UpdatedAt,
	).Scan(&newID)
	if err != nil {
		return 0, err
	}
	return newID, nil
}

// InsertRoomRestrictions insert row to room_restrictions in database
func (m *postgresDBRepo) InsertRoomRestrictions(res *models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `
		INSERT INTO room_restrictions (start_date, end_date, restriction_id, reservation_id, room_id, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
	`
	_, err := m.DB.ExecContext(ctx, stmt,
		res.StartDate,
		res.EndDate,
		res.RestrictionID,
		res.ReservationID,
		res.RoomID,
		res.CreatedAt,
		res.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
