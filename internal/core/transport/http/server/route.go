package server

import (
	"net/http"

	core_http_middleware "github.com/QBL25079/teams/internal/core/transport/http/middleware"
)

type Route struct {
	Method     string
	Path       string
	Handler    http.HandlerFunc
	MiddleWare []core_http_middleware.Middleware
}

func NewRoute(method, path string, handler http.HandlerFunc) Route {
	return Route{Method: method, Path: path, Handler: handler}
}

func (r *Route) WithMiddleWare() http.Handler {
	return core_http_middleware.ChainMiddleware(r.Handler, r.MiddleWare...)
}
