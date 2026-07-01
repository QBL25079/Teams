package user_repository

import (
	"context"
	"fmt"

	"github.com/QBL25079/teams/internal/core/domain"
)

func (r *UserRepository) GetUsers(ctx context.Context, limit, offset, teamID *int) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT id, first_name, last_name, birth_year, team_id, created_at, updated_at 
	          FROM teams.user 
	          WHERE ($1::int IS NULL OR team_id = $1)
	          ORDER BY id ASC 
	          LIMIT $2 OFFSET $3;`

	rows, err := r.pool.Query(ctx, query, teamID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("Error to get users from repo: %w", err)
	}
	defer rows.Close()

	var userModels []UserModel

	for rows.Next() {
		var userModel UserModel
		if err := rows.Scan(
			&userModel.ID,
			&userModel.FirstName,
			&userModel.LastName,
			&userModel.BirthYear,
			&userModel.TeamID,
			&userModel.CreatedAt,
			&userModel.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("Scan users error: %w", err)
		}
		userModels = append(userModels, userModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	return userDomainsFromModels(userModels), nil
}
