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

func JsonApi(handler ApiHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httpCode, resultCode, message, result := handler(ctx)
		if httpCode != http.StatusOK && httpCode != http.StatusBadRequest {
			logger.Error(message)
			webhook := viper.GetString("dingding_webhook")
			env := viper.GetString("env")
			if webhook != "" && env == "release" {
				dingtalk.GetInstance().Notify(&dingtalk.DingTalkMsg{
					MsgType: "markdown",
					Markdown: dingtalk.Markdown{
						Title: "监控报警",
						Text:  fmt.Sprintf("## [%s] 接口请求异常: %d\n\n > 错误信息: %s", ctx.Request.URL, httpCode, message),
					},
					At: dingtalk.At{
						AtMobiles: []string{"18583872978"},
						IsAtAll:   false,
					},
				})
			}
		}
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
