package api

import (
	"TODO-MIS/application"

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
	// 示例：简单返回一个创建成功的占位数据
	Success(c, gin.H{
		"id":   1,
		"name": "example todo",
	})
}

func (todo *Todo) Delete(c *gin.Context) {
	id := c.Param("id")
	// 这里可以调用应用层删除逻辑，现在先返回成功占位
	Success(c, gin.H{
		"id":      id,
		"deleted": true,
	})
}

func (todo *Todo) List(c *gin.Context) {
	// 示例：返回一个空列表
	Success(c, gin.H{
		"items": []interface{}{},
	})
}

func (todo *Todo) Complete(c *gin.Context) {
	id := c.Param("id")
	// 示例：标记完成
	Success(c, gin.H{
		"id":        id,
		"completed": true,
	})
}
