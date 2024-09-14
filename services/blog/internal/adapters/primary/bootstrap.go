package primary

import (
	"fmt"

	grpc_adapter "github.com/tahadostifam/go-hexagonal-architecture/internal/adapters/primary/grpc"
	sqlite_adapter "github.com/tahadostifam/go-hexagonal-architecture/internal/adapters/secondary/sqlite"
	article_service "github.com/tahadostifam/go-hexagonal-architecture/internal/core/services/article"
	"github.com/tahadostifam/go-hexagonal-architecture/utils"
	"gorm.io/gorm"
)

type BootstrapRequirements struct {
	Grpc struct {
		Host string
		Port int
	}
	Dialector gorm.Dialector
}

func Bootstrap(requirements *BootstrapRequirements) {
	// Init secondary adapters
	articleRepo, err := sqlite_adapter.NewArticleRepository(requirements.Dialector)
	utils.HandleError(err)

	// Init business logic
	articleService := article_service.NewService(&article_service.Requirements{
		Repo: articleRepo,
	})

	// Init primary adapters
	fmt.Printf("Grpc server is listening at %s:%d\n", requirements.Grpc.Host, requirements.Grpc.Port)
	app := grpc_adapter.NewGrpcServer(articleService, requirements.Grpc.Host, requirements.Grpc.Port)
	err = app.Run()
	utils.HandleError(err)
}
