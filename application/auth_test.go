package application

import (
	"TODO-MIS/domain/auth"
	"TODO-MIS/domain/auth/mock"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type AuthTestSuite struct {
	suite.Suite
	ctrl         *gomock.Controller
	mockProvider *mock.MockOAuthProvider
	authApp      *Auth
	ctx          context.Context
	authService  *auth.AuthService
}

func (s *AuthTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.ctrl = gomock.NewController(s.T())
	s.mockProvider = mock.NewMockOAuthProvider(s.ctrl)
	s.authService = auth.NewAuthService(s.mockProvider)
	s.authApp = &Auth{
		authService: s.authService,
	}
}

func (s *AuthTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *AuthTestSuite) TestGetAuthURL_Success() {
	callbackUrl := "http://localhost:8080/callback"
	channel := "google"
	expectedURL := "https://accounts.google.com/o/oauth/authorize?client_id=xxx&redirect_uri=http://localhost:8080/callback&response_type=code&scope=email profile"

	s.mockProvider.EXPECT().GetOAuthURL(s.ctx, callbackUrl, channel).Return(expectedURL, nil)
	url, err := s.authApp.GetAuthURL(s.ctx, callbackUrl, channel)
	s.Nil(err)
	s.Equal(expectedURL, url)
}

func (s *AuthTestSuite) TestGetAuthURL_Fail() {
	callbackUrl := "http://localhost:8080/callback"
	channel := "google"

	s.mockProvider.EXPECT().GetOAuthURL(s.ctx, callbackUrl, channel).Return("", errors.New("provider error"))
	url, err := s.authApp.GetAuthURL(s.ctx, callbackUrl, channel)
	s.NotNil(err)
	s.Equal("", url)
	s.Equal("provider error", err.Error())
}

func (s *AuthTestSuite) TestGetTokenWithCode_Success() {
	code := "oauth_code_12345"
	channel := "google"
	expectedToken := "access_token_67890"

	s.mockProvider.EXPECT().ExchangeTokenWithCode(s.ctx, code, channel).Return(expectedToken, nil)
	token, err := s.authApp.GetTokenWithCode(s.ctx, code, channel)
	s.Nil(err)
	s.Equal(expectedToken, token)
}

func (s *AuthTestSuite) TestGetTokenWithCode_Fail() {
	code := "oauth_code_12345"
	channel := "google"

	s.mockProvider.EXPECT().ExchangeTokenWithCode(s.ctx, code, channel).Return("", errors.New("exchange failed"))
	token, err := s.authApp.GetTokenWithCode(s.ctx, code, channel)
	s.NotNil(err)
	s.Equal("", token)
	s.Equal("exchange failed", err.Error())
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}
