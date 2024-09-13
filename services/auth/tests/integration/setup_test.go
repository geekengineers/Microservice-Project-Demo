package auth_integration_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/redis/go-redis/v9"
	auth_manager "github.com/tahadostifam/go-auth-manager"
	redis_adapter "github.com/tahadostifam/go-hexagonal-architecture/internal/adapters/secondary/redis"
	sqlite_adapter "github.com/tahadostifam/go-hexagonal-architecture/internal/adapters/secondary/sqlite"
	auth_service "github.com/tahadostifam/go-hexagonal-architecture/internal/core/services/auth"
	"github.com/tahadostifam/go-hexagonal-architecture/pkg/otp_manager"
	"github.com/tahadostifam/go-hexagonal-architecture/pkg/sms"
	"github.com/tahadostifam/go-hexagonal-architecture/utils"
	"gorm.io/driver/sqlite"
)

var authService auth_service.Api

func TestMain(m *testing.M) {
	wd, _ := os.Getwd()
	databaseFilename := fmt.Sprintf("%s/%s", wd, "../../database/test.db ")
	dialector := sqlite.Open(databaseFilename)

	repo, err := sqlite_adapter.NewAuthRepositorySecondaryPort(dialector)
	utils.HandleError(err)

	redis_adapter.GetRedisTestInstance(func(redisClient *redis.Client) {
		requirements := auth_service.Requirements{
			OtpManager: otp_manager.NewOtpManger(redisClient),
			AuthManager: auth_manager.NewAuthManager(nil, auth_manager.AuthManagerOpts{
				PrivateKey: "no-matter-at-the-moment",
			}),
			Repo:       repo,
			SmsService: sms.NewSMSDevelopment(),
		}

		authService = auth_service.NewService(&requirements)

		m.Run()

		err = os.Remove(databaseFilename)
		utils.HandleError(err)
	})
}
