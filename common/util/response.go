package util

import (
	_const "TODO-MIS/common/const"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"` // code: 0 success
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	TraceID string      `json:"trace_id,omitempty"` // trace_id: optionally request trace ID (reserved field)
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func Fail(c *gin.Context, httpStatus int, code int) {
	c.JSON(httpStatus, Response{
		Code:    code,
		Message: _const.GetErrorByCode(code).Error(),
	})
}
