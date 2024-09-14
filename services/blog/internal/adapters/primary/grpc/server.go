package grpc_adapter

import (
	"context"
	"fmt"
	"net"

	auth_service "github.com/tahadostifam/go-hexagonal-architecture/internal/core/services/auth"
	"github.com/tahadostifam/go-hexagonal-architecture/protobuf/auth"
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
	panic("unimplemented")
}

func (a authServerImpl) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	_, err := a.authService.Login(ctx, req.PhoneNumber)
	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{}, nil
}

// SubmitOtp implements auth.AuthServer.
func (a authServerImpl) SubmitOtp(context.Context, *auth.SubmitOtpRequest) (*auth.SubmitOtpResponse, error) {
	panic("unimplemented")
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
