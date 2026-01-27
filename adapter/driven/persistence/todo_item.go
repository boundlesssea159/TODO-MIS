package persistence

import "TODO-MIS/domain/entity"

type TodoItem struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string `json:"title" gorm:"size:255;not null;uniqueIndex"`
	Description string `json:"description" gorm:"type:text"`
	Status      int    `json:"status" gorm:"not null;default:0"`
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
	}
}
