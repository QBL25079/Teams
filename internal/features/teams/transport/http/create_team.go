package teams_transport

import (
	"net/http"

	"github.com/QBL25079/teams/internal/core/domain"
	"github.com/QBL25079/teams/internal/core/logger"
	request "github.com/QBL25079/teams/internal/core/transport/http/request"
	"github.com/QBL25079/teams/internal/core/transport/http/response"
)

type CreateTeamRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=15"`
	ParentID *int   `json:"parent_id" validate:"omitempty"`
}

type CreateTeamResponse TeamDTOResponse

func (h *TeamHTTPHandler) CreateTeam(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)

	responseHandler := response.NewHTTPResponseHandler(log, rw)

	var req CreateTeamRequest

	if err := request.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	teamDomain := domain.NewTeamUninitialized(
		req.Name,
		req.ParentID,
	)

	team, err := h.teamService.CreateTeam(ctx, teamDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create team")
		return
	}

	resp := CreateTeamResponse(teamDTOFromDomain(team))

	responseHandler.JSONResponse(resp, http.StatusCreated)
}
