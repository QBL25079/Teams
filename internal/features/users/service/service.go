package user_service

import (
	"context"

	domain "github.com/QBL25079/teams/internal/core/domain"
)

type UserService struct {
	userRepository UserRepository
}

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit, offset, teamID *int) ([]domain.User, error)
	DeleteUser(ctx context.Context, userID int) error
	GetUser(ctx context.Context, id int) (domain.User, error)
}

func NewUserService(userRepositry UserRepository) *UserService {
	return &UserService{userRepository: userRepositry}
}
