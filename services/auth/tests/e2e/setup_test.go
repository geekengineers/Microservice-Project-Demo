package auth_integration_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/geekengineers/Microservice-Project-Demo/protobuf/auth/authconnect"
	"github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/adapters/primary"
	redis_adapter "github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/adapters/secondary/redis"
	"github.com/geekengineers/Microservice-Project-Demo/services/auth/pkg/otp_manager"
	"github.com/geekengineers/Microservice-Project-Demo/services/auth/utils"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/sqlite"
)

var client authconnect.AuthServiceClient
var otpManager otp_manager.OtpManager

func TestMain(m *testing.M) {
	err := os.Setenv("GO_ENV", "test")
	utils.HandleError(err)

	redis_adapter.GetRedisTestInstance(func(redisClient *redis.Client) {
		const grpc_port = 8006

		dialector := sqlite.Open("../../database/test.db")

		// Bootstrap Application (Test ENV)
		go func() {
			primary.Bootstrap(&primary.BootstrapRequirements{
				RedisClient: redisClient,
				Dialector:   dialector,
				Grpc: struct {
					Host string
					Port int
				}{
					Host: "127.0.0.1",
					Port: grpc_port,
				},
				Jwt: struct{ PrivateKey string }{
					PrivateKey: "samplejwtsecret",
				},
			})
		}()

		otpManager = *otp_manager.NewOtpManger(redisClient)

		client = authconnect.NewAuthServiceClient(http.DefaultClient, fmt.Sprintf("%s:%d", "http://127.0.0.1", grpc_port))

		time.Sleep(2 * time.Second)

		m.Run()
	})
}
