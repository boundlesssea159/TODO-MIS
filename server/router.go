package server

import (
	"TODO-MIS/adapter/driving/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, todo *api.Todo) {
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	apiGroup := r.Group("/api/v1")
	apiGroup.POST("/todo-items", todo.Create)
	apiGroup.GET("/todo-items", todo.List)
	apiGroup.DELETE("/todo-items/:id", todo.Delete)
	apiGroup.PATCH("/todo-items/:id/complete", todo.Complete)
}
