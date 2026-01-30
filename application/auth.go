package application

import (
	"TODO-MIS/domain/auth"
	"context"
)

type Auth struct {
	authService *auth.AuthService
}

func NewAuth(authService *auth.AuthService) *Auth {
	return &Auth{
		authService: authService,
	}
}

func (a *Auth) GetAuthURL(c context.Context, callbackUrl, channel string) (string, error) {
	return a.authService.GetAuthURL(c, callbackUrl, channel)
}

func (a *Auth) GetTokenWithCode(c context.Context, code, channel string) (string, error) {
	return a.authService.GetTokenWithCode(c, code, channel)
}

func (a *Auth) GenerateJWT(ctx context.Context, email, code string) (string, error) {
	return a.authService.GenerateJWT(ctx, email, code)
}
