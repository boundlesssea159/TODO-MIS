package persistence

import (
	_const "TODO-MIS/common/const"
	"TODO-MIS/domain/todo"
	"TODO-MIS/domain/todo/entity"
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MysqlRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewMysqlRepository(db *gorm.DB, logger *zap.Logger) todo.TodoRepository {
	return &MysqlRepository{db: db, logger: logger}
}

func (r *MysqlRepository) Create(ctx context.Context, title, description string, userId int) (int, error) {
	item := &TodoItem{
		Title:       title,
		Description: description,
		Status:      _const.TodoItemCreatedStatus,
		UserID:      userId,
	}
	if err := r.db.WithContext(ctx).Create(item).Error; err != nil {
		r.logger.Error("create item error", zap.Error(err))
		return 0, err
	}
	return item.ID, nil
}

func (r *MysqlRepository) Delete(ctx context.Context, id int) error {
	result := r.db.WithContext(ctx).Model(&TodoItem{}).Where("id = ?", id).Update("status", _const.TodoItemDeletedStatus)
	err := result.Error
	if err != nil {
		r.logger.Error("delete item status error", zap.Error(err))
	}
	return err
}

func (r *MysqlRepository) List(ctx context.Context, userId int) ([]*entity.TodoItem, error) {
	var items []*TodoItem
	if err := r.db.WithContext(ctx).
		Where("user_id=? and status!=?", userId, _const.TodoItemDeletedStatus).
		Order("created_at desc").
		Find(&items).Error; err != nil {
		return nil, err
	}

	result := make([]*entity.TodoItem, 0, len(items))
	for _, item := range items {
		domainEntity := item.ToDomainEntity()
		result = append(result, &domainEntity)
	}
	return result, nil
}

func (r *MysqlRepository) Complete(ctx context.Context, id int) error {
	err := r.db.WithContext(ctx).Model(&TodoItem{}).
		Where("id = ? and user_id=?", id).
		Update("status", _const.TodoItemDoneStatus).Error
	if err != nil {
		r.logger.Error("complete item status error", zap.Error(err))
	}
	return err
}
