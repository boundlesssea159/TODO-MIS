package server

import (
	"TODO-MIS/adapter/driving/api"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, todo *api.Todo) {
	apiGroup := r.Group("/api/v1")
	apiGroup.POST("/todo-items", todo.Create)
	apiGroup.DELETE("/todo-items/:id", todo.Delete)
	apiGroup.GET("/todo-items", todo.List)
	apiGroup.PATCH("/todo-items/:id/complete", todo.Complete)
}
