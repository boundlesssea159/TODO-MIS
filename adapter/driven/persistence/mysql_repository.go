package persistence

import (
	"TODO-MIS/domain"
	"TODO-MIS/domain/entity"
	"context"
)

// todo: title should be set unique index

type MysqlRepository struct {
}

func NewMysqlRepository() domain.TodoRepository {
	return &MysqlRepository{}
}

func (r *MysqlRepository) Create(ctx context.Context, title string, description string) (int, error) {
	return 1, nil
}

func (r *MysqlRepository) Delete(ctx context.Context, id int) error {
	return nil
}

func (r *MysqlRepository) List(ctx context.Context) ([]*entity.TodoItem, error) {
	return []*entity.TodoItem{}, nil
}

func (r *MysqlRepository) Complete(ctx context.Context, id int) error {
	return nil
}
