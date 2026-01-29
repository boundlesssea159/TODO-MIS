package server

import (
	"TODO-MIS/adapter/driving/api"
	"TODO-MIS/common/middware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, todo *api.Todo, auth *api.Auth) {
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	r.Use(gin.Logger(), gin.Recovery(), middware.AuthMiddleware())
	// business routes
	apiGroup := r.Group("/api/v1")
	apiGroup.POST("/todo-items", todo.Create)
	apiGroup.GET("/todo-items", todo.List)
	apiGroup.DELETE("/todo-items/:id", todo.Delete)
	apiGroup.PATCH("/todo-items/:id/complete", todo.Complete)
	apiGroup.GET("/")
	apiGroup.GET("")

	// OAuth routes
	authGroup := apiGroup.Group("/auth")
	authGroup.GET("/url", auth.GetAuthURL)
	authGroup.GET("/token", auth.GetTokenWithCode)
}
