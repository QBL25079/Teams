package user_transport_http

import "time"

type DTOUserResponse struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthYear int       `json:"birth_year"`
	GroupID   int       `json:"group_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
