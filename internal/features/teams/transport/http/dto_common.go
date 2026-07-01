package teams_transport

import (
	"time"

	"github.com/QBL25079/teams/internal/core/domain"
)

type TeamDTOResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	ParentID  *int      `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func teamDTOFromDomain(team domain.Team) TeamDTOResponse {
	return TeamDTOResponse{ID: team.ID, Name: team.Name, ParentID: team.ParentID, CreatedAt: team.CreatedAt, UpdatedAt: team.UpdatedAt}
}

func teamsDTOsFromDomains(tasks []domain.Team) []TeamDTOResponse {
	dtos := make([]TeamDTOResponse, len(tasks))

	for i, task := range tasks {
		dtos[i] = teamDTOFromDomain(task)
	}
	return dtos
}
