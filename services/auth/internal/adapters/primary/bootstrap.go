package primary

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	auth_manager "github.com/tahadostifam/go-auth-manager"
	"github.com/tahadostifam/go-hexagonal-architecture/config"
	grpc_adapter "github.com/tahadostifam/go-hexagonal-architecture/internal/adapters/primary/grpc"
	sqlite_adapter "github.com/tahadostifam/go-hexagonal-architecture/internal/adapters/secondary/sqlite"
	auth_service "github.com/tahadostifam/go-hexagonal-architecture/internal/core/services/auth"
	"github.com/tahadostifam/go-hexagonal-architecture/pkg/otp_manager"
	"github.com/tahadostifam/go-hexagonal-architecture/pkg/sms"
	"github.com/tahadostifam/go-hexagonal-architecture/utils"
	"gorm.io/gorm"
)

type BootstrapRequirements struct {
	Grpc struct {
		Host string
		Port int
	}
	Jwt struct {
		PrivateKey string
	}
	RedisClient *redis.Client
	Dialector   gorm.Dialector
}

func Bootstrap(requirements *BootstrapRequirements) {
	// Init secondary adapters
	authRepo, err := sqlite_adapter.NewUserRepository(requirements.Dialector)
	utils.HandleError(err)

	var smsService sms.Service
	if config.CurrentEnv != config.Production {
		smsService = sms.NewSMSDevelopment()
	}

	// Init business logic
	authService := auth_service.NewService(&auth_service.Requirements{
		Repo:       authRepo,
		OtpManager: otp_manager.NewOtpManger(requirements.RedisClient),
		AuthManager: auth_manager.NewAuthManager(nil, auth_manager.AuthManagerOpts{
			PrivateKey: requirements.Jwt.PrivateKey,
		}),
		SmsService: smsService,
	})

	// Init primary adapters
	fmt.Printf("Grpc server is listening at %s:%d\n", requirements.Grpc.Host, requirements.Grpc.Port)
	app := grpc_adapter.NewGrpcServer(authService, requirements.Grpc.Host, requirements.Grpc.Port)
	err = app.Run()
	utils.HandleError(err)
}
