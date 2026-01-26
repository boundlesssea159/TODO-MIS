package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 定义统一的响应结构
// code: 业务状态码，0 表示成功
// message: 提示信息
// data: 业务数据
// trace_id: 可选的请求追踪 ID（这里先预留字段）
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	TraceID string      `json:"trace_id,omitempty"`
}

// Success 返回统一成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// Error 返回统一错误响应
func Error(c *gin.Context, httpStatus int, code int, message string, data interface{}) {
	c.JSON(httpStatus, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
