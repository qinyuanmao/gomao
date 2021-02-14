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

func (ctx *Context) Reject(resultCode ResultCode, message string) {
	ctx.Context.Abort()
	response := map[string]interface{}{
		"code":    resultCode,
		"message": message,
	}
	ctx.JSON(resultCode.getHttpCode(), response)
}

type Middleware func(*Context) (reject bool, resultCode ResultCode, message string)

func CreateMiddleware(middleware Middleware) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context := &Context{ctx}
		reject, resultCode, message := middleware(context)
		if reject {
			ctx.Abort()
			ctx.JSON(resultCode.getHttpCode(), map[string]interface{}{
				"code":    resultCode,
				"message": message,
			})
			return
		}
		ctx.Next()
	}
}

func Cors() gin.HandlerFunc {
	return CreateMiddleware(func(ctx *Context) (reject bool, resultCode ResultCode, message string) {
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
	return CreateMiddleware(func(ctx *Context) (reject bool, resultCode ResultCode, message string) {
		openID := ctx.Request.Header.Get(key)
		if openID == "" {
			return true, NOLOGIN, "You are logout."
		}
		if getUserByOpenID == nil {
			ctx.Set(GET_OPEN_ID_KEY, openID)
			return
		}
		user, err := getUserByOpenID(ctx.Request.Context(), openID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return true, NOLOGIN, "You are logout."
			} else {
				return true, SERVER_ERROR, err.Error()
			}
		}
		ctx.Set(GIN_USER_KEY, user)
		ctx.Set(GET_OPEN_ID_KEY, openID)
		return
	})
}
