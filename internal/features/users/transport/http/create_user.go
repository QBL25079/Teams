package user_transport_http

import (
	"net/http"

	core_logger "github.com/QBL25079/teams/internal/core/logger"
	core_http_response "github.com/QBL25079/teams/internal/core/transport/http/response"
)

type CreateUserResponse struct {
	FirstName string `json:"first_name" validate:"required, min=3, max=10"`
	LastName string `json:"last_name" validate:"required, min=3, max=10"`
	BirthYear int `json:"birth_day" required:"true"`
	GroupID   *int `json:"group_id" validate:"omitempty"`
}

func (h *UserHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw) 

	var request 
}