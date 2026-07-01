package user_repository

import (
	"context"
	"fmt"

	"github.com/QBL25079/teams/internal/core/errors"
)

func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `DELETE FROM teams.user WHERE id=$1;`

	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("user with id='%d': %w", id, errors.ErrNotFound)
	}

	return nil
}
