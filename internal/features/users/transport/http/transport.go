package user_transport_http

import (
	"context"
	"net/http"

	core_domain "github.com/QBL25079/teams/internal/core/domain"
	core_http_server "github.com/QBL25079/teams/internal/core/transport/http/server"
)

type UserHTTPHandler struct {
	userService UserService
}

type UserService interface {
	CreateUser(ctx context.Context, user core_domain.User) (core_domain.User, error)
	GetUsers(ctx context.Context, limit, offset *int) ([]core_domain.User, error)
}

func NewUsersHTTPHandler(usersService UserService) *UserHTTPHandler {
	return &UserHTTPHandler{userService: usersService}
}

func (h *UserHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
		{
			Method: http.MethodGet,
			Path: "/users",
			Handler: h.GetUsers,
		},
	}
}
