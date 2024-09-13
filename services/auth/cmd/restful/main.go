package main

import (
	restful_adapter "github.com/tahadostifam/go-hexagonal-architecture/internal/adapters/primary/restful"
	sqlite_adapter "github.com/tahadostifam/go-hexagonal-architecture/internal/adapters/secondary/sqlite"
	user_service "github.com/tahadostifam/go-hexagonal-architecture/internal/core/services/user"
	"github.com/tahadostifam/go-hexagonal-architecture/utils"

	"gorm.io/driver/sqlite"
)

func main() {
	// Init database dialector
	dialector := sqlite.Open("./database/development.db")

	// Init secondary adapters
	userRepo, err := sqlite_adapter.NewUserRepositorySecondaryPort(dialector)
	utils.HandleError(err)

	// Init business logic
	userService := user_service.NewService(userRepo)

	// Init primary adapters
	app := restful_adapter.NewApp(restful_adapter.ServicesApi{
		UserApi: userService,
	})
	err = app.Run()
	utils.HandleError(err)
}
