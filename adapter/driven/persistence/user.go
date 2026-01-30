package persistence

import "TODO-MIS/domain/auth/entity"

type User struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"name"`
	Email string `json:"email" gorm:"email"`
}

func (user *User) ToEntity() *entity.User {
	return &entity.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func (user *User) IsValid() bool {
	return user.ID > 0
}
