package api

import (
	"fmt"
)

func SuccessResponse(response interface{}) (code ResultCode, message string, result interface{}) {
	return SUCCESS, "成功！", response
}

func WithParamNotFound(paramName string) (code ResultCode, message string, result interface{}) {
	return PARAMS_NOT_FOUNT, fmt.Sprintf("请求必须添加 %s 参数", paramName), nil
}

func WithParamError(paramName, errorMessage string) (code ResultCode, message string, result interface{}) {
	return PARAMS_ERROR, fmt.Sprintf("%s 参数不正确: %s!", paramName, errorMessage), nil
}

func WithLogout(errMessage ...string) (code ResultCode, message string, result interface{}) {
	if errMessage != nil && len(errMessage) > 0 {
		return NOLOGIN, errMessage[0], nil
	}
	return NOLOGIN, fmt.Sprintf("用户未登录。"), nil
}

func WithServerError(err error) (code ResultCode, message string, result interface{}) {
	return SERVER_ERROR, err.Error(), nil
}

func WithRecordNotFound(errMessage ...string) (code ResultCode, message string, result interface{}) {
	if errMessage != nil && len(errMessage) > 0 {
		return RECORD_NOT_FOUND, errMessage[0], nil
	}
	return RECORD_NOT_FOUND, "未找到记录。", nil
}

func WithForbidden() (code ResultCode, message string, result interface{}) {
	return FORBIDDEN, "权限不足！", nil
}

func WithResponseError(err error, response interface{}) (code ResultCode, message string, result interface{}) {
	if err != nil {
		return WithServerError(err)
	}
	return SuccessResponse(response)
}

func WithRequestError(err error) (code ResultCode, message string, result interface{}) {
	return PARAMS_ERROR, err.Error(), nil
}
