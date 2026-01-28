package auth

import (
	"context"
)

type AuthService struct {
	oauthProvider OAuthProvider
}

func NewAuthService(oauthProvider OAuthProvider) *AuthService {
	return &AuthService{
		oauthProvider: oauthProvider,
	}
}

func (service *AuthService) GetAuthURL(c context.Context, callbackUrl, channel string) (string, error) {
	return service.oauthProvider.GetOAuthURL(c, callbackUrl, channel)
}
func (service *AuthService) GetTokenWithCode(c context.Context, code, channel string) (string, error) {
	return service.oauthProvider.ExchangeTokenWithCode(c, code, channel)
}
