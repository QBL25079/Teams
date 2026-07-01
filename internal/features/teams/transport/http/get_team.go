package teams_transport

import (
	"net/http"

	"github.com/QBL25079/teams/internal/core/logger"
	"github.com/QBL25079/teams/internal/core/transport/http/request"
	"github.com/QBL25079/teams/internal/core/transport/http/response"
)

type GetTeamResponse TeamDTOResponse

func (h TeamHTTPHandler) GetTeam(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)

	responseHandler := response.NewHTTPResponseHandler(log, rw)

	teamID, err := request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get team id path value")
		return
	}

	teamDomain, err := h.teamService.GetTeam(ctx, teamID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get team")
		return
	}

	response := GetTeamResponse(teamDTOFromDomain(teamDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}
