package application

import "TODO-MIS/domain"

type Todo struct {
	service *domain.Todo
}

func NewTodo(service *domain.Todo) *Todo {
	return &Todo{
		service: service,
	}
}
