package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const GIN_USER_KEY = "GIN_USER_KEY"

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

func CheckOpenID(getUserByOpenID func(openID string) (interface{}, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		openID := ctx.Request.Header.Get("open_id")
		if openID == "" {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    NOLOGIN,
				"message": "You are logout.",
			})
		} else {
			if getUserByOpenID == nil {
				ctx.Set("open_id", openID)
				ctx.Next()
				return
			}
			user, err := getUserByOpenID(openID)
			if err != nil {
				ctx.Abort()
				if gorm.IsRecordNotFoundError(err) {
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
			ctx.Set("GIN_USER_KEY", user)
			ctx.Set("open_id", openID)
			ctx.Next()
		}
	}
}
