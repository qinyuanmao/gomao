package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"e.coding.net/tssoft/repository/gomao/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Any("/*action", api.JsonProxyRequestHandler(func(ctx *gin.Context) *httputil.ReverseProxy {
		proxy, err := api.ProxyHandler("https://f.3bsoft.cn", func(res *http.Request) {
			res.Header.Add("X-Test", "test")
		}, func(resp *http.Response) error {
			fmt.Printf("response: %v\n", resp)
			return nil
		})

		if err != nil {
			log.Panic(err)
		}
		return proxy
	}))
	router.Run(":8080")
}
