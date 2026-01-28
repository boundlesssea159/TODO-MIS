package auth

import "context"

type OAuthProvider interface {
	GetOAuthURL(ctx context.Context, callbackUrl, channel string) (string, error)
	ExchangeTokenWithCode(ctx context.Context, code, channel string) (string, error)
}
