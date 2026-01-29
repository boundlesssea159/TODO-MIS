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
	ErrInvalidParameter = errors.New("invalid parameter")
	ErrJsonParse        = errors.New("json parse error")
	ErrInternal         = errors.New("internal server error")
	ErrUnauthorized     = errors.New("unauthorized")
)

func GetErrorByCode(code int) error {
	switch code {
	case InvalidParameterCode:
		return ErrInvalidParameter
	case JsonParseErrorCode:
		return ErrJsonParse
	case UnauthorizedCode:
		return ErrUnauthorized
	default:
		return ErrInternal
	}
}
