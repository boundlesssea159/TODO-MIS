package _const

import "errors"

var (
	// 1001 ~ 1999 common error code
	InvalidParameterCode = 1001

	// 2000~2999 business error code
)

var (
	InvalidParameterError = errors.New("invalid parameter")
)

func GetErrorByCode(code int) error {
	switch code {
	case InvalidParameterCode:
		return InvalidParameterError
	default:
		return errors.New("server error")
	}
}
