package auth_service

import (
	"context"

	"github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/core/domain/user"
)

type Api interface {
	Login(ctx context.Context, phoneNumber string) (int, error)
	SubmitOtp(ctx context.Context, phoneNumber string, otpCode int) (*user.User, string, error)
	Authenticate(ctx context.Context, token string) (*user.User, error)
}
