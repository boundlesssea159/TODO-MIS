package main

import (
	"TODO-MIS/adapter/driven/auth"
	"TODO-MIS/adapter/driven/persistence"
	"TODO-MIS/adapter/driving/api"
	"TODO-MIS/application"
	auth2 "TODO-MIS/domain/auth"
	"TODO-MIS/domain/todo"
	"TODO-MIS/server"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	_ = godotenv.Load(".env")
	app := fx.New(
		fx.Provide(
			// common provider
			server.NewLogger,
			server.NewDB,
			server.NewGinEngine,
			// business provider
			persistence.NewMysqlRepository,
			todo.NewTodo,
			application.NewTodo,
			api.NewTodoAPI,
			// auth provider
			auth.NewOAuthFactory,
			auth2.NewAuthService,
			application.NewAuth,
			api.NewAuth,
		),
		fx.Invoke(
			server.RegisterRoutes,
			startServer,
		),
	)
	app.Run()
}

func startServer(lc fx.Lifecycle, engin *gin.Engine) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	svr := &http.Server{
		Addr:    ":" + port,
		Handler: engin,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := svr.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					panic(fmt.Sprintf("listen: %s\n", err.Error()))
				}
			}()
			fmt.Println("http server start on :" + port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("shutdown http server")
			return svr.Shutdown(ctx)
		},
	})
}
