package auth_integration_test

import (
	"fmt"
	"os"
	"testing"

	gorm_adapter "github.com/geekengineers/Microservice-Project-Demo/services/blog/internal/adapters/secondary/gorm"
	article_service "github.com/geekengineers/Microservice-Project-Demo/services/blog/internal/core/services/article"
	"github.com/geekengineers/Microservice-Project-Demo/services/blog/utils"
	"gorm.io/driver/sqlite"
)

var articleService article_service.Api

func TestMain(m *testing.M) {
	wd, _ := os.Getwd()
	databaseFilename := fmt.Sprintf("%s/%s", wd, "../../database/test.db ")
	dialector := sqlite.Open(databaseFilename)

	repo, err := gorm_adapter.NewArticleRepository(dialector)
	utils.HandleError(err)

	requirements := article_service.Requirements{
		Repo: repo,
	}

	articleService = article_service.NewService(&requirements)

	m.Run()

	err = os.Remove(databaseFilename)
	utils.HandleError(err)
}
