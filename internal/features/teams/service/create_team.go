package team_service

import (
	"context"
	"fmt"

	"github.com/QBL25079/teams/internal/core/domain"
)

func (s *TeamService) CreateTeam(ctx context.Context, team domain.Team) (domain.Team, error) {
	if err := team.Validate(); err != nil {
		return domain.Team{}, fmt.Errorf("validate user domain: %w", err)
	}

	team, err := s.teamRepository.CreateTeam(ctx, team)
	if err != nil {
		return domain.Team{}, fmt.Errorf("Can`t create team: %w", err)
	}

	return team, nil
}