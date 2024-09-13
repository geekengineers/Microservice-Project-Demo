package restful_adapter

import (
	user_service "github.com/tahadostifam/go-hexagonal-architecture/internal/core/services/auth"
)

type ServicesApi struct {
	UserApi user_service.Api
}

type App struct {
	port        int
	servicesApi ServicesApi
}

func NewGrpcServer(servicesApi ServicesApi) *App {
	s := &App{

		servicesApi: servicesApi,
		port:        8000,
	}

	return s
}

func (a *App) Run() error {
	return nil
}
