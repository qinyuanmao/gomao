package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qinyuanmao/gomao/dingtalk"
	"github.com/qinyuanmao/gomao/logger"
	"github.com/spf13/viper"
)

type ApiHandler func(ctx *gin.Context) (httpCode, resultCode int, message string, result interface{})
type FileHandler func(ctx *gin.Context) (httpCode int, message string, bytes []byte, fileName, contentType string)

func JsonApi(handler ApiHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httpCode, resultCode, message, result := handler(ctx)
		sendDingTalk(ctx.Request.URL.String(), message, httpCode)
		response := map[string]interface{}{
			"code":    resultCode,
			"message": message,
		}
		if result != nil {
			response["result"] = result
		}
		ctx.JSON(httpCode, response)
	}
}

func FileApi(hander FileHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httpCode, message, bytes, fileName, contentType := hander(ctx)
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
	if httpCode != http.StatusOK {
		logger.Error(message)
		webhook := viper.GetString("dingding_webhook")
		env := viper.GetString("env")
		if webhook != "" {
			dingtalk.GetInstance().Notify(&dingtalk.DingTalkMsg{
				MsgType: "markdown",
				Markdown: dingtalk.Markdown{
					Title: "监控报警",
					Text:  fmt.Sprintf("## 【%s】[%s] 接口请求异常: %d\n\n > 错误信息: %s", env, url, httpCode, message),
				},
				At: dingtalk.At{
					AtMobiles: []string{"18583872978"},
					IsAtAll:   false,
				},
			})
		}
	}
}
