package ping

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	pingService := NewService()
	res, err := pingService.Ping(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, "PONG", res)
}
