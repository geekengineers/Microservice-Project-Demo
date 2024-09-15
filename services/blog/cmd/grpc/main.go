package main

import (
	"github.com/geekengineers/Microservice-Project-Demo/services/blog/config"
	"github.com/geekengineers/Microservice-Project-Demo/services/blog/internal/adapters/primary"

	"gorm.io/driver/sqlite"
)

func main() {
	cfg := config.Read()

	// Init database dialector
	dialector := sqlite.Open("./database/development.db")

	primary.Bootstrap(&primary.BootstrapRequirements{
		Grpc: struct {
			Host           string
			Port           int
			AuthServiceUrl string
		}{
			Host:           cfg.Grpc.Host,
			Port:           cfg.Grpc.Port,
			AuthServiceUrl: cfg.Grpc.AuthServiceUrl,
		},
		Dialector: dialector,
	})
}
