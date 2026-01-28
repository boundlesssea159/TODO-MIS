package entity

import _const "TODO-MIS/common/const"

type TodoItem struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	UserID      int    `json:"user_id"`
}

func (todoItem TodoItem) ConvertStatus() string {
	return _const.GetTodoItemStatusText(todoItem.Status)
}
