package teams_transport

import (
	"context"
	"net/http"

	"github.com/QBL25079/teams/internal/core/domain"
	"github.com/QBL25079/teams/internal/core/transport/http/server"
)

type TeamHTTPHandler struct {
	teamService TeamService
}

type TeamService interface {
	CreateTeam(ctx context.Context, team domain.Team) (domain.Team, error)
	GetTeams(ctx context.Context, limit, offset *int) ([]domain.Team, error)
	DeleteTeam(ctx context.Context, teamID int) error
	UpdateTeam(ctx context.Context, id int, patch domain.TeamPatch) (domain.Team, error)
	GetTeam(ctx context.Context, teamID int) (domain.Team, error)
}

func NewTeamHTTPHandler(teamService TeamService) *TeamHTTPHandler {
	return &TeamHTTPHandler{teamService: teamService}
}

func (h *TeamHTTPHandler) Routes() []server.Route {
	return []server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/team",
			Handler: h.CreateTeam,
		},
		{
			Method:  http.MethodGet,
			Path:    "/team",
			Handler: h.GetTeams,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/team/{id}",
			Handler: h.DeleteTeam,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/team/{id}",
			Handler: h.UpdateTeam,
		},
		{
			Method: http.MethodGet,
			Path: "/team/{id}",
			Handler: h.GetTeam,
		},
	}
}
