package user_service

import (
	"context"
	"fmt"

	"github.com/QBL25079/teams/internal/core/domain"
)

func (s *UserService) GetUser(ctx context.Context, id int) (domain.User, error) {
	user, err := s.userRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("Get user from repository: %w", err)
	}
	return user, nil
}
