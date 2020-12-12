package api

import "strconv"

func (ctx *Context) GetQueryBool(key string) (bool, error) {
	v := ctx.Query(key)
	if v == "" {
		return false, nil
	}
	result, err := strconv.ParseBool(v)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (ctx *Context) GetParamBool(key string) (bool, error) {
	v := ctx.Param(key)
	if v == "" {
		return false, nil
	}
	result, err := strconv.ParseBool(v)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (ctx *Context) GetFormBool(key string) (bool, error) {
	v := ctx.PostForm(key)
	if v == "" {
		return false, nil
	}
	result, err := strconv.ParseBool(v)
	if err != nil {
		return false, err
	}
	return result, nil
}
