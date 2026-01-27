package domain

import (
	"TODO-MIS/domain/entity"
	"context"
)

type TodoRepository interface {
	Create(ctx context.Context, title string, description string) (int, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]*entity.TodoItem, error)
	Complete(ctx context.Context, id int) error
}
