package auth_integration_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	grpc_transformer "github.com/tahadostifam/go-hexagonal-architecture/internal/adapters/primary/grpc/transformer"
	"github.com/tahadostifam/go-hexagonal-architecture/internal/core/domain/user"
	"github.com/tahadostifam/go-hexagonal-architecture/protobuf/auth"
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

	_, err := client.Login(ctx, &auth.LoginRequest{
		PhoneNumber: s.phoneNumber,
	})
	require.NoError(s.T(), err)

	otpCodeStr, err := otpManager.Get(ctx, s.phoneNumber)
	require.NoError(s.T(), err)

	otpCode, err := strconv.Atoi(otpCodeStr)
	require.NoError(s.T(), err)

	s.otpCode = otpCode
}

func (s *AuthServiceSuite) TestB_SubmitOtp() {
	ctx := context.TODO()

	res, err := client.SubmitOtp(ctx, &auth.SubmitOtpRequest{
		PhoneNumber: s.phoneNumber,
		OtpCode:     int64(s.otpCode),
	})
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res.User)
	require.NotEmpty(s.T(), res.AccessToken)

	s.accessToken = res.AccessToken
	s.user = grpc_transformer.GrpcUserToDomain(res.User)
}

func (s *AuthServiceSuite) TestC_Authenticate() {
	ctx := context.TODO()

	res, err := client.Authenticate(ctx, &auth.AuthenticateRequest{
		AccessToken: s.accessToken,
	})
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res.User)
	require.Equal(s.T(), res.User.Name, s.user.Name)
	require.Equal(s.T(), res.User.PhoneNumber, s.user.PhoneNumber)
}

func TestAuthServiceSuite(t *testing.T) {
	t.Helper()
	suite.Run(t, new(AuthServiceSuite))
}
