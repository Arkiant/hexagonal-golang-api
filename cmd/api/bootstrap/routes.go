package bootstrap

import (
	"github.com/arkiant/hexagonal-golang-api/internal/platform/server/handler"
	"github.com/arkiant/hexagonal-golang-api/kit/cqrs/command"
	"github.com/arkiant/hexagonal-golang-api/kit/cqrs/query"
	"github.com/arkiant/hexagonal-golang-api/kit/http/server"
)

func routes(queryBus query.Bus, commandBus command.Bus) []server.Route {
	return []server.Route{
		server.NewRoute("GET", "ping", handler.PingHandler(queryBus)),
	}
}
