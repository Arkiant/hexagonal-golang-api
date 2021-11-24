package handler

import (
	"net/http"

	"github.com/arkiant/hexagonal-golang-api/internal/ping"
	"github.com/arkiant/hexagonal-golang-api/kit/cqrs/query"
	"github.com/gin-gonic/gin"
)

type ResponsePong struct {
	Response interface{} `json:"response"`
}

func PingHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := queryBus.Dispatch(ctx, ping.NewPingQuery())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, ResponsePong{Response: response})
	}
}
