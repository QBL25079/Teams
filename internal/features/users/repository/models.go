package user_repository

import (
	"time"

	"github.com/QBL25079/teams/internal/core/domain"
)

type UserModel struct {
	ID        int
	FirstName string
	LastName  string
	BirthYear int
	TeamID   *int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func userDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))

	for i, user := range users {
		userDomains[i] = domain.NewUser(
			user.ID,
			user.FirstName,
			user.LastName,
			user.BirthYear,
			user.TeamID,
			user.CreatedAt,
			user.UpdatedAt,
		)
	}
	return userDomains
}
