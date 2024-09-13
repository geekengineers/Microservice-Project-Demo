package redis_adapter

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/ory/dockertest/v3"
	"github.com/redis/go-redis/v9"
)

var (
	redisLock     = &sync.Mutex{}
	redisInstance *redis.Client
)

type Config struct {
	Host, Password string
	Port, DB       int
}

func GetRedisDBInstance(config *Config) *redis.Client {
	if redisInstance == nil {
		redisLock.Lock()
		defer redisLock.Unlock()

		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
			Password: config.Password,
			DB:       config.DB,
		})

		redisInstance = client
	}
	return redisInstance
}

func GetRedisTestInstance(callback func(redisClient *redis.Client)) {
	dockerContainerEnvVariables := []string{}

	err := os.Setenv("ENV", "test")
	if err != nil {
		log.Fatalf("Could not set the environment variable to test: %s", err)
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	var client *redis.Client

	resource, err := pool.Run("redis", "latest", dockerContainerEnvVariables)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	ipAddr := resource.Container.NetworkSettings.IPAddress + ":6379"

	// Kill the container
	defer func() {
		if err = pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}()

	err = pool.Retry(func() error {
		fmt.Printf("Docker redis container network ip address: %s\n", ipAddr)

		client = redis.NewClient(&redis.Options{
			Addr: ipAddr,
			DB:   0,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Fatalf("Could not connect to Redis: %s", err)
	}

	callback(client)
}
