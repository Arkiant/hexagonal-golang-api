package ping

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arkiant/hexagonal-golang-api/kit/cqrs/query/querymocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	queryBus := new(querymocks.Bus)
	queryBus.On(
		"Dispatch",
		mock.Anything,
		mock.AnythingOfType("ping.PingQuery"),
	).Return("PONG", nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/ping", Handler(queryBus))

	t.Run("it returns 200", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/ping", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"data":"PONG"}`, rec.Body.String())
	})
}
