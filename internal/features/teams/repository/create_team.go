package team_repository

import (
	"context"
	"fmt"

	"github.com/QBL25079/teams/internal/core/domain"
)

func (r *TeamRepository) CreateTeam(ctx context.Context, team domain.Team) (domain.Team, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO teams.team (name, parent_id)
		VALUES ($1, $2)
		RETURNING id, name, parent_id, created_at, updated_at;
	`

	row := r.pool.QueryRow(ctx, query, team.Name, team.ParentID)

	var teamModel TeamModel

	err := row.Scan(
		&teamModel.ID,
		&teamModel.Name,
		&teamModel.ParentID,
		&teamModel.CreatedAt,
		&teamModel.UpdatedAt,
	)
	if err != nil {
		return domain.Team{}, fmt.Errorf("scan error: %w", err)
	}

	teamDomain := domain.NewTeam(
		teamModel.ID,
		teamModel.Name,
		teamModel.ParentID,
		teamModel.CreatedAt,
		teamModel.UpdatedAt,
	)

	return teamDomain, nil
}
