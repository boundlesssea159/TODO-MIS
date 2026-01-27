package _const

import "errors"

var (
	// 1001 ~ 1999 common error code
	InvalidParameterCode = 1001
	JsonParseErrorCode   = 1002
	InternalErrorCode    = 1003
	UnauthorizedCode     = 1004

	// 2000~ business error code
)

var (
	InvalidParameterError = errors.New("invalid parameter")
	JsonParseError        = errors.New("json parse error")
	InternalError         = errors.New("internal server error")
	UnauthorizedError     = errors.New("unauthorized")
)

func GetErrorByCode(code int) error {
	switch code {
	case InvalidParameterCode:
		return InvalidParameterError
	case JsonParseErrorCode:
		return JsonParseError
	case UnauthorizedCode:
		return UnauthorizedError
	default:
		return InternalError
	}
}
