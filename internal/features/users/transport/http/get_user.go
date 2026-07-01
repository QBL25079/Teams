package user_transport_http

import (
	"net/http"

	"github.com/QBL25079/teams/internal/core/logger"
	"github.com/QBL25079/teams/internal/core/transport/http/request"
	"github.com/QBL25079/teams/internal/core/transport/http/response"
)

type GetUserResponse DTOUserResponse

func (h *UserHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	userID, err := request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "Missing ID value")
		return
	}
	user, err := h.userService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "Cant get user with this ID")
		return
	}

	response := GetUserResponse(DTOUserResponse(user))

	responseHandler.JSONResponse(response, http.StatusOK)
}
