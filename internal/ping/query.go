package ping

import (
	"context"
	"errors"

	"github.com/arkiant/hexagonal-golang-api/kit/cqrs/query"
)

const PingQueryType query.Type = "query.ping"

type PingQuery struct{}

func NewPingQuery() PingQuery {
	return PingQuery{}
}

func (f PingQuery) Type() query.Type {
	return PingQueryType
}

type PingQueryHandler struct {
	service PingService
}

func NewPingQueryHandler(service PingService) PingQueryHandler {
	return PingQueryHandler{service: service}
}

func (f PingQueryHandler) Handle(ctx context.Context, query query.Query) (interface{}, error) {

	_, ok := query.(PingQuery)
	if !ok {
		return "", errors.New("unexpected query")
	}

	return f.service.Ping(ctx)

}
