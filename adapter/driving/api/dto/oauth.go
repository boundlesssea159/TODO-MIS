package dto

type GetAuthURLRequest struct {
	Channel     string `json:"channel" binding:"required"`
	CallbackUrl string `json:"callback_url" binding:"required"`
}

type GetAuthURLResponse struct {
	OAuthURL string `json:"oauth_url"`
}

type GetTokenWithCodeRequest struct {
	Channel string `json:"channel" binding:"required"`
	Code    string `json:"code" binding:"required"`
}

type GetTokenWithCodeResponse struct {
	AccessToken string `json:"access_token"`
}
