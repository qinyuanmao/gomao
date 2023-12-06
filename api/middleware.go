package api

import (
	"context"
	"errors"
	"net/http"

	"e.coding.net/tssoft/repository/gomao/security"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

		//放行所有 OPTIONS 方法
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		return
	})
}

func CheckJWTToken[T any](getUser func(context.Context, int64) (T, error)) gin.HandlerFunc {
	return CreateMiddleware(func(ctx *Context) (resultCode ResultCode, message string, result interface{}) {
		authorization := ctx.Request.Header.Get("Authorization")
		c, err := security.JwtDecode(authorization)
		if err != nil {
			return WithLogout()
		}
		ctx.Set("JWT", c)
		var userId, error = c.GetUserId()
		if error != nil {
			return WithLogout()
		}
		ctx.Set("USER_ID", userId)
		user, err := getUser(ctx.Request.Context(), userId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return WithLogout()
			} else {
				return WithServerError(err)
			}
		}
		ctx.Set("USER", user)
		ctx.Next()
		return
	})
}
