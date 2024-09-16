package primary

import (
	"fmt"

	grpc_adapter "github.com/geekengineers/Microservice-Project-Demo/services/blog/internal/adapters/primary/grpc"
	gorm_adapter "github.com/geekengineers/Microservice-Project-Demo/services/blog/internal/adapters/secondary/gorm"
	article_service "github.com/geekengineers/Microservice-Project-Demo/services/blog/internal/core/services/article"
	"github.com/geekengineers/Microservice-Project-Demo/services/blog/utils"
	"gorm.io/gorm"
)

type BootstrapRequirements struct {
	Grpc struct {
		Host           string
		Port           int
		AuthServiceUrl string
	}
	Dialector gorm.Dialector
}

func Bootstrap(requirements *BootstrapRequirements) {
	// Init secondary adapters
	articleRepo, err := gorm_adapter.NewArticleRepository(requirements.Dialector)
	utils.HandleError(err)

	// Init business logic
	articleService := article_service.NewService(&article_service.Requirements{
		Repo: articleRepo,
	})

	// Init primary adapters
	fmt.Printf("Grpc server is listening at %s:%d\n", requirements.Grpc.Host, requirements.Grpc.Port)
	app := grpc_adapter.NewGrpcServer(articleService, requirements.Grpc.AuthServiceUrl, requirements.Grpc.Host, requirements.Grpc.Port)
	err = app.Run()
	utils.HandleError(err)
}
