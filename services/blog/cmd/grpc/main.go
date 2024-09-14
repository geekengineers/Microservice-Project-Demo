package main

import (
	"github.com/tahadostifam/go-hexagonal-architecture/config"
	"github.com/tahadostifam/go-hexagonal-architecture/internal/adapters/primary"

	"gorm.io/driver/sqlite"
)

func main() {
	cfg := config.Read()

	// Init database dialector
	dialector := sqlite.Open("./database/development.db")

	primary.Bootstrap(&primary.BootstrapRequirements{
		Grpc: struct {
			Host string
			Port int
		}{
			Host: cfg.Grpc.Host,
			Port: cfg.Grpc.Port,
		},
		Dialector: dialector,
	})
}
