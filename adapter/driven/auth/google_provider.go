package auth

import "context"

type GmailProvider struct {
}

func NewGmailProvider() *GmailProvider {
	return &GmailProvider{}
}

func (provider *GmailProvider) GetOAuthURL(ctx context.Context, callbackUrl, channel string) (string, error) {
	return "https://google_redirect_url", nil
}

func (provider *GmailProvider) ExchangeTokenWithCode(ctx context.Context, code, channel string) (string, error) {
	return "token", nil
}
