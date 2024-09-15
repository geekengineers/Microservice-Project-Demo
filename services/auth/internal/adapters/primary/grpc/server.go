package grpc_adapter

import (
	"context"
	"fmt"
	"net"

	"github.com/geekengineers/Microservice-Project-Demo/protobuf/auth"
	grpc_transformer "github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/adapters/primary/grpc/transformer"
	auth_service "github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/core/services/auth"
	"google.golang.org/grpc"
)

type App struct {
	authService auth_service.Api
	host        string
	port        int
	server      *grpc.Server
}

type authServerImpl struct {
	auth.UnimplementedAuthServer
	authService auth_service.Api
}

func (a authServerImpl) Authenticate(ctx context.Context, req *auth.AuthenticateRequest) (*auth.AuthenticateResponse, error) {
	user, err := a.authService.Authenticate(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}

	res := &auth.AuthenticateResponse{
		User: grpc_transformer.DomainToGrpcUser(user),
	}

	return res, nil
}

func (a authServerImpl) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	_, err := a.authService.Login(ctx, req.PhoneNumber)
	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{}, nil
}

func (a authServerImpl) SubmitOtp(ctx context.Context, req *auth.SubmitOtpRequest) (*auth.SubmitOtpResponse, error) {
	user, accessToken, err := a.authService.SubmitOtp(ctx, req.PhoneNumber, int(req.OtpCode))
	if err != nil {
		return nil, err
	}

	res := &auth.SubmitOtpResponse{
		User:        grpc_transformer.DomainToGrpcUser(user),
		AccessToken: accessToken,
	}

	return res, nil
}

func NewGrpcServer(authService auth_service.Api, host string, port int) *App {
	s := grpc.NewServer()

	auth.RegisterAuthServer(s, authServerImpl{authService: authService})

	return &App{authService, host, port, s}
}

func (a *App) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.host, a.port))
	if err != nil {
		return err
	}

	err = a.server.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}
