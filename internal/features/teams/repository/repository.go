package team_repository

import core_postgres_pool "github.com/QBL25079/teams/internal/core/repository/postgres/pool"

type TeamRepository struct {
	pool core_postgres_pool.Pool
}

func NewTeamRepository(pool core_postgres_pool.Pool) *TeamRepository {
	return &TeamRepository{pool: pool}
}
