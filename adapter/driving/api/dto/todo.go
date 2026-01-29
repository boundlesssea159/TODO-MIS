package dto

import (
	"TODO-MIS/domain/todo/entity"
)

type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"max=256,omitempty"`
}

type CreateTodoResponse struct {
	ID int `json:"id"`
}

type UpdateTodoRequest struct {
	Title       string `json:"title" binding:"omitempty,max=50"`
	Description string `json:"description" binding:"max=256,omitempty"`
	Status      int    `json:"status" binding:"omitempty,min=0,max=2"`
}

type UpdateTodoResponse struct {
	ID int `json:"id"`
}

type GetTodoListResponse struct {
	Items []TodoItem `json:"items"`
	Total int        `json:"total"`
}

type TodoItem struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (TodoItem) From(todoItem entity.TodoItem) TodoItem {
	return TodoItem{
		ID:          todoItem.ID,
		Title:       todoItem.Title,
		Description: todoItem.Description,
		Status:      todoItem.ConvertStatus(),
		CreatedAt:   todoItem.CreatedAt.String(),
		UpdatedAt:   todoItem.UpdatedAt.String(),
	}
}
