package team_service

import (
	"context"
	"fmt"

	"github.com/QBL25079/teams/internal/core/domain"
	"github.com/QBL25079/teams/internal/core/errors"
)

func (s *TeamService) GetTeams(ctx context.Context, limit, offset *int) ([]domain.Team, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("limit must be more than 0: %w", errors.ErrInvalidArgument)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("Offset must be more than 0: %w", errors.ErrInvalidArgument)
	}

	teams, err := s.teamRepository.GetTeams(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("Get users from repo: %w", err)
	}

	return teams, nil
}
