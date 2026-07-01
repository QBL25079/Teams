package team_service

import (
	"context"
	"fmt"

	"github.com/QBL25079/teams/internal/core/domain"
)

func (s *TeamService) UpdateTeam(ctx context.Context, id int, patch domain.TeamPatch) (domain.Team, error) {
	team, err := s.teamRepository.GetTeam(ctx, id)
	if err != nil {
		return domain.Team{}, fmt.Errorf("team does not exist: %w", err)
	}

	if err := team.ApplyPatch(patch); err != nil {
		return domain.Team{}, fmt.Errorf("apply team patch: %w", err)
	}

	patchedTeam, err := s.teamRepository.UpdateTeam(ctx, id, team)
	if err != nil {
		return domain.Team{}, fmt.Errorf("failed to save team: %w", err)
	}

	return patchedTeam, nil
}
