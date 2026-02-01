package middware

import (
	_const "TODO-MIS/common/const"
	"TODO-MIS/common/util"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

const (
	UserIDKey = "user_id"
	TokenType = "Bearer"
)

var SecretKey = os.Getenv("SECRET_KEY")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			util.Fail(c, http.StatusUnauthorized, _const.InternalErrorCode)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, TokenType+" ")
		if tokenString == authHeader {
			tokenString = authHeader
		}

		claims, err := ParseToken(tokenString)
		if err != nil {
			util.Fail(c, http.StatusUnauthorized, _const.InternalErrorCode)
			c.Abort()
			return
		}

		// set user id to context
		c.Set(UserIDKey, claims.UserID)
		c.Next()
	}
}

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken generate JWT token
func GenerateToken(userID int) (string, error) {
	claims := Claims{
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

// ParseToken parse JWT token
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// GetUserIDFromContext Get user id from context
func GetUserIDFromContext(c *gin.Context) (int, bool) {
	userID, exists := c.Get(UserIDKey)
	if !exists {
		return 0, false
	}
	return userID.(int), true
}
