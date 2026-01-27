package api

import (
	"TODO-MIS/adapter/driving/api/dto"
	"TODO-MIS/application"
	_const "TODO-MIS/common/const"
	"TODO-MIS/common/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	application *application.Todo
}

func NewTodoAPI(application *application.Todo) *Todo {
	return &Todo{
		application: application,
	}
}

func (todo *Todo) Create(c *gin.Context) {
	req := &dto.CreateTodoRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		util.Fail(c, http.StatusBadRequest, _const.InvalidParameterCode)
		return
	}
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
