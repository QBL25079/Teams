package user_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/QBL25079/teams/internal/core/domain"
	errrors "github.com/QBL25079/teams/internal/core/errors"
	postgres_pool "github.com/QBL25079/teams/internal/core/repository/postgres/pool"
)

func (r *UserRepository) GetUser(ctx context.Context, id int) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())

	defer cancel()

	query := `SELECT id, first_name, last_name, birth_year, team_id, created_at, updated_at FROM teams.user WHERE id = $1;`

	row := r.pool.QueryRow(ctx, query, id)

	var userModel UserModel
	err := row.Scan(&userModel.ID, &userModel.FirstName, &userModel.LastName, &userModel.BirthYear, &userModel.GroupID, &userModel.CreatedAt, &userModel.UpdatedAt)
	if err != nil {
		if errors.Is(err, postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with id='%d': %w", id, errrors.ErrNotFound)
		}
		return domain.User{}, fmt.Errorf("scan error %w", err)
	}
	userDomain := domain.NewUser(userModel.ID, userModel.FirstName, userModel.LastName, userModel.BirthYear, userModel.GroupID, userModel.CreatedAt, userModel.UpdatedAt)
	return userDomain, nil
}
