package domain

import "TODO-MIS/domain/entity"

type TodoRepository interface {
	Create(title string, description string) (int, error)
	Delete(id int) error
	List() ([]*entity.TodoItem, error)
	Complete(id int) error
}
