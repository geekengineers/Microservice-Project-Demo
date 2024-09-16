package grpc_adapter

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"github.com/geekengineers/Microservice-Project-Demo/common/interceptor"
	"github.com/geekengineers/Microservice-Project-Demo/protobuf/auth"
	"github.com/geekengineers/Microservice-Project-Demo/protobuf/auth/authconnect"
	grpc_transformer "github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/adapters/primary/grpc/transformer"
	auth_service "github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/core/services/auth"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type App struct {
	authService auth_service.Api
	host        string
	port        int
	mux         *http.ServeMux
}

type authServerImpl struct {
	authService auth_service.Api
}

func NewGrpcServer(authService auth_service.Api, host string, port int) *App {
	loggerInterceptor := interceptor.LoggerInterceptor()

	mux := http.NewServeMux()

	path, handler := authconnect.NewAuthServiceHandler(&authServerImpl{authService}, connect.WithInterceptors(loggerInterceptor))
	mux.Handle(path, handler)

	return &App{authService, host, port, mux}
}

func (a *App) Run() error {
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", a.host, a.port), h2c.NewHandler(a.mux, &http2.Server{}))
	if err != nil {
		return err
	}

	return nil
}

func (a *authServerImpl) Authenticate(ctx context.Context, req *connect.Request[auth.AuthenticateRequest]) (*connect.Response[auth.AuthenticateResponse], error) {
	user, err := a.authService.Authenticate(ctx, req.Msg.AccessToken)
	if err != nil {
		return nil, err
	}

	res := &connect.Response[auth.AuthenticateResponse]{
		Msg: &auth.AuthenticateResponse{
			User: grpc_transformer.DomainToGrpcUser(user),
		},
	}

	return res, nil
}

func (a *authServerImpl) Login(ctx context.Context, req *connect.Request[auth.LoginRequest]) (*connect.Response[auth.LoginResponse], error) {
	_, err := a.authService.Login(ctx, req.Msg.PhoneNumber)
	if err != nil {
		return nil, err
	}

	res := &connect.Response[auth.LoginResponse]{
		Msg: &auth.LoginResponse{},
	}

	return res, nil
}

func (a *authServerImpl) SubmitOtp(ctx context.Context, req *connect.Request[auth.SubmitOtpRequest]) (*connect.Response[auth.SubmitOtpResponse], error) {
	user, accessToken, err := a.authService.SubmitOtp(ctx, req.Msg.PhoneNumber, int(req.Msg.OtpCode))
	if err != nil {
		return nil, err
	}

	res := &connect.Response[auth.SubmitOtpResponse]{
		Msg: &auth.SubmitOtpResponse{
			User:        grpc_transformer.DomainToGrpcUser(user),
			AccessToken: accessToken,
		},
	}

	return res, nil
}
