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
	GroupID   *int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func userDomainsFromModels(users []UserModel) []core_domain.User {
	userDomains := make([]core_domain.User, len(users))

	for i, user := range users {
		userDomains[i] = core_domain.NewUser(
			user.ID,
			user.FirstName,
			user.LastName,
			user.BirthYear,
			user.GroupID,
			user.CreatedAt,
			user.UpdatedAt,
		)
	}
	return userDomains
}
