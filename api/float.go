package api

import "strconv"

func (ctx *Context) GetQueryFloat(key string) (float32, error) {
	v := ctx.Query(key)
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseFloat(v, 10)
	if err != nil {
		return 0, err
	}
	return float32(result), nil
}

func (ctx *Context) GetQueryFloat64(key string) (float64, error) {
	v := ctx.Query(key)
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseFloat(v, 10)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (ctx *Context) GetParamFloat(key string) (float32, error) {
	v := ctx.Param(key)
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseFloat(v, 10)
	if err != nil {
		return 0, err
	}
	return float32(result), nil
}

func (ctx *Context) GetParamFloat64(key string) (float64, error) {
	v := ctx.Param(key)
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseFloat(v, 10)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (ctx *Context) GetFormFloat32(key string) (float32, error) {
	v := ctx.PostForm(key)
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseFloat(v, 10)
	if err != nil {
		return 0, err
	}
	return float32(result), nil
}

func (ctx *Context) GetFormFloat64(key string) (float64, error) {
	v := ctx.PostForm(key)
	if v == "" {
		return 0, nil
	}
	result, err := strconv.ParseFloat(v, 10)
	if err != nil {
		return 0, err
	}
	return result, nil
}
