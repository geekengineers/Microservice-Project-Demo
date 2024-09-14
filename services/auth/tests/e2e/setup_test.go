package auth_integration_test

import (
	"fmt"
	"testing"

	"github.com/tahadostifam/go-hexagonal-architecture/config"
	redis_adapter "github.com/tahadostifam/go-hexagonal-architecture/internal/adapters/secondary/redis"
	"github.com/tahadostifam/go-hexagonal-architecture/pkg/otp_manager"
	"github.com/tahadostifam/go-hexagonal-architecture/protobuf/auth"
	"github.com/tahadostifam/go-hexagonal-architecture/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client auth.AuthClient
var otpManager otp_manager.OtpManager

func TestMain(m *testing.M) {
	cfg := config.Read()

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", cfg.Grpc.Host, cfg.Grpc.Port), grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	utils.HandleError(err)
	defer conn.Close()

	redisClient := redis_adapter.GetRedisDBInstance(&redis_adapter.Config{
		Host:     cfg.Redis.Host,
		Port:     cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	otpManager = *otp_manager.NewOtpManger(redisClient)

	client = auth.NewAuthClient(conn)

	m.Run()
}
