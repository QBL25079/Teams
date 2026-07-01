package team_repository

import (
	"time"

	"github.com/QBL25079/teams/internal/core/domain"
)

type TeamModel struct {
	ID        int
	Name      string
	ParentID  *int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func teamDomainsFromModels(teamModels []TeamModel) []domain.Team {
	teamDomains := make([]domain.Team, len(teamModels))

	for i, model := range teamModels {
		teamDomains[i] = domain.NewTeam(
			model.ID,
			model.Name,
			model.ParentID,
			model.CreatedAt,
			model.UpdatedAt,
		)
	}

	return teamDomains
}