package user_transport_http

import (
	"context"
	"net/http"

	"github.com/QBL25079/teams/internal/core/domain"
	core_http_server "github.com/QBL25079/teams/internal/core/transport/http/server"
)

type UserHTTPHandler struct {
	userService UserService
}

type UserService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit, offset, teamID *int) ([]domain.User, error)
	DeleteUser(ctx context.Context, userID int) error
	GetUser(ctx context.Context, userID int) (domain.User, error)
	UpdateUser(ctx context.Context, id int, patch domain.UserPatch) (domain.User, error)
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
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetUsers,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: h.DeleteUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.GetUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: h.UpdateUser,
		},
	}
}
