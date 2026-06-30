package user_transport_http

import (
	"encoding/json"
	"net/http"

	core_domain "github.com/QBL25079/teams/internal/core/domain"
	core_logger "github.com/QBL25079/teams/internal/core/logger"
	core_http_response "github.com/QBL25079/teams/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FirstName string `json:"first_name" validate:"required, min=3, max=10"`
	LastName string `json:"last_name" validate:"required, min=3, max=10"`
	BirthYear int `json:"birth_yaer" required:"true"`
	GroupID   *int `json:"group_id" validate:"omitempty"`
}

type CreateUserResponse DTOUserResponse

func (h *UserHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw) 

	var request CreateUserRequest
	log.Debug("Invoce CreateUser Handler")

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		responseHandler.ErrorResponse(err, "failed to decode and validate create user request")
		return 
	}
	userDomain := domainFromDTO(request)
	userDomain, err := h.userService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	response := CreateUserResponse(UserDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) core_domain.User {
	return core_domain.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		BirthYear: dto.BirthYear,
		GroupID:   dto.GroupID,
	}
}