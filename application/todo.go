package application

import (
	"TODO-MIS/adapter/driving/api/dto"
	"TODO-MIS/domain/todo"
	"TODO-MIS/domain/todo/entity"
	"context"

	"go.uber.org/zap"
)

type Todo struct {
	service *todo.Todo
	logger  *zap.Logger
}

func NewTodo(service *todo.Todo, logger *zap.Logger) *Todo {
	return &Todo{
		service: service,
		logger:  logger,
	}
}

func (todo *Todo) Create(c context.Context, req *dto.CreateTodoRequest, userId int) (int, error) {
	return todo.service.Create(c, req.Title, req.Description, userId)
}

func (todo *Todo) Delete(c context.Context, id int, userId int) error {
	return todo.service.Delete(c, id, userId)
}

func (todo *Todo) List(c context.Context, userId int) ([]*entity.TodoItem, error) {
	return todo.service.List(c, userId)
}

func (todo *Todo) Complete(c context.Context, id int, userId int) error {
	return todo.service.Complete(c, id, userId)
}
