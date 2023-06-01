package dbrepo

import (
	"context"
	"errors"
	"time"

	"github.com/ishanshre/Booking-App/internal/models"
	"golang.org/x/crypto/bcrypt"
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
		res.Room.ID,
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

// returns true if availability exists else returns false
func (m *postgresDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `
		SELECT 
			COUNT(id)
		FROM
			room_restrictions
		WHERE
			room_id = $1
			$2 < end_date and $3 > start_date;
	`
	var numRows int
	row := m.DB.QueryRowContext(ctx, query, roomID, start, end)
	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}
	if numRows == 0 {
		return true, nil
	}
	return false, nil
}

// Returns slic of rooms avalibale in given dates
func (m *postgresDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]*models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var rooms []*models.Room
	query := `
		SELECT
			r.id, r.room_name
		FROM
			rooms as r
		WHERE r.id not in
			(SELECT 
				room_id 
			FROM room_restrictions as rr 
			WHERE $1 < rr.end_date and $2 > rr.start_date
			)
	`
	rows, err := m.DB.QueryContext(
		ctx,
		query,
		start,
		end,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		room := models.Room{}
		err := rows.Scan(
			&room.ID,
			&room.RoomName,
		)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}
	return rooms, nil
}

// Get Room By Id
func (m *postgresDBRepo) GetRoomByID(id int) (*models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	room := &models.Room{}
	query := `
		SELECT *
		FROM rooms 
		WHERE id=$1
	`
	rows := m.DB.QueryRowContext(ctx, query, id)
	if err := rows.Scan(
		&room.ID,
		&room.RoomName,
		&room.CreatedAt,
		&room.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return room, nil
}

func (m *postgresDBRepo) GetUserById(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `
		SELECT id, first_name, last_name, email, password, access_level, created_at, updated_at, last_login
		FROM users
		WHERE id=$1
	`
	row := m.DB.QueryRowContext(ctx, query, id)
	var u *models.User
	err := row.Scan(
		u.ID,
		u.FirstName,
		u.LastName,
		u.Email,
		u.Password,
		u.AccessLevel,
		u.CreatedAt,
		u.UpdatedAt,
		u.LastLogin,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (m *postgresDBRepo) UpdateUser(u *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `
		UPDATE users
		SET first_name = $2, last_name = $3, email = $4, access_level = $5, updated_at = $6
		WHERE id = $1
	`
	_, err := m.DB.ExecContext(
		ctx,
		stmt,
		u.ID,
		u.FirstName,
		u.LastName,
		u.Email,
		u.AccessLevel,
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

// Authenticates a user
func (m *postgresDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int
	var hashedPassword string
	query := `SELECT id, password FROM users WHERE email=$1`
	row := m.DB.QueryRowContext(ctx, query, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return id, "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("incorrect password")
	} else if err != nil {
		return 0, "", err
	}
	return id, hashedPassword, nil
}
