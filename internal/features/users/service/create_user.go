package user_service

import (
	"context"
	"fmt"

	core_domain "github.com/QBL25079/teams/internal/core/domain"
)

func (s UserService) CreateUser(ctx context.Context, user core_domain.User) (core_domain.User, error) {
	if err := user.Validate(); err != nil {
		return core_domain.User{}, fmt.Errorf("validate user domain: %w", err)
	}

	user, err := s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return core_domain.User{}, fmt.Errorf("Can`t create user: %w", err)
	}

	return user, nil
}