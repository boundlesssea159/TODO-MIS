package api

import (
	"TODO-MIS/adapter/driving/api/dto"
	"TODO-MIS/application"
	_const "TODO-MIS/common/const"
	"TODO-MIS/common/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	application *application.Auth
}

func NewAuth(application *application.Auth) *Auth {
	return &Auth{
		application: application,
	}
}

func (a *Auth) GetAuthURL(c *gin.Context) {
	var req dto.GetAuthURLRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		util.Fail(c, http.StatusBadRequest, _const.InvalidParameterCode)
		return
	}
	url, err := a.application.GetAuthURL(c.Request.Context(), req.CallbackUrl, req.Channel)
	if err != nil {
		util.Fail(c, http.StatusInternalServerError, _const.InternalErrorCode)
		return
	}

	util.Success(c, dto.GetAuthURLResponse{
		OAuthURL: url,
	})
}

func (a *Auth) GetTokenWithCode(c *gin.Context) {
	var req dto.GetTokenWithCodeRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		util.Fail(c, http.StatusBadRequest, _const.InvalidParameterCode)
		return
	}

	token, err := a.application.GetTokenWithCode(c.Request.Context(), req.Code, req.Channel)
	if err != nil {
		util.Fail(c, http.StatusInternalServerError, _const.InternalErrorCode)
		return
	}

	util.Success(c, dto.GetTokenWithCodeResponse{
		AccessToken: token,
	})
}
