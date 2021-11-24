package ping

import (
	"context"

	"github.com/arkiant/hexagonal-golang-api/kit/cqrs/event"
)

type PingService struct {
	eventBus event.Bus
}

func NewPingService(eventBus event.Bus) PingService {
	return PingService{eventBus: eventBus}
}

func (f PingService) Ping(ctx context.Context) (interface{}, error) {
	return "PONG", nil
}
