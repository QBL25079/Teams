package domain

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
	TeamID    *int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserPatch struct {
	FirstName *string
	LastName  *string
	TeamID    *int
}

func NewUserPatch(firstName, lastName *string, teamID *int) UserPatch {
	return UserPatch{
		FirstName: firstName,
		LastName:  lastName,
		TeamID:    teamID,
	}
}

func NewUser(ID int, FirstName, LastName string, BirthYear int, TeamID *int, CreatedAt, UpdatedAt time.Time) User {
	return User{ID: ID, FirstName: FirstName, LastName: LastName, BirthYear: BirthYear, TeamID: TeamID, CreatedAt: CreatedAt, UpdatedAt: UpdatedAt}
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

	if u.TeamID != nil {
		if *u.TeamID <= 0 {
			return fmt.Errorf("invalid group_id %d: %w", *u.TeamID, core_errors.ErrInvalidArgument)
		}
	}

	return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
	tmp := *u

	if patch.FirstName != nil {
		tmp.FirstName = *patch.FirstName
	}

	if patch.LastName != nil {
		tmp.LastName = *patch.LastName
	}

	if patch.TeamID != nil {
		tmp.TeamID = patch.TeamID
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate updated user: %w", err)
	}

	*u = tmp
	return nil
}
