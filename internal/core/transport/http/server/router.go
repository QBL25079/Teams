package server

import (
	"fmt"
	"net/http"

	core_http_middleware "github.com/QBL25079/teams/internal/core/transport/http/middleware"
)

type ApiVersion string

var (
	ApiVersion1 = ApiVersion("/v1")
)

type ApiVersionRouter struct {
	*http.ServeMux
	apiVersion ApiVersion
	middleWare []core_http_middleware.Middleware
}

func NewApiVersionRouter(apiVersion ApiVersion, middleWare ...core_http_middleware.Middleware) *ApiVersionRouter {
	return &ApiVersionRouter{ServeMux: http.NewServeMux(), apiVersion: apiVersion, middleWare: middleWare}
}

func (a *ApiVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		a.Handle(pattern, route.WithMiddleWare())
	}
}

func (a *ApiVersionRouter) WithMiddleWare() http.Handler {
	return core_http_middleware.ChainMiddleware(a, a.middleWare...)
}
