package team_service

import (
	"context"

	"github.com/QBL25079/teams/internal/core/domain"
)

type TeamService struct {
	teamRepository TeamRepository
}

type TeamRepository interface {
	CreateTeam(ctx context.Context, team domain.Team) (domain.Team, error)
	GetTeams(ctx context.Context, limit, offset *int) ([]domain.Team, error)
	DeleteTeam(ctx context.Context, teamID int) error
	UpdateTeam(ctx context.Context, id int, team domain.Team) (domain.Team, error)
	GetTeam(ctx context.Context, teamID int) (domain.Team, error)
}

func NewTeamService(teamRepository TeamRepository) *TeamService {
	return &TeamService{teamRepository: teamRepository}
}
