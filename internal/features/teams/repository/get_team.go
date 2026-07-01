package team_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/QBL25079/teams/internal/core/domain"
	errrors "github.com/QBL25079/teams/internal/core/errors"
	postgres_pool "github.com/QBL25079/teams/internal/core/repository/postgres/pool"
)

func (r *TeamRepository) GetTeam(ctx context.Context, taskID int) (domain.Team, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT id, name, parent_id, created_at, updated_at FROM teams.team WHERE id = $1`
	row := r.pool.QueryRow(ctx, query, taskID)

	var teamModel TeamModel
	err := row.Scan(&teamModel.ID, &teamModel.Name, &teamModel.ParentID, &teamModel.CreatedAt, &teamModel.UpdatedAt)
	if err != nil {
		if errors.Is(err, postgres_pool.ErrNoRows) {
			return domain.Team{}, fmt.Errorf("task with id='%d': %w", taskID, errrors.ErrNotFound)
		}
		return domain.Team{}, fmt.Errorf("scan error %w", err)

	}

	teamDomain := domain.NewTeam(teamModel.ID, teamModel.Name, teamModel.ParentID, teamModel.CreatedAt, teamModel.UpdatedAt)

	return teamDomain, nil
}
