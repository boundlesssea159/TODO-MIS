package api

import (
	"TODO-MIS/common/middware"
	"strconv"

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
	userId, ok := middware.GetUserIDFromContext(c)
	if !ok {
		util.Fail(c, http.StatusUnauthorized, _const.UnauthorizedCode)
		return
	}
	id, err := todo.application.Create(c, req, userId)
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
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		todo.logger.Warn("invalid id parameter", zap.String("id", idStr), zap.Error(err))
		util.Fail(c, http.StatusBadRequest, _const.InvalidParameterCode)
		return
	}

	err = todo.application.Delete(c, id)
	if err != nil {
		todo.logger.Error("Delete todo failed", zap.Int("id", id), zap.Error(err))
		util.Fail(c, http.StatusInternalServerError, _const.InternalErrorCode)
		return
	}
	util.Success(c, nil)
}

func (todo *Todo) List(c *gin.Context) {
	userId, ok := middware.GetUserIDFromContext(c)
	if !ok {
		util.Fail(c, http.StatusUnauthorized, _const.UnauthorizedCode)
		return
	}
	items, err := todo.application.List(c, userId)
	if err != nil {
		todo.logger.Error("List todos failed", zap.Error(err))
		util.Fail(c, http.StatusInternalServerError, _const.InternalErrorCode)
		return
	}

	todoItems := make([]dto.TodoItem, len(items))
	for _, item := range items {
		dtoItem := dto.TodoItem{}
		todoItems = append(todoItems, dtoItem.From(*item))
	}

	util.Success(c, dto.GetTodoListResponse{
		Items: todoItems,
		Total: len(todoItems),
	})
}

func (todo *Todo) Complete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		todo.logger.Warn("Invalid ID parameter", zap.String("id", idStr), zap.Error(err))
		util.Fail(c, http.StatusBadRequest, _const.InvalidParameterCode)
		return
	}

	err = todo.application.Complete(c, id)
	if err != nil {
		todo.logger.Error("Complete todo failed", zap.Int("id", id), zap.Error(err))
		util.Fail(c, http.StatusInternalServerError, _const.InternalErrorCode)
		return
	}
	util.Success(c, nil)
}
