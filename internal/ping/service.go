package ping

import (
	"context"

	"github.com/arkiant/ddd-golang-framework/kit/cqrs/event"
)

type PingService struct {
	eventBus event.Bus
}

func NewPingService(eventBus event.Bus) PingService {
	return PingService{eventBus: eventBus}
}

func (f PingService) GetPing(ctx context.Context) (interface{}, error) {
	return "PONG", nil
}