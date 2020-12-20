package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	GIN_USER_KEY    = "GIN_USER_KEY"
	GET_OPEN_ID_KEY = "GIN_OPEN_ID_KEY"
)

type Middleware func(*Context)

func CreateMiddleware(middleware Middleware) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context := &Context{ctx}
		middleware(context)
	}
}

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method

		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, open_id, Authorization, token, sign, timestamp")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		ctx.Next()
	}
}

func CheckLogin(key string, getUserByOpenID func(context.Context, string) (interface{}, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		openID := ctx.Request.Header.Get("key")
		if openID == "" {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    NOLOGIN,
				"message": "You are logout.",
			})
			return
		}
		if getUserByOpenID == nil {
			ctx.Set(GET_OPEN_ID_KEY, openID)
			ctx.Next()
			return
		}
		user, err := getUserByOpenID(ctx.Request.Context(), openID)
		if err != nil {
			ctx.Abort()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
					"code":    NOLOGIN,
					"message": "User not found.",
				})
			} else {
				ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    SERVER_ERROR,
					"message": err.Error(),
				})
			}
			return
		}
		ctx.Set(GIN_USER_KEY, user)
		ctx.Set(GET_OPEN_ID_KEY, openID)
		ctx.Next()
	}
}
