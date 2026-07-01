package user_service

import (
	"context"
	"fmt"

	"github.com/QBL25079/teams/internal/core/domain"
	errors "github.com/QBL25079/teams/internal/core/errors"
)

func (s *UserService) GetUsers(ctx context.Context, limit, offset, teamID *int) ([]domain.User, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("limit must be more than 0: %w", errors.ErrInvalidArgument)
	}
	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("offset must be more than 0: %w", errors.ErrInvalidArgument)
	}

	users, err := s.userRepository.GetUsers(ctx, limit, offset, teamID)
	if err != nil {
		return nil, fmt.Errorf("Get users from repo: %w", err)
	}

	return users, nil
}
