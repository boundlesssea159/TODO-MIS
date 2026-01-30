package auth

import (
	"TODO-MIS/domain/auth/entity"
	"context"
)

//go:generate mockgen -destination=mock/mock_repository.go -package=mock . AuthRepository
type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}
