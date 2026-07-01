package teams_transport

import (
	"net/http"

	"github.com/QBL25079/teams/internal/core/domain"
	"github.com/QBL25079/teams/internal/core/logger"
	"github.com/QBL25079/teams/internal/core/transport/http/request"
	"github.com/QBL25079/teams/internal/core/transport/http/response"
)

type PatchTeamRequest struct {
	Name     *string `json:"name"`
	ParentID *int    `json:"parent_id"`
}

type PatchTeamResponse TeamDTOResponse

func (h *TeamHTTPHandler) UpdateTeam(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)
	userID, err := request.GetIntPathValue(r, "id")

	if err != nil {
		responseHandler.ErrorResponse(err, "Cant get user ID")
		return
	}

	var req PatchTeamRequest

	if err := request.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	teamPatch := teamPatchFromRequest(req)

	teamDomain, err := h.teamService.UpdateTeam(ctx, userID, teamPatch)

	response := PatchTeamResponse(teamDTOFromDomain(teamDomain))

	responseHandler.JSONResponse(response, http.StatusOK)

}


func teamPatchFromRequest(request PatchTeamRequest) domain.TeamPatch {
	return domain.TeamPatch{
		Name:     request.Name,
		ParentID: request.ParentID,
	}
}
