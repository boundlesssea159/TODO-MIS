package dto

type GetAuthURLRequest struct {
	Channel     string `form:"channel" binding:"required"`
	CallbackUrl string `form:"callback_url" binding:"required"`
}

type GetAuthURLResponse struct {
	OAuthURL string `json:"oauth_url"`
}

type GetTokenWithCodeRequest struct {
	Channel string `form:"channel" binding:"required"`
	Code    string `form:"code" binding:"required"`
}

type GetTokenWithCodeResponse struct {
	AccessToken string `json:"access_token"`
}
