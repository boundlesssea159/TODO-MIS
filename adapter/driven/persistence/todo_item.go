package persistence

import (
	"TODO-MIS/domain/todo/entity"
	"time"
)

type TodoItem struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"size:255;not null"`
	Description string    `json:"description" gorm:"type:text"`
	Status      int       `json:"status" gorm:"not null;default:0"`
	UserID      int       `json:"user_id" gorm:"index;not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
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

func (item TodoItem) ToDomainEntity() entity.TodoItem {
	return entity.TodoItem{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		Status:      item.Status,
		UserID:      item.UserID,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}
}
