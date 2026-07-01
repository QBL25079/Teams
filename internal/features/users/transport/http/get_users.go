package user_transport_http

import (
	"fmt"
	"net/http"
	"strconv"

	core_logger "github.com/QBL25079/teams/internal/core/logger"
	core_http_request "github.com/QBL25079/teams/internal/core/transport/http/request"
	core_http_response "github.com/QBL25079/teams/internal/core/transport/http/response"
)

type GetUsersResponse []DTOUserResponse

func (h *UserHTTPHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit / offset query params")
		return
	}

	teamID, err := getTeamIDQueryParam(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get team_id query param")
		return
	}

	userDomains, err := h.userService.GetUsers(ctx, limit, offset, teamID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get list of users")
		return
	}

	response := GetUsersResponse(UsersDTOFromDomains(userDomains))

	responseHandler.JSONResponse(response, http.StatusOK)
}

// Существующая функция
func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	const (
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)

	limit, err := core_http_request.GetQueryParam(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' query params: %w", err)
	}

	offset, err := core_http_request.GetQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf("get 'offset' query params: %w", err)
	}

	return limit, offset, nil
}

// Новая функция
func getTeamIDQueryParam(r *http.Request) (*int, error) {
	teamIDStr := r.URL.Query().Get("team_id")
	if teamIDStr == "" {
		return nil, nil
	}

	teamID, err := strconv.Atoi(teamIDStr)
	if err != nil {
		return nil, fmt.Errorf("team_id must be a valid integer")
	}

	if teamID <= 0 {
		return nil, fmt.Errorf("team_id must be greater than 0")
	}

	return &teamID, nil
}
