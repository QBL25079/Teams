package user_repository

import (
	"context"
	"fmt"

	"github.com/QBL25079/teams/internal/core/domain"
)

func (r *UserRepository) UpdateUser(ctx context.Context, id int, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `UPDATE teams.user SET first_name = $1, last_name = $2, team_id = $3, updated_at = NOW() WHERE id = $4 RETURNING id, first_name, last_name, birth_year, team_id, created_at, updated_at;`

	var userModel UserModel

	err := r.pool.QueryRow(ctx, query, user.FirstName, user.LastName, user.TeamID, id).Scan(&userModel.ID, &userModel.FirstName, &userModel.LastName, &userModel.BirthYear, &userModel.TeamID, &userModel.CreatedAt, &userModel.UpdatedAt)

	if err != nil {
		return domain.User{}, fmt.Errorf("patch user: %w", err)
	}

	return domain.NewUser(userModel.ID, userModel.FirstName, userModel.LastName, userModel.BirthYear, userModel.TeamID, userModel.CreatedAt, userModel.UpdatedAt), nil
}
