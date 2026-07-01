package teams_transport

import (
	"net/http"

	"github.com/QBL25079/teams/internal/core/logger"
	"github.com/QBL25079/teams/internal/core/transport/http/request"
	"github.com/QBL25079/teams/internal/core/transport/http/response"
)

func (h *TeamHTTPHandler) DeleteTeam(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)

	responseHandler := response.NewHTTPResponseHandler(log, rw)

	teamID, err := request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get teamID path value")
		return
	}

	if err := h.teamService.DeleteTeam(ctx, teamID); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete team")
		return
	}

	responseHandler.NoContentResponse()
}
