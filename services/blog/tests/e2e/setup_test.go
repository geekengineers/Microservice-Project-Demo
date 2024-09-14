package auth_integration_test

import (
	"fmt"
	"testing"

	"github.com/tahadostifam/go-hexagonal-architecture/config"
	"github.com/tahadostifam/go-hexagonal-architecture/protobuf/auth"
	"github.com/tahadostifam/go-hexagonal-architecture/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client auth.AuthClient

func TestMain(m *testing.M) {
	cfg := config.Read()

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", cfg.Grpc.Host, cfg.Grpc.Port), grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	utils.HandleError(err)
	defer conn.Close()

	// TODO - Check auth service to know how e2e and grpc server bootstrap (primary.Bootstrap) works

	client = auth.NewAuthClient(conn)

	m.Run()
}
