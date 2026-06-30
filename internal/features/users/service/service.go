package user_service

import (
	"context"

	core_domain "github.com/QBL25079/teams/internal/core/domain"
)

type UserService struct {
	userRepository UserRepository
}

type UserRepository interface {
	CreateUser(ctx context.Context, user core_domain.User) (core_domain.User, error)
}

func NewUserService(userRepositry UserRepository) *UserService {
	return &UserService{userRepository: userRepositry}
}