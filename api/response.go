package api

import (
	"fmt"
	"net/http"
)

func SuccessResponse(response interface{}) (httpCode, resultCode int, message string, result interface{}) {
	return http.StatusOK, SUCCESS, "Success", response
}

func WithParamNotFound(paramName string) (httpCode, resultCode int, message string, result interface{}) {
	return http.StatusBadRequest, PARAMS_NOT_FOUNT, fmt.Sprintf("%s parameter not found!", paramName), nil
}

func WithParamError(paramName, errorMessage string) (httpCode, resultCode int, message string, result interface{}) {
	return http.StatusBadRequest, PARAMS_ERROR, fmt.Sprintf("%s parameter not error: %s!", paramName, errorMessage), nil
}

func WithLogout() (httpCode, resultCode int, message string, result interface{}) {
	return http.StatusUnauthorized, NOLOGIN, fmt.Sprintf("You are logout."), nil
}

func WithServerError(err error) (httpCode, resultCode int, message string, result interface{}) {
	return http.StatusInternalServerError, SERVER_ERROR, err.Error(), nil
}