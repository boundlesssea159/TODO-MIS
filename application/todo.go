package application

import (
	"TODO-MIS/domain"
	"context"
)

type Todo struct {
	service *domain.Todo
}

func NewTodo(service *domain.Todo) *Todo {
	return &Todo{
		service: service,
	}
}

func (todo *Todo) Create(c context.Context) {
}

func (todo *Todo) Delete(c context.Context) {

}

func (todo *Todo) List(c context.Context) {

}

func (todo *Todo) Complete(c context.Context) {

}
