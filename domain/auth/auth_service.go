package auth

import (
	"context"
)

type AuthService struct {
	oauthProviderFactory OAuthProvider
}

func NewAuthService(oauthProviderFactory OAuthProvider) *AuthService {
	return &AuthService{
		oauthProviderFactory: oauthProviderFactory,
	}
}

func (service *AuthService) GetAuthURL(c context.Context, callbackUrl, channel string) (string, error) {
	return service.oauthProviderFactory.GetOAuthURL(c, callbackUrl, channel)
}
func (service *AuthService) GetTokenWithCode(c context.Context, code, channel string) (string, error) {
	return service.oauthProviderFactory.ExchangeTokenWithCode(c, code, channel)
}
