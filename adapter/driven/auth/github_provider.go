package auth

import "context"

type GithubProvider struct {
}

func NewGithubProvider() *GithubProvider {
	return &GithubProvider{}
}

func (provider *GithubProvider) GetOAuthURL(ctx context.Context, callbackUrl, channel string) (string, error) {
	// TODO: implement building real GitHub OAuth redirect URL
	return "https://github_redirect_url", nil
}

func (provider *GithubProvider) ExchangeTokenWithCode(ctx context.Context, code, channel string) (string, error) {
	// TODO: implement exchanging GitHub authorization code for access token
	return "token", nil
}
