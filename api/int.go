package api

import "strconv"

func (ctx *Context) GetQueryInt32(key string) (int32, error) {
	v := ctx.Query(key)
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return int32(result), nil
}

func (ctx *Context) GetQueryInt64(key string) (int64, error) {
	v := ctx.Query(key)
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (ctx *Context) GetParamInt32(key string) (int32, error) {
	v := ctx.Param(key)
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return int32(result), nil
}

func (ctx *Context) GetParamInt64(key string) (int64, error) {
	v := ctx.Param(key)
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (ctx *Context) GetFormInt(key string) (int, error) {
	v := ctx.PostForm(key)
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(result), nil
}

func (ctx *Context) GetFormInt32(key string) (int32, error) {
	v := ctx.PostForm(key)
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return int32(result), nil
}

func (ctx *Context) GetFormInt64(key string) (int64, error) {
	v := ctx.PostForm(key)
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}
