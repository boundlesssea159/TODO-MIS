package todo

import (
	"TODO-MIS/domain/todo/entity"
	"context"

	"go.uber.org/zap"
)

type Todo struct {
	repository TodoRepository
	logger     *zap.Logger
}

func NewTodo(repository TodoRepository, logger *zap.Logger) *Todo {
	return &Todo{
		repository: repository,
		logger:     logger,
	}
}

func (todo *Todo) Create(ctx context.Context, title, description string, userId int) (int, error) {
	return todo.repository.Create(ctx, title, description, userId)
}

func (todo *Todo) Delete(ctx context.Context, id int, userId int) error {
	return todo.repository.Delete(ctx, id, userId)
}

func (todo *Todo) List(ctx context.Context, userId int) ([]*entity.TodoItem, error) {
	return todo.repository.List(ctx, userId)
}

func (todo *Todo) Complete(ctx context.Context, id int, userId int) error {
	return todo.repository.Complete(ctx, id, userId)
}
