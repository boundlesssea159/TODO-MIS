package main

import (
	"TODO-MIS/adapter/driving/api"
	"TODO-MIS/application"
	"TODO-MIS/domain"
	"TODO-MIS/server"
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
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

func startServer(lc fx.Lifecycle, engin *gin.Engine) {
	svr := &http.Server{
		Addr:    "localhost:8080",
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
