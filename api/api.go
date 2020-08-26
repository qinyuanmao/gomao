package api

import (
	"github.com/gin-gonic/gin"
)

type ApiHandler func(ctx *gin.Context) (httpCode, resultCode int, message string, result interface{})

func JsonApi(handler ApiHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httpCode, resultCode, message, result := handler(ctx)
		ctx.JSON(httpCode, map[string]interface{}{
			"code":    resultCode,
			"message": message,
			"result":  result,
		})
	}
}
