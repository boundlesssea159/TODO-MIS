package persistence

import (
	"TODO-MIS/domain/todo/entity"
)

type TodoItem struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string `json:"title" gorm:"size:255;not null"`
	Description string `json:"description" gorm:"type:text"`
	Status      int    `json:"status" gorm:"not null;default:0"`
	UserID      int    `json:"user_id" gorm:"index;not null"`
	CreatedAt   int64  `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt   int64  `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}

func (item TodoItem) From(todoItem entity.TodoItem) TodoItem {
	return TodoItem{
		ID:          todoItem.ID,
		Title:       todoItem.Title,
		Description: todoItem.Description,
		Status:      todoItem.Status,
		UserID:      todoItem.UserID,
	}
}
