package auth_integration_test

import (
	"fmt"
	"os"
	"testing"

	sqlite_adapter "github.com/tahadostifam/go-hexagonal-architecture/internal/adapters/secondary/sqlite"
	article_service "github.com/tahadostifam/go-hexagonal-architecture/internal/core/services/article"
	"github.com/tahadostifam/go-hexagonal-architecture/utils"
	"gorm.io/driver/sqlite"
)

var articleService article_service.Api

func TestMain(m *testing.M) {
	wd, _ := os.Getwd()
	databaseFilename := fmt.Sprintf("%s/%s", wd, "../../database/test.db ")
	dialector := sqlite.Open(databaseFilename)

	repo, err := sqlite_adapter.NewArticleRepository(dialector)
	utils.HandleError(err)

	requirements := article_service.Requirements{
		Repo: repo,
	}

	articleService = article_service.NewService(&requirements)

	m.Run()

	err = os.Remove(databaseFilename)
	utils.HandleError(err)
}
