package domain

import (
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

func (todo *Todo) Create(ctx context.Context, title string, description string) (int, error) {
	return todo.repository.Create(ctx, title, description)
}
