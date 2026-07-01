package team_repository

import (
	"context"
	"fmt"

	"github.com/QBL25079/teams/internal/core/domain"
)

func (r *TeamRepository) GetTeams(ctx context.Context, limit, offset *int) ([]domain.Team, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT id, name, parent_id, created_at, updated_at FROM teams.team ORDER BY id ASC LIMIT $1 OFFSET $2;`

	rows, err := r.pool.Query(ctx, query, limit, offset)

	if err != nil {
		return nil, fmt.Errorf("Error to get teams from repo: %w", err)
	}

	defer rows.Close()

	var teamModels []TeamModel

	for rows.Next() {
		var teamModel TeamModel

		if err := rows.Scan(
			&teamModel.ID,
			&teamModel.Name,
			&teamModel.ParentID,
			&teamModel.CreatedAt,
			&teamModel.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("Scan teams error: %w", err)
		}

		teamModels = append(teamModels, teamModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	teamDomains := teamDomainsFromModels(teamModels)

	return teamDomains, nil
}