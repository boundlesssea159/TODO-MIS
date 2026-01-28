package application

import "github.com/gin-gonic/gin"

type Auth struct {
}

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) GetAuthURL(c *gin.Context) {

}
func (a *Auth) GetTokenWithCode(c *gin.Context) {}
