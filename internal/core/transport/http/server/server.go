package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_logger "github.com/QBL25079/teams/internal/core/logger"
	core_http_middleware "github.com/QBL25079/teams/internal/core/transport/http/middleware"
	"go.uber.org/zap"
)

type HTTPServer struct {
	mux        *http.ServeMux
	config     Config
	log        *core_logger.Logger
	middleware []core_http_middleware.Middleware
}

func NewHTTPServer(config Config, logger *core_logger.Logger, middlware ...core_http_middleware.Middleware) *HTTPServer {
	return &HTTPServer{mux: http.NewServeMux(), config: config, log: logger, middleware: middlware}
}

func (h *HTTPServer) Run(ctx context.Context) error {
	mux := core_http_middleware.ChainMiddleware(h.mux, h.middleware...)

	server := &http.Server{Addr: h.config.Addr, Handler: mux}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)
		h.log.Info("Start HTTP Server,", zap.String("addr", h.config.Addr))

		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("ListenAndServer HTTP: %w", err)
		}
	case <-ctx.Done():
		h.log.Info("shutdown http server...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), h.config.ShutdownTimeout)

		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown HTTP server: %w", err)
		}
		h.log.Info("HTTP server stopped")
	}
	return nil
}

func (s *HTTPServer) RegisterAPIRouters(routers ...*ApiVersionRouter) {
	for _, router := range routers {
		prefix := "/api" + string(router.apiVersion)

		s.mux.Handle(prefix+"/", http.StripPrefix(prefix, router.WithMiddleWare()))
	}
}
