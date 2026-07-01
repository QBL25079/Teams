package team_repository

import (
	"context"
	"fmt"

	"github.com/QBL25079/teams/internal/core/domain"
)

func (r *TeamRepository) UpdateTeam(ctx context.Context, id int, team domain.Team) (domain.Team, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		UPDATE teams.team
		SET name = $1,
		    parent_id = $2,
		    updated_at = NOW()
		WHERE id = $3
		RETURNING id, name, parent_id, created_at, updated_at
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		team.Name,
		team.ParentID,
		id,
	)

	var teamModel TeamModel

	err := row.Scan(
		&teamModel.ID,
		&teamModel.Name,
		&teamModel.ParentID,
		&teamModel.CreatedAt,
		&teamModel.UpdatedAt,
	)

	if err != nil {
		return domain.Team{}, fmt.Errorf("update team: %w", err)
	}

	return domain.Team{
		ID:        teamModel.ID,
		Name:      teamModel.Name,
		ParentID:  teamModel.ParentID,
		CreatedAt: teamModel.CreatedAt,
		UpdatedAt: teamModel.UpdatedAt,
	}, nil
}
