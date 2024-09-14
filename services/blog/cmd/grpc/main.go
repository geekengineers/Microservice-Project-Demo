package main

import (
	"fmt"

	"github.com/tahadostifam/go-hexagonal-architecture/config"
	grpc_adapter "github.com/tahadostifam/go-hexagonal-architecture/internal/adapters/primary/grpc"
	sqlite_adapter "github.com/tahadostifam/go-hexagonal-architecture/internal/adapters/secondary/sqlite"
	"github.com/tahadostifam/go-hexagonal-architecture/pkg/sms"
	"github.com/tahadostifam/go-hexagonal-architecture/utils"

	"gorm.io/driver/sqlite"
)

func main() {
	cfg := config.Read()

	// Init database dialector
	dialector := sqlite.Open("./database/development.db")

	// Init secondary adapters
	authRepo, err := sqlite_adapter.NewAuthRepositorySecondaryPort(dialector)
	utils.HandleError(err)

	var smsService sms.Service
	if config.CurrentEnv == config.Development {
		smsService = sms.NewSMSDevelopment()
	}

	// Init business logic

	// Init primary adapters
	fmt.Printf("Grpc server is listening at %s:%d\n", cfg.Grpc.Host, cfg.Grpc.Port)
	app := grpc_adapter.NewGrpcServer(authService, cfg.Grpc.Host, cfg.Grpc.Port)
	err = app.Run()
	utils.HandleError(err)
}
