package api

import "strconv"

func (ctx *Context) GetFormInt64Array(key string) ([]int64, error) {
	result := make([]int64, 0)
	queryResult := ctx.PostFormArray(key)
	for _, qResult := range queryResult {
		v, err := strconv.ParseInt(qResult, 10, 64)
		if err != nil {
			return []int64{}, err
		}
		result = append(result, v)
	}
	return result, nil
}

func (ctx *Context) GetQueryInt64Array(key string) ([]int64, error) {
	result := make([]int64, 0)
	queryResult := ctx.QueryArray(key)
	for _, qResult := range queryResult {
		v, err := strconv.ParseInt(qResult, 10, 64)
		if err != nil {
			return []int64{}, err
		}
		result = append(result, v)
	}
	return result, nil
}

func (ctx *Context) GetFormInt32Array(key string) ([]int32, error) {
	result := make([]int32, 0)
	queryResult := ctx.PostFormArray(key)
	for _, qResult := range queryResult {
		v, err := strconv.ParseInt(qResult, 10, 64)
		if err != nil {
			return []int32{}, err
		}
		result = append(result, int32(v))
	}
	return result, nil
}

func (ctx *Context) GetQueryInt32Array(key string) ([]int32, error) {
	result := make([]int32, 0)
	queryResult := ctx.QueryArray(key)
	for _, qResult := range queryResult {
		v, err := strconv.ParseInt(qResult, 10, 64)
		if err != nil {
			return []int32{}, err
		}
		result = append(result, int32(v))
	}
	return result, nil
}

func (ctx *Context) GetFormFloat64Array(key string) ([]float64, error) {
	result := make([]float64, 0)
	queryResult := ctx.PostFormArray(key)
	for _, qResult := range queryResult {
		v, err := strconv.ParseFloat(qResult, 10)
		if err != nil {
			return []float64{}, err
		}
		result = append(result, v)
	}
	return result, nil
}

func (ctx *Context) GetQueryFloat64Array(key string) ([]float64, error) {
	result := make([]float64, 0)
	queryResult := ctx.QueryArray(key)
	for _, qResult := range queryResult {
		v, err := strconv.ParseFloat(qResult, 10)
		if err != nil {
			return []float64{}, err
		}
		result = append(result, v)
	}
	return result, nil
}

func (ctx *Context) GetFormFloat32Array(key string) ([]float32, error) {
	result := make([]float32, 0)
	queryResult := ctx.PostFormArray(key)
	for _, qResult := range queryResult {
		v, err := strconv.ParseFloat(qResult, 10)
		if err != nil {
			return []float32{}, err
		}
		result = append(result, float32(v))
	}
	return result, nil
}

func (ctx *Context) GetQueryFloat32Array(key string) ([]float32, error) {
	result := make([]float32, 0)
	queryResult := ctx.QueryArray(key)
	for _, qResult := range queryResult {
		v, err := strconv.ParseFloat(qResult, 10)
		if err != nil {
			return []float32{}, err
		}
		result = append(result, float32(v))
	}
	return result, nil
}

func (ctx *Context) GetFormBoolArray(key string) ([]bool, error) {
	result := make([]bool, 0)
	queryResult := ctx.PostFormArray(key)
	for _, qResult := range queryResult {
		v, err := strconv.ParseBool(qResult)
		if err != nil {
			return []bool{}, err
		}
		result = append(result, v)
	}
	return result, nil
}

func (ctx *Context) GetQueryBoolArray(key string) ([]bool, error) {
	result := make([]bool, 0)
	queryResult := ctx.QueryArray(key)
	for _, qResult := range queryResult {
		v, err := strconv.ParseBool(qResult)
		if err != nil {
			return []bool{}, err
		}
		result = append(result, v)
	}
	return result, nil
}
