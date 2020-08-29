package api

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/qinyuanmao/gomao/dingtalk"
	"github.com/qinyuanmao/gomao/logger"
	"github.com/spf13/viper"
)

var instance *dingtalk.Client
var once sync.Once

func getInstance() *dingtalk.Client {
	once.Do(func() {
		instance = dingtalk.NewClient(viper.GetString("dingding_webhook"))
	})
	return instance
}

type ApiHandler func(ctx *gin.Context) (httpCode, resultCode int, message string, result interface{})

func JsonApi(handler ApiHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httpCode, resultCode, message, result := handler(ctx)
		if httpCode != http.StatusOK && httpCode != http.StatusBadRequest {
			logger.Error(message)
			webhook := viper.GetString("dingding_webhook")
			env := viper.Get("env")
			if webhook == "" && env == "release" {
				pc, _, line, _ := runtime.Caller(1)
				f := runtime.FuncForPC(pc)
				getInstance().Notify(&dingtalk.DingTalkMsg{
					MsgType: "监控报警",
					Markdown: dingtalk.Markdown{
						Title: fmt.Sprintf("接口请求异常: %d", httpCode),
						Text:  fmt.Sprintf("%s:%d", message, f, line),
					},
					Text: dingtalk.Text{
						Content: fmt.Sprintf("错误信息: %s", message),
					},
					At: dingtalk.At{
						AtMobiles: []string{"18583872978"},
					},
				})
			}
		}
		ctx.JSON(httpCode, map[string]interface{}{
			"code":    resultCode,
			"message": message,
			"result":  result,
		})
	}
}
