package bootstrap

import (
	"github.com/arkiant/ddd-golang-framework/internal/platform/server/handler"
	"github.com/arkiant/ddd-golang-framework/kit/cqrs/command"
	"github.com/arkiant/ddd-golang-framework/kit/cqrs/query"
	"github.com/arkiant/ddd-golang-framework/kit/http/server"
)

func routes(queryBus query.Bus, commandBus command.Bus) []server.Route {
	return []server.Route{
		server.NewRoute("GET", "ping", handler.PingHandler(queryBus)),
	}
}
