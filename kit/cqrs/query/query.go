package query

import "context"

// Bus defines the expected behaviour from a query bus.
type Bus interface {
	// Dispatch is the method used to dispatch new queries.
	Dispatch(context.Context, Query) (interface{}, error)
	// Register is the method used to register a new query handler.
	Register(Type, Handler)
}

//go:generate mockery --case=snake --outpkg=querymocks --output=querymocks --name=Bus

// Type represents an application query type.
type Type string

// Query represents an application query.
type Query interface {
	Type() Type
}

// Handler defines the expected behaviour from a query handler.
type Handler interface {
	Handle(context.Context, Query) (interface{}, error)
}
