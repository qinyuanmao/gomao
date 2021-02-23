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

func CreateMiddleware(middleware ApiHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context := &Context{ctx}
		resultCode, message, _ := middleware(context)
		if resultCode != SUCCESS {
			ctx.Abort()
			ctx.String(resultCode.getHttpCode(), message)
			sendDingTalk(ctx.Request.URL.String(), message, resultCode.getHttpCode())
		}
		ctx.Next()
	}
}

func Cors() gin.HandlerFunc {
	return CreateMiddleware(func(ctx *Context) (resultCode ResultCode, message string, result interface{}) {
		method := ctx.Request.Method

		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, open_id, Authorization, token, sign, timestamp, Token, OpenID")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		return
	})
}

func CheckLogin(key string, getUserByOpenID func(context.Context, string) (interface{}, error)) gin.HandlerFunc {
	return CreateMiddleware(func(ctx *Context) (resultCode ResultCode, message string, result interface{}) {
		openID := ctx.Request.Header.Get(key)
		if openID == "" {
			return WithLogout()
		}
		if getUserByOpenID == nil {
			ctx.Set(GET_OPEN_ID_KEY, openID)
			return
		}
		user, err := getUserByOpenID(ctx.Request.Context(), openID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return WithLogout()
			} else {
				return WithServerError(err)
			}
		}
		ctx.Set(GIN_USER_KEY, user)
		ctx.Set(GET_OPEN_ID_KEY, openID)
		return
	})
}
