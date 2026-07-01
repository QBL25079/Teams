package team_repository

import (
	"context"
	"fmt"

	"github.com/QBL25079/teams/internal/core/errors"
)

func (r *TeamRepository) DeleteTeam(ctx context.Context, teamID int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `DELETE FROM teams.team WHERE id=$1`

	cmdTag, err := r.pool.Exec(ctx, query, teamID)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("team with id='%d': %w", teamID, errors.ErrNotFound)
	}

	return nil
}
