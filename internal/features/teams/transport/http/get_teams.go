package teams_transport

import (
	"fmt"
	"net/http"

	"github.com/QBL25079/teams/internal/core/logger"
	"github.com/QBL25079/teams/internal/core/transport/http/request"
	"github.com/QBL25079/teams/internal/core/transport/http/response"
)

type GetTeamsResponse []TeamDTOResponse

func (h *TeamHTTPHandler) GetTeams(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)

	responseHandler := response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit / offset query params")
		return
	}

	teamDomains, err := h.teamService.GetTeams(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get list of teams")
		return
	}

	response := GetTeamsResponse(teamsDTOsFromDomains(teamDomains))

	responseHandler.JSONResponse(response, http.StatusOK)

}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	const (
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)

	limit, err := request.GetQueryParam(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' query params: %w", err)
	}

	offset, err := request.GetQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf("get 'offset' query params: %w", err)
	}

	return limit, offset, nil
}
