package user_transport_http

import (
	"time"

	core_domain "github.com/QBL25079/teams/internal/core/domain"
)

type DTOUserResponse struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthYear int       `json:"birth_year"`
	TeamID    *int      `json:"team_id" validate:"omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func UserDTOFromDomain(user core_domain.User) DTOUserResponse {
	return DTOUserResponse{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName, BirthYear: user.BirthYear, TeamID: user.TeamID, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt}
}

func UsersDTOFromDomains(users []core_domain.User) []DTOUserResponse {
	if users == nil {
		return nil
	}
	userDTOs := make([]DTOUserResponse, len(users))

	for i, user := range users {
		userDTOs[i] = UserDTOFromDomain(user)
	}

	return userDTOs
}
