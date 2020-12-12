package api

import (
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

func (ctx *Context) OpenID() string {
	return ctx.GetHeader("open_id")
}

func (ctx *Context) Token() string {
	return ctx.GetHeader("token")
}
