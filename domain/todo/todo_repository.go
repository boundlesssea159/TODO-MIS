package todo

import (
	"TODO-MIS/domain/todo/entity"
	"context"
)

//go:generate mockgen -destination=mock/mock_todo_repository.go -package=mock . TodoRepository
type TodoRepository interface {
	Create(ctx context.Context, title, description string, userId int) (int, error)
	Delete(ctx context.Context, id int, userId int) error
	List(ctx context.Context, userId int) ([]*entity.TodoItem, error)
	Complete(ctx context.Context, id int, userId int) error
}
