package grpc_transformer

import (
	"github.com/tahadostifam/go-hexagonal-architecture/internal/core/domain/user"
	"github.com/tahadostifam/go-hexagonal-architecture/protobuf/auth"
)

func GrpcUserToDomain(u *auth.User) *user.User {
	return &user.User{
		Name:        u.Name,
		PhoneNumber: u.PhoneNumber,
	}
}

func DomainToGrpcUser(u *user.User) *auth.User {
	return &auth.User{
		Name:        u.Name,
		PhoneNumber: u.PhoneNumber,
	}
}
