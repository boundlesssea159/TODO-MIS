package auth

import "context"

type FacebookProvider struct {
}

func NewFacebookProvider() *FacebookProvider {
	return &FacebookProvider{}
}

func (provider *FacebookProvider) GetOAuthURL(ctx context.Context, callbackUrl, channel string) (string, error) {
	// TODO: implement building real Facebook OAuth redirect URL
	return "https://facebook_redirect_url", nil
}

func (provider *FacebookProvider) ExchangeTokenWithCode(ctx context.Context, code, channel string) (string, error) {
	// TODO: implement exchanging Facebook authorization code for access token
	return "token", nil
}
