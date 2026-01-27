package _const

import "errors"

var (
	// 1001 ~ 1999 common error code
	InvalidParameterCode = 1001
	JsonParseErrorCode   = 1002

	// 2000~2999 business error code
)

var (
	InvalidParameterError = errors.New("invalid parameter")
	JsonParseError        = errors.New("json parse error")
)

func GetErrorByCode(code int) error {
	switch code {
	case InvalidParameterCode:
		return InvalidParameterError
	case JsonParseErrorCode:
		return JsonParseError
	default:
		return errors.New("server error")
	}
}
