package user_transport_http

import (
	"context"

	core_domain "github.com/QBL25079/teams/internal/core/domain"
)

type UserHTTPHandler struct {
	userService UserService
}

type UserService interface {
	CreateUser(ctx context.Context, user core_domain.User) (core_domain.User, error)
}