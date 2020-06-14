package helper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/qinyuanmao/go-utils/logutl"
)

func GetParam(ctx *gin.Context, key string) string {
	var value string
	value = ctx.PostForm(key)
	if value == "" {
		value = ctx.Query(key)
	}
	if value == "" {
		value = ctx.Param(key)
	}
	if value == "" {
		var values map[string]string
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		buf := bytes.NewBuffer(body)
		ctx.Request.Body = ioutil.NopCloser(buf)
		_ = json.Unmarshal(body, &values)
		value = values[key]
	}
	return strings.TrimSpace(value)
}

func GetIntParam(ctx *gin.Context, key string) int {
	vStr := GetParam(ctx, key)
	if vStr == "" {
		var values map[string]int
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		buf := bytes.NewBuffer(body)
		ctx.Request.Body = ioutil.NopCloser(buf)
		_ = json.Unmarshal(body, &values)
		return values[key]
	} else {
		v, err := strconv.Atoi(vStr)
		if err != nil {
			logutl.Error(err.Error())
		}
		return v
	}
}

func GetIntArrayParam(ctx *gin.Context, key string) []int {
	vStr := GetParam(ctx, key)
	if vStr == "" {
		var values map[string][]int
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		buf := bytes.NewBuffer(body)
		ctx.Request.Body = ioutil.NopCloser(buf)
		_ = json.Unmarshal(body, &values)
		v := values[key]
		if v == nil {
			var values map[string]string
			body, _ := ioutil.ReadAll(ctx.Request.Body)
			buf := bytes.NewBuffer(body)
			ctx.Request.Body = ioutil.NopCloser(buf)
			_ = json.Unmarshal(body, &values)
			v2 := values[key]
			if v2 == "" {
				return nil
			} else {
				var r []int
				err := json.Unmarshal([]byte(v2), &r)
				if err != nil {
					logutl.Error(err.Error())
				}
				return r
			}
		} else {
			return v
		}
	} else {
		var r []int
		err := json.Unmarshal([]byte(vStr), &r)
		if err != nil {
			logutl.Error(err.Error())
		}
		return r
	}
}

func GetInt64Param(ctx *gin.Context, key string) int64 {
	vStr := GetParam(ctx, key)
	if vStr == "" {
		var values map[string]int64
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		buf := bytes.NewBuffer(body)
		ctx.Request.Body = ioutil.NopCloser(buf)
		_ = json.Unmarshal(body, &values)
		return values[key]
	} else {
		v, err := strconv.ParseInt(vStr, 10, 64)
		if err != nil {
			logutl.Error(err.Error())
		}
		return v
	}
}

func GetFloat64Param(ctx *gin.Context, key string) float64 {
	vStr := GetParam(ctx, key)
	if vStr == "" {
		var values map[string]float64
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		buf := bytes.NewBuffer(body)
		ctx.Request.Body = ioutil.NopCloser(buf)
		_ = json.Unmarshal(body, &values)
		return values[key]
	} else {
		v, err := strconv.ParseFloat(vStr, 64)
		if err != nil {
			logutl.Error(err.Error())
		}
		return v
	}
}

func GetBoolParam(ctx *gin.Context, key string) bool {
	vStr := GetParam(ctx, key)
	if vStr == "" {
		var values map[string]bool
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		buf := bytes.NewBuffer(body)
		ctx.Request.Body = ioutil.NopCloser(buf)
		_ = json.Unmarshal(body, &values)
		return values[key]
	} else {
		v, err := strconv.ParseBool(vStr)
		if err != nil {
			logutl.Error(err.Error())
		}
		return v
	}
}

func Bind(ctx *gin.Context, val interface{}) (err error) {
	if err = ctx.ShouldBindWith(val, binding.Default(ctx.Request.Method, ctx.Request.Header.Get("Content-Type"))); err != nil {
		logutl.Error(err)
	}
	return
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, openId, Authorization, Token, sign, timestamp")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
