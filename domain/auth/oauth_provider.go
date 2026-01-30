package auth

import "context"

//go:generate mockgen -destination=mock/mock_oauth_provider.go -package=mock . OAuthProvider
type OAuthProvider interface {
	GetOAuthURL(ctx context.Context, callbackUrl, channel string) (string, error)
	ExchangeTokenWithCode(ctx context.Context, code, channel string) (string, error)
}
