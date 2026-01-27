package application

import (
	"TODO-MIS/adapter/driving/api/dto"
	"TODO-MIS/domain"
	"context"

	"go.uber.org/zap"
)

type Todo struct {
	service *domain.Todo
	logger  *zap.Logger
}

func NewTodo(service *domain.Todo, logger *zap.Logger) *Todo {
	return &Todo{
		service: service,
		logger:  logger,
	}
}

func (todo *Todo) Create(c context.Context, req *dto.CreateTodoRequest) (int, error) {
	return todo.service.Create(c, req.Title, req.Description)
}

func (todo *Todo) Delete(c context.Context) {

}

func (todo *Todo) List(c context.Context) {

}

func (todo *Todo) Complete(c context.Context) {

}
