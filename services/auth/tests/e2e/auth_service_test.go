package auth_integration_test

import (
	"context"
	"strconv"
	"testing"

	"connectrpc.com/connect"
	"github.com/geekengineers/Microservice-Project-Demo/protobuf/auth"
	grpc_transformer "github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/adapters/primary/grpc/transformer"
	"github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/core/domain/user"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
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

	_, err := client.Login(ctx, &connect.Request[auth.LoginRequest]{
		Msg: &auth.LoginRequest{
			PhoneNumber: s.phoneNumber,
		},
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

	res, err := client.SubmitOtp(ctx, &connect.Request[auth.SubmitOtpRequest]{
		Msg: &auth.SubmitOtpRequest{
			PhoneNumber: s.phoneNumber,
			OtpCode:     int64(s.otpCode),
		},
	})
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res.Msg.User)
	require.NotEmpty(s.T(), res.Msg.AccessToken)

	s.accessToken = res.Msg.AccessToken
	s.user = grpc_transformer.GrpcUserToDomain(res.Msg.User)
}

func (s *AuthServiceSuite) TestC_Authenticate() {
	ctx := context.TODO()

	res, err := client.Authenticate(ctx, &connect.Request[auth.AuthenticateRequest]{
		Msg: &auth.AuthenticateRequest{
			AccessToken: s.accessToken,
		},
	})
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res.Msg.User)
	require.Equal(s.T(), res.Msg.User.Name, s.user.Name)
	require.Equal(s.T(), res.Msg.User.PhoneNumber, s.user.PhoneNumber)
}

func TestAuthServiceSuite(t *testing.T) {
	t.Helper()
	suite.Run(t, new(AuthServiceSuite))
}
