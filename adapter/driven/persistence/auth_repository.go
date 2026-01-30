package persistence

import (
	_const "TODO-MIS/common/const"
	auth2 "TODO-MIS/domain/auth"
	"TODO-MIS/domain/auth/entity"
	"context"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth2.AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (repo *AuthRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user User
	if err := repo.db.WithContext(ctx).
		Where("email=?", email).
		Find(&user).Error; err != nil {
		return nil, err
	}
	if !user.IsValid() {
		return nil, _const.ErrDataNotFound
	}
	return user.ToEntity(), nil
}
