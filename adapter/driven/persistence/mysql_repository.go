package persistence

import (
	"TODO-MIS/domain"
	"TODO-MIS/domain/entity"
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MysqlRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewMysqlRepository(db *gorm.DB, logger *zap.Logger) domain.TodoRepository {
	return &MysqlRepository{db: db, logger: logger}
}

func (r *MysqlRepository) Create(ctx context.Context, title string, description string) (int, error) {
	item := &TodoItem{
		Title:       title,
		Description: description,
	}
	if err := r.db.WithContext(ctx).Create(item).Error; err != nil {
		r.logger.Error("error creating item", zap.Error(err))
		return 0, err
	}
	return item.ID, nil
}

func (r *MysqlRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&entity.TodoItem{}, id).Error
}

func (r *MysqlRepository) List(ctx context.Context) ([]*entity.TodoItem, error) {
	var items []*entity.TodoItem
	if err := r.db.WithContext(ctx).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *MysqlRepository) Complete(ctx context.Context, id int) error {
	err := r.db.WithContext(ctx).Model(&entity.TodoItem{}).
		Where("id = ?", id).
		Update("completed", true).Error
	if err != nil {
		r.logger.Error("error updating item", zap.Error(err))
	}
	return err
}
