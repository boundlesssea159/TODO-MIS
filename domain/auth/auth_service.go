package auth

import (
	"TODO-MIS/common/middware"
	"context"
)

type AuthService struct {
	oauthProviderFactory OAuthProvider
	authRepository       AuthRepository
}

func NewAuthService(oauthProviderFactory OAuthProvider, authRepository AuthRepository) *AuthService {
	return &AuthService{
		oauthProviderFactory: oauthProviderFactory,
		authRepository:       authRepository,
	}
}

func (service *AuthService) GetAuthURL(c context.Context, callbackUrl, channel string) (string, error) {
	return service.oauthProviderFactory.GetOAuthURL(c, callbackUrl, channel)
}
func (service *AuthService) GetTokenWithCode(c context.Context, code, channel string) (string, error) {
	return service.oauthProviderFactory.ExchangeTokenWithCode(c, code, channel)
}

func (service *AuthService) GenerateJWT(ctx context.Context, email, code string) (string, error) {
	user, err := service.authRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	// todo call third party to verify code and get third party token
	token, err := middware.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
