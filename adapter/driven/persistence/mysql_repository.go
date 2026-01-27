package persistence

import (
	"TODO-MIS/domain"
	"TODO-MIS/domain/entity"
)

type MysqlRepository struct {
}

func NewMysqlRepository() domain.TodoRepository {
	return &MysqlRepository{}
}

func (r *MysqlRepository) Create(title string, description string) (int, error) {
	return 1, nil
}

func (r *MysqlRepository) Delete(id int) error {
	return nil
}

func (r *MysqlRepository) List() ([]*entity.TodoItem, error) {
	return []*entity.TodoItem{}, nil
}

func (r *MysqlRepository) Complete(id int) error {
	return nil
}
