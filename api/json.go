package api

import (
	"io/ioutil"

	"github.com/thoas/go-funk"
	"github.com/tidwall/gjson"
)

func (ctx *Context) GetJsonInt(key string) (int64, error) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return 0, err
	}
	value := gjson.Get(string(body), key)
	return value.Int(), err
}

func (ctx *Context) GetJsonBool(key string) (bool, error) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return false, err
	}
	value := gjson.Get(string(body), key)
	return value.Bool(), err
}

func (ctx *Context) GetJsonFloat(key string) (float64, error) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return 0, err
	}
	value := gjson.Get(string(body), key)
	return value.Float(), err
}

func (ctx *Context) GetJsonString(key string) (string, error) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return "", err
	}
	value := gjson.Get(string(body), key)
	return value.String(), err
}

func (ctx *Context) GetJsonStringArray(key ...string) ([]string, error) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return []string{}, err
	}
	value := gjson.Get(string(body), key)
	return funk.Map(value.Array(), func(value gjson.Result) string {
		return value.String()
	}).([]string), nil
}

func (ctx *Context) GetJsonIntArray(key string) ([]int64, error) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return []int64{}, err
	}
	value := gjson.Get(string(body), key)
	return funk.Map(value.Array(), func(value gjson.Result) int64 {
		return value.Int()
	}).([]int64), nil
}

func (ctx *Context) GetJsonFloatArray(key string) ([]float64, error) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return []float64{}, err
	}
	value := gjson.Get(string(body), key)
	return funk.Map(value.Array(), func(value gjson.Result) float64 {
		return value.Float()
	}).([]float64), nil
}

func (ctx *Context) GetJsonBoolArray(key string) ([]bool, error) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return []bool{}, err
	}
	value := gjson.Get(string(body), key)
	return funk.Map(value.Array(), func(value gjson.Result) bool {
		return value.Bool()
	}).([]bool), nil
}
