package main

import (
	"TODO-MIS/adapter/driven/persistence"
	"TODO-MIS/adapter/driving/api"
	"TODO-MIS/application"
	"TODO-MIS/domain"
	"TODO-MIS/server"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	_ = godotenv.Load(".env")
	app := fx.New(
		fx.Provide(
			NewLogger,
			persistence.NewMysqlRepository,
			domain.NewTodo,
			application.NewTodo,
			api.NewTodoAPI,
			server.NewGinEngine,
		),
		fx.Invoke(
			server.RegisterRoutes,
			startServer,
		),
	)
	app.Run()
}

func NewLogger() (*zap.Logger, error) {
	if os.Getenv("APP_ENV") == "prod" {
		return zap.NewProduction()
	}
	return zap.NewDevelopment()
}

func startServer(lc fx.Lifecycle, engin *gin.Engine) {
	port := os.Getenv("PORT")
	svr := &http.Server{
		Addr:    "localhost:" + port,
		Handler: engin,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := svr.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					panic(fmt.Sprintf("listen: %s\n", err.Error()))
				}
			}()
			fmt.Println("http server start on localhost:8080")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("shutdown http server")
			return svr.Shutdown(ctx)
		},
	})
}
