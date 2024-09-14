package grpc_adapter

import (
	"fmt"
	"net"

	article_service "github.com/tahadostifam/go-hexagonal-architecture/internal/core/services/article"
	"github.com/tahadostifam/go-hexagonal-architecture/protobuf/auth"
	"google.golang.org/grpc"
)

type App struct {
	articleService article_service.Api
	host           string
	port           int
	server         *grpc.Server
}

type articleServerImpl struct {
	auth.UnimplementedAuthServer
	articleService article_service.Api
}

func NewGrpcServer(articleService article_service.Api, host string, port int) *App {
	s := grpc.NewServer()

	auth.RegisterAuthServer(s, articleServerImpl{articleService: articleService})

	return &App{articleService, host, port, s}
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
