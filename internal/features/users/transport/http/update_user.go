package user_transport_http

import (
	"net/http"

	"github.com/QBL25079/teams/internal/core/domain"
	"github.com/QBL25079/teams/internal/core/logger"
	"github.com/QBL25079/teams/internal/core/transport/http/request"
	"github.com/QBL25079/teams/internal/core/transport/http/response"
)

type UpdateUserRequest struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	TeamID    *int    `json:"team_id"`
}

type UpdateUserResponse DTOUserResponse

func (h *UserHTTPHandler) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)
	userID, err := request.GetIntPathValue(r, "id")

	if err != nil {
		responseHandler.ErrorResponse(err, "Cant get user ID")
		return
	}

	var req UpdateUserRequest

	if err := request.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	userPatch := userPatchFromRequest(req)

	userDomain, err := h.userService.UpdateUser(ctx, userID, userPatch)

	response := UpdateUserResponse(UserDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusOK)

}

func userPatchFromRequest(req UpdateUserRequest) domain.UserPatch {
	return domain.NewUserPatch(
		req.FirstName,
		req.LastName,
		req.TeamID,
	)
}