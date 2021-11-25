package ping

import (
	"context"
)

type PingService struct {
}

func NewPingService() PingService {
	return PingService{}
}

func (f PingService) Ping(ctx context.Context) (interface{}, error) {
	return "PONG", nil
}
