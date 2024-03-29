package api

import (
	"fmt"
	"net/http"
	"strings"

	"e.coding.net/tssoft/repository/gomao/dingtalk"
	"e.coding.net/tssoft/repository/gomao/logger"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"github.com/spf13/viper"
)

type ApiHandler func(ctx *Context) (resultCode ResultCode, message string, result interface{})

func JsonApi(handler ApiHandler, resultHandler ...func(result any) any) gin.HandlerFunc {
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
				for _, handler := range resultHandler {
					response["result"] = handler(response["result"])
				}
			}
			ctx.JSON(resultCode.getHttpCode(), response)
		} else {
			ctx.String(resultCode.getHttpCode(), message)
		}
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
						strings.Join(lo.Map(atMobiles, func(item string, _ int) string {
							return fmt.Sprintf("@%s", item)
						}), ", ")),
				},
				At: dingtalk.At{
					AtMobiles: atMobiles,
					IsAtAll:   false,
				},
			})
		}
	}
}
