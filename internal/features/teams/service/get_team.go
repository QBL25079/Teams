package team_service

import (
	"context"
	"fmt"

	"github.com/QBL25079/teams/internal/core/domain"
)

func (s *TeamService) GetTeam(ctx context.Context, taskID int) (domain.Team, error) {
	task, err := s.teamRepository.GetTeam(ctx, taskID)
	if err != nil {
		return domain.Team{}, fmt.Errorf("get team from repository: %w", err)
	}

	return task, nil
}
