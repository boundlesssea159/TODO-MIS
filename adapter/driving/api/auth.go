package api

import (
	"TODO-MIS/application"

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

}
func (a *Auth) GetTokenWithCode(c *gin.Context) {}
