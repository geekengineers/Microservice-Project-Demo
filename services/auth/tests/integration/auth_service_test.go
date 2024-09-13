package auth_integration_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/tahadostifam/go-hexagonal-architecture/internal/core/domain/user"
)

type AuthServiceSuite struct {
	suite.Suite

	phoneNumber string
	otpCode     int
	accessToken string
	user        *user.User
}

func (s *AuthServiceSuite) SetupSuite() {
	s.phoneNumber = "+989345678900"
}

func (s *AuthServiceSuite) TestA_Login() {
	ctx := context.TODO()

	otpCode, err := authService.Login(ctx, s.phoneNumber)
	require.NoError(s.T(), err)
	require.NotZero(s.T(), otpCode)

	s.otpCode = otpCode

	s.T().Log("Otp Code", s.otpCode)
}

func (s *AuthServiceSuite) TestB_SubmitOtp() {
	ctx := context.TODO()

	user, accessToken, err := authService.SubmitOtp(ctx, s.phoneNumber, s.otpCode)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), user)
	require.NotEmpty(s.T(), accessToken)

	s.accessToken = accessToken
	s.user = user

	s.T().Log("AccessToken", s.accessToken)
}

func (s *AuthServiceSuite) TestC_Authenticate() {
	ctx := context.TODO()

	user, err := authService.Authenticate(ctx, s.accessToken)
	require.NoError(s.T(), err)
	require.Equal(s.T(), user, s.user)
}

func TestAuthServiceSuite(t *testing.T) {
	t.Helper()
	suite.Run(t, new(AuthServiceSuite))
}
