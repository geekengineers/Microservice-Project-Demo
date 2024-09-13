package main

import (
	"fmt"

	auth_manager "github.com/tahadostifam/go-auth-manager"
	"github.com/tahadostifam/go-hexagonal-architecture/config"
	grpc_adapter "github.com/tahadostifam/go-hexagonal-architecture/internal/adapters/primary/grpc"
	redis_adapter "github.com/tahadostifam/go-hexagonal-architecture/internal/adapters/secondary/redis"
	sqlite_adapter "github.com/tahadostifam/go-hexagonal-architecture/internal/adapters/secondary/sqlite"
	auth_service "github.com/tahadostifam/go-hexagonal-architecture/internal/core/services/auth"
	"github.com/tahadostifam/go-hexagonal-architecture/pkg/otp_manager"
	"github.com/tahadostifam/go-hexagonal-architecture/pkg/sms"
	"github.com/tahadostifam/go-hexagonal-architecture/utils"

	"gorm.io/driver/sqlite"
)

func main() {
	cfg := config.Read()

	// Init database dialector
	dialector := sqlite.Open("./database/development.db")

	redisClient := redis_adapter.GetRedisDBInstance(&redis_adapter.Config{
		Host:     cfg.Redis.Host,
		Port:     cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	// Init secondary adapters
	authRepo, err := sqlite_adapter.NewAuthRepositorySecondaryPort(dialector)
	utils.HandleError(err)

	var smsService sms.Service
	if config.CurrentEnv == config.Development {
		smsService = sms.NewSMSDevelopment()
	}

	// Init business logic
	authService := auth_service.NewService(&auth_service.Requirements{
		Repo:       authRepo,
		OtpManager: otp_manager.NewOtpManger(redisClient),
		AuthManager: auth_manager.NewAuthManager(nil, auth_manager.AuthManagerOpts{
			PrivateKey: cfg.Jwt.PrivateKey,
		}),
		SmsService: smsService,
	})

	// Init primary adapters
	fmt.Printf("Grpc server is listening at %s:%d\n", cfg.Grpc.Host, cfg.Grpc.Port)
	app := grpc_adapter.NewGrpcServer(authService, cfg.Grpc.Host, cfg.Grpc.Port)
	err = app.Run()
	utils.HandleError(err)
}
