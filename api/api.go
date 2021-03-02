package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qinyuanmao/gomao/dingtalk"
	"github.com/qinyuanmao/gomao/logger"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

type ApiHandler func(ctx *Context) (resultCode ResultCode, message string, result interface{})
type FileHandler func(ctx *Context) (httpCode int, message string, bytes []byte, fileName, contentType string)

func JsonApi(handler ApiHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resultCode, message, result := handler(&Context{ctx})
		sendDingTalk(ctx.Request.URL.String(), message, resultCode.getHttpCode())
		if resultCode == SUCCESS {
			response := map[string]interface{}{
				"code":    resultCode,
				"message": message,
			}
			if result != nil {
				response["result"] = result
			}
			ctx.JSON(resultCode.getHttpCode(), response)
		} else {
			ctx.String(resultCode.getHttpCode(), message)
		}
	}
}

func FileApi(hander FileHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httpCode, message, bytes, fileName, contentType := hander(&Context{ctx})
		sendDingTalk(ctx.Request.URL.String(), message, httpCode)

		if httpCode != http.StatusOK {
			ctx.Writer.WriteHeader(httpCode)
			ctx.Writer.Write([]byte(message))
			return
		}

		ctx.Writer.Header().Set("Content-Type", contentType)
		ctx.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", fileName))
		ctx.Writer.Header().Set("Accept-Length", fmt.Sprintf("%d", len(bytes)))
		ctx.Writer.Write(bytes)
	}
}

func sendDingTalk(url, message string, httpCode int) {
	if httpCode == http.StatusInternalServerError {
		logger.Error(message)
		webhook := viper.GetString("dingtalk.webhook")
		env := viper.GetString("env")
		if webhook != "" {
			var atMobiles = make([]string, 0)
			if env == "production" {
				atMobiles = viper.GetStringSlice("dingtalk.at_mobiles")
			}
			dingtalk.GetInstance().Notify(&dingtalk.DingTalkMsg{
				MsgType: "markdown",
				Markdown: dingtalk.Markdown{
					Title: "监控报警",
					Text: fmt.Sprintf("## %s \n\n ### 【%s】[%s] 接口请求异常: \n\n > 错误信息: %s \n\n %s", viper.GetString("project_name"), env, url, message,
						strings.Join(funk.Map(atMobiles, func(item string) string {
							return fmt.Sprintf("@%s", item)
						}).([]string), ",")),
				},
				At: dingtalk.At{
					AtMobiles: atMobiles,
					IsAtAll:   false,
				},
			})
		}
	}
}
