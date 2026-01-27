package api

import (
	"TODO-MIS/adapter/driving/api/dto"
	"TODO-MIS/application"
	_const "TODO-MIS/common/const"
	"TODO-MIS/common/util"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Todo struct {
	application *application.Todo
	logger      *zap.Logger
}

func NewTodoAPI(application *application.Todo, logger *zap.Logger) *Todo {
	return &Todo{
		application: application,
		logger:      logger,
	}
}

func (todo *Todo) Create(c *gin.Context) {
	req := &dto.CreateTodoRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		todo.logger.Warn("Invalid create todo request", zap.Error(err))
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			util.Fail(c, http.StatusBadRequest, _const.InvalidParameterCode)
			return
		}
		util.Fail(c, http.StatusBadRequest, _const.InternalErrorCode)
		return
	}
	id, err := todo.application.Create(c, req)
	if err != nil {
		todo.logger.Error("Create todo failed", zap.Error(err))
		util.Fail(c, http.StatusInternalServerError, _const.InternalErrorCode)
		return
	}
	util.Success(c, dto.CreateTodoResponse{
		ID: id,
	})
}

func (todo *Todo) Delete(c *gin.Context) {
	id := c.Param("id")
	// 这里可以调用应用层删除逻辑，现在先返回成功占位
	util.Success(c, gin.H{
		"id":      id,
		"deleted": true,
	})
}

func (todo *Todo) List(c *gin.Context) {
	// 示例：返回一个空列表
	util.Success(c, gin.H{
		"items": []interface{}{},
	})
}

func (todo *Todo) Complete(c *gin.Context) {
	id := c.Param("id")
	// 示例：标记完成
	util.Success(c, gin.H{
		"id":        id,
		"completed": true,
	})
}
