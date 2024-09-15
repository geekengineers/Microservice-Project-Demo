package grpc_transformer

import (
	"github.com/geekengineers/Microservice-Project-Demo/protobuf/auth"
	"github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/core/domain/user"
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
