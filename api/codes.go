package api

import "net/http"

type ResultCode int

const (
	SUCCESS ResultCode = iota
	PARAMS_NOT_FOUNT
	PARAMS_ERROR
	SERVER_ERROR
	NOLOGIN
	RECORD_NOT_FOUND
	FORBIDDEN
	TIME_OUT
)

func (code ResultCode) getHttpCode() int {
	switch code {
	case SUCCESS:
		return http.StatusOK
	case PARAMS_NOT_FOUNT:
		return http.StatusBadRequest
	case PARAMS_ERROR:
		return http.StatusBadRequest
	case SERVER_ERROR:
		return http.StatusInternalServerError
	case NOLOGIN:
		return http.StatusUnauthorized
	case RECORD_NOT_FOUND:
		return http.StatusNotFound
	case FORBIDDEN:
		return http.StatusForbidden
	case TIME_OUT:
		return http.StatusRequestTimeout
	}
	return http.StatusNotFound
}
