package domain

import (
	"TODO-MIS/domain/entity"
	"context"
)

type TodoRepository interface {
	Create(ctx context.Context, title, description string, userId int) (int, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, userId int) ([]*entity.TodoItem, error)
	Complete(ctx context.Context, id int) error
}
