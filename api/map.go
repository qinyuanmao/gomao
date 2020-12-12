package api

import (
	"strconv"

	"github.com/pkg/errors"
)

func (ctx *Context) GetQueryMapInt32(key string, hashKey string) (int32, error) {
	mp, exist := ctx.GetQueryMap(key)
	if !exist {
		return 0, errors.Errorf("%s not found.", key)
	}
	v, exist := mp[hashKey]
	if !exist {
		return 0, errors.Errorf("%s.%s not found.", key, hashKey)
	}
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return int32(result), nil
}

func (ctx *Context) GetQueryMapInt64(key string, hashKey string) (int64, error) {
	mp, exist := ctx.GetQueryMap(key)
	if !exist {
		return 0, errors.Errorf("%s not found.", key)
	}
	v, exist := mp[hashKey]
	if !exist {
		return 0, errors.Errorf("%s.%s not found.", key, hashKey)
	}
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return int64(result), nil
}

func (ctx *Context) GetFormMapInt32(key string, hashKey string) (int32, error) {
	mp, exist := ctx.GetPostFormMap(key)
	if !exist {
		return 0, errors.Errorf("%s not found.", key)
	}
	v, exist := mp[hashKey]
	if !exist {
		return 0, errors.Errorf("%s.%s not found.", key, hashKey)
	}
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return int32(result), nil
}

func (ctx *Context) GetFormMapInt64(key string, hashKey string) (int64, error) {
	mp, exist := ctx.GetPostFormMap(key)
	if !exist {
		return 0, errors.Errorf("%s not found.", key)
	}
	v, exist := mp[hashKey]
	if !exist {
		return 0, errors.Errorf("%s.%s not found.", key, hashKey)
	}
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return int64(result), nil
}

func (ctx *Context) GetQueryMapFloat32(key string, hashKey string) (float32, error) {
	mp, exist := ctx.GetQueryMap(key)
	if !exist {
		return 0, errors.Errorf("%s not found.", key)
	}
	v, exist := mp[hashKey]
	if !exist {
		return 0, errors.Errorf("%s.%s not found.", key, hashKey)
	}
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseFloat(v, 10)
	if err != nil {
		return 0, err
	}
	return float32(result), nil
}

func (ctx *Context) GetQueryMapFloat64(key string, hashKey string) (float64, error) {
	mp, exist := ctx.GetQueryMap(key)
	if !exist {
		return 0, errors.Errorf("%s not found.", key)
	}
	v, exist := mp[hashKey]
	if !exist {
		return 0, errors.Errorf("%s.%s not found.", key, hashKey)
	}
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseFloat(v, 10)
	if err != nil {
		return 0, err
	}
	return float64(result), nil
}

func (ctx *Context) GetFormMapFloat32(key string, hashKey string) (float32, error) {
	mp, exist := ctx.GetPostFormMap(key)
	if !exist {
		return 0, errors.Errorf("%s not found.", key)
	}
	v, exist := mp[hashKey]
	if !exist {
		return 0, errors.Errorf("%s.%s not found.", key, hashKey)
	}
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseFloat(v, 10)
	if err != nil {
		return 0, err
	}
	return float32(result), nil
}

func (ctx *Context) GetFormMapFloat64(key string, hashKey string) (float64, error) {
	mp, exist := ctx.GetPostFormMap(key)
	if !exist {
		return 0, errors.Errorf("%s not found.", key)
	}
	v, exist := mp[hashKey]
	if !exist {
		return 0, errors.Errorf("%s.%s not found.", key, hashKey)
	}
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseFloat(v, 10)
	if err != nil {
		return 0, err
	}
	return float64(result), nil
}

func (ctx *Context) GetQueryMapBool(key string, hashKey string) (bool, error) {
	mp, exist := ctx.GetQueryMap(key)
	if !exist {
		return false, errors.Errorf("%s not found.", key)
	}
	v, exist := mp[hashKey]
	if !exist {
		return false, errors.Errorf("%s.%s not found.", key, hashKey)
	}
	if v == "" {
		return false, nil
	}
	result, err := strconv.ParseBool(v)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (ctx *Context) GetFormMapBool(key string, hashKey string) (bool, error) {
	mp, exist := ctx.GetPostFormMap(key)
	if !exist {
		return false, errors.Errorf("%s not found.", key)
	}
	v, exist := mp[hashKey]
	if !exist {
		return false, errors.Errorf("%s.%s not found.", key, hashKey)
	}
	if v == "" {
		return false, nil
	}
	result, err := strconv.ParseBool(v)
	if err != nil {
		return false, err
	}
	return result, nil
}
