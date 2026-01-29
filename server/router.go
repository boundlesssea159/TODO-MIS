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

	apiGroup := r.Group("/api/v1")
	// OAuth
	authGroup := apiGroup.Group("/auth")
	authGroup.GET("/url", auth.GetAuthURL)
	authGroup.GET("/token", auth.GetTokenWithCode)

	apiGroup.Use(gin.Logger(), gin.Recovery(), middware.AuthMiddleware())
	// business routes
	todoItemsGroup := apiGroup.Group("/todo-items")
	todoItemsGroup.POST("", todo.Create)
	todoItemsGroup.GET("", todo.List)
	todoItemsGroup.DELETE("/:id", todo.Delete)
	todoItemsGroup.PATCH("/:id/complete", todo.Complete)
}
