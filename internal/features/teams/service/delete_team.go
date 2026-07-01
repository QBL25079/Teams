package team_service

import (
	"context"
	"fmt"
)

func (s *TeamService) DeleteTeam(ctx context.Context, teamID int) error {
	if err := s.teamRepository.DeleteTeam(ctx, teamID); err != nil {
		return fmt.Errorf("delete task: %w", err)
	}

	return nil
}