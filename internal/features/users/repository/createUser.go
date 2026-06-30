package user_repository

import (
	"context"
	"fmt"

	core_domain "github.com/QBL25079/teams/internal/core/domain"
)

func (r *UserRepository) CreateUser(ctx context.Context, user core_domain.User) (core_domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO teams.users (first_name, last_name, birth_year, group_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, first_name, last_name, birth_year, group_id, created_at, updated_at;
	`

	row := r.pool.QueryRow(ctx, query,
		user.FirstName,
		user.LastName,
		user.BirthYear,
		user.GroupID,
	)

	var userModel UserModel

	err := row.Scan(
		&userModel.ID,
		&userModel.FirstName,
		&userModel.LastName,
		&userModel.BirthYear,
		&userModel.GroupID,
		&userModel.CreatedAt,
		&userModel.UpdatedAt,
	)
	if err != nil {
		return core_domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := core_domain.NewUser(
		userModel.ID,
		userModel.FirstName,
		userModel.LastName,
		userModel.BirthYear,
		userModel.GroupID,
		userModel.CreatedAt,
		userModel.UpdatedAt,
	)

	return userDomain, nil
}
