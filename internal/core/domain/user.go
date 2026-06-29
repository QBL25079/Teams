package core_domain

import "time"

type User struct {
	ID        int
	FirstName string
	LastName  string
	BirthYear int
	GroupID   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(ID, BirthYear, GroupID int, FirstName, LastName string, CreatedAt, UpdatedAt time.Time) User {
	return User{ID: ID, FirstName: FirstName, LastName: LastName, BirthYear: BirthYear, GroupID: GroupID, CreatedAt: CreatedAt, UpdatedAt: UpdatedAt}
}
