package inmemory

import (
	"context"

	"github.com/arkiant/hexagonal-golang-api/kit/cqrs/query"
)

// QueryBus is an in-memory implementation of the command.Bus.
type QueryBus struct {
	handlers map[query.Type]query.Handler
}

// NewQueryBus initializes a new instance of QueryBus.
func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[query.Type]query.Handler),
	}
}

// Dispatch implements the query.Bus interface.
func (b *QueryBus) Dispatch(ctx context.Context, q query.Query) (interface{}, error) {
	handler, ok := b.handlers[q.Type()]
	if !ok {
		return "", nil
	}

	return handler.Handle(ctx, q)
}

// Register implements the query.Bus interface.
func (b *QueryBus) Register(queryType query.Type, handler query.Handler) {
	b.handlers[queryType] = handler
}
