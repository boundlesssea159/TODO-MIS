package auth

import (
	"TODO-MIS/domain/auth/entity"
	"context"
)

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}
