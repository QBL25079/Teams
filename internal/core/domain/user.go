package core_domain

import (
	"fmt"
	"time"

	core_errors "github.com/QBL25079/teams/internal/core/errors"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	BirthYear int
	GroupID   *int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(ID, BirthYear int, FirstName, LastName string, CreatedAt, UpdatedAt time.Time,  GroupID *int) User {
	return User{ID: ID, FirstName: FirstName, LastName: LastName, BirthYear: BirthYear, GroupID: GroupID, CreatedAt: CreatedAt, UpdatedAt: UpdatedAt}
}

func (u *User) Validate() error {
	firstNameLen := len([]rune(u.FirstName))
	if firstNameLen < 1 || firstNameLen > 100 {
		return fmt.Errorf("invalid first_name length %d: %w", firstNameLen, core_errors.ErrInvalidArgument)
	}

	lastNameLen := len([]rune(u.LastName))
	if lastNameLen < 1 || lastNameLen > 100 {
		return fmt.Errorf("invalid last_name length %d: %w", lastNameLen, core_errors.ErrInvalidArgument)
	}

	currentYear := time.Now().Year()
	if u.BirthYear < 1900 || u.BirthYear > currentYear {
		return fmt.Errorf("invalid birth_year %d: %w", u.BirthYear, core_errors.ErrInvalidArgument)
	}

	if u.GroupID != nil {
		if *u.GroupID <= 0 {
			return fmt.Errorf("invalid group_id %d: %w", *u.GroupID, core_errors.ErrInvalidArgument)
		}
	}

	return nil
}