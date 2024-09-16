package primary

import (
	"fmt"

	"github.com/geekengineers/Microservice-Project-Demo/services/auth/config"
	grpc_adapter "github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/adapters/primary/grpc"
	gorm_adapter "github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/adapters/secondary/gorm"
	auth_service "github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/core/services/auth"
	"github.com/geekengineers/Microservice-Project-Demo/services/auth/pkg/otp_manager"
	"github.com/geekengineers/Microservice-Project-Demo/services/auth/pkg/sms"
	"github.com/geekengineers/Microservice-Project-Demo/services/auth/utils"
	"github.com/redis/go-redis/v9"
	auth_manager "github.com/tahadostifam/go-auth-manager"
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
	authRepo, err := gorm_adapter.NewUserRepository(requirements.Dialector)
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
