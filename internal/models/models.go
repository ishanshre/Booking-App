package models

// It holds reservation data
type Reservation struct {
	FristName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
