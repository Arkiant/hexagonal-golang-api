package ping

import (
	"context"
)

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (f Service) Ping(ctx context.Context) (interface{}, error) {
	return "PONG", nil
}
