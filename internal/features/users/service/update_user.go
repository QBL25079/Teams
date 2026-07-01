package user_service

import (
	"context"
	"fmt"

	"github.com/QBL25079/teams/internal/core/domain"
)

func (s *UserService) UpdateUser(ctx context.Context, id int, patch domain.UserPatch) (domain.User, error) {
	user, err := s.userRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("user does not exist: %w", err)
	}

	if err := user.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("apply user patch: %w", err)
	}

	patchedUser, err := s.userRepository.UpdateUser(ctx, id, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to save user: %w", err)
	}

	return patchedUser, nil
}