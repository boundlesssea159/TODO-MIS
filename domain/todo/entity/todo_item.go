package entity

import (
	_const "TODO-MIS/common/const"
	"time"
)

type TodoItem struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	UserID      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (todoItem TodoItem) ConvertStatus() string {
	return _const.GetTodoItemStatusText(todoItem.Status)
}
