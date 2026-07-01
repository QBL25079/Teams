package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/QBL25079/teams/internal/core/logger"
	core_pgx_pool "github.com/QBL25079/teams/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/QBL25079/teams/internal/core/transport/http/middleware"
	core_http_server "github.com/QBL25079/teams/internal/core/transport/http/server"
	team_repository "github.com/QBL25079/teams/internal/features/teams/repository"
	team_service "github.com/QBL25079/teams/internal/features/teams/service"
	teams_transport "github.com/QBL25079/teams/internal/features/teams/transport/http"
	user_repository "github.com/QBL25079/teams/internal/features/users/repository"
	user_service "github.com/QBL25079/teams/internal/features/users/service"
	user_transport_http "github.com/QBL25079/teams/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init app logger: %w", err)
		os.Exit(1)
	}

	defer logger.Close()

	logger.Debug("initializing connection pool")
	pool, err := core_pgx_pool.NewPool(ctx, core_pgx_pool.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}

	defer pool.Close()

	logger.Debug("Initializing feature ", zap.String("feature", "users"))
	usersRepository := user_repository.NewUsersRepository(pool)
	userService := user_service.NewUserService(usersRepository)
	usersTransportHTTP := user_transport_http.NewUsersHTTPHandler(userService)

	logger.Debug("Initializing feature ", zap.String("feature", "teams"))
	teamsRepository := team_repository.NewTeamRepository(pool)
	teamService := team_service.NewTeamService(teamsRepository)
	teamHandler := teams_transport.NewTeamHTTPHandler(teamService)

	logger.Debug("initializing HTTP server")

	httpServer := core_http_server.NewHTTPServer(core_http_server.NewConfigMust(), logger, core_http_middleware.RequestID(), core_http_middleware.Logger(logger), core_http_middleware.Trace(), core_http_middleware.Recover())

	apiVersionRouter := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)

	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)
	apiVersionRouter.RegisterRoutes(teamHandler.Routes()...)

	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
