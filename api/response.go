package api

import (
	"fmt"
)

func SuccessResponse(response interface{}) (code ResultCode, message string, result interface{}) {
	return SUCCESS, "Success", response
}

func WithParamNotFound(paramName string) (code ResultCode, message string, result interface{}) {
	return PARAMS_NOT_FOUNT, fmt.Sprintf("%s parameter is required!", paramName), nil
}

func WithParamError(paramName, errorMessage string) (code ResultCode, message string, result interface{}) {
	return PARAMS_ERROR, fmt.Sprintf("%s parameter not error: %s!", paramName, errorMessage), nil
}

func WithLogout() (code ResultCode, message string, result interface{}) {
	return NOLOGIN, fmt.Sprintf("You are logout."), nil
}

func WithServerError(err error) (code ResultCode, message string, result interface{}) {
	return SERVER_ERROR, err.Error(), nil
}

func WithRecordNotFound() (code ResultCode, message string, result interface{}) {
	return RECORD_NOT_FOUND, "Record not found.", nil
}

func WithForbidden() (code ResultCode, message string, result interface{}) {
	return FORBIDDEN, "Forbidden", nil
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
