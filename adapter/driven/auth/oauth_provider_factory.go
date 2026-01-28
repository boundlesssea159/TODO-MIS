package auth

import (
	"TODO-MIS/domain/auth"
	"context"
)

type OAuthProviderFactory struct {
}

func NewOAuthProvider() *OAuthProviderFactory {
	return &OAuthProviderFactory{}
}

func (factory *OAuthProviderFactory) getProvider(channel string) auth.OAuthProvider {
	switch channel {
	case "google":
		return NewGmailProvider()
	case "github":
		return NewGithubProvider()
	case "facebook":
		return NewFacebookProvider()
	default:
		return nil
	}
}

func (factory *OAuthProviderFactory) GetOAuthURL(ctx context.Context, callbackUrl, channel string) (string, error) {
	provider := factory.getProvider(channel)
	return provider.GetOAuthURL(ctx, callbackUrl, channel)
}

func (factory *OAuthProviderFactory) ExchangeTokenWithCode(ctx context.Context, code, channel string) (string, error) {
	provider := factory.getProvider(channel)
	return provider.ExchangeTokenWithCode(ctx, code, channel)
}
