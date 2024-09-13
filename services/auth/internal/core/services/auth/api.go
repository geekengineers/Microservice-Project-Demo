package auth_service

import (
	"context"

	"github.com/tahadostifam/go-hexagonal-architecture/internal/core/domain/user"
)

type Api interface {
	Login(ctx context.Context, phoneNumber string) (int, error)
	SubmitOtp(ctx context.Context, phoneNumber string, otpCode int) (*user.User, string, error)
	Authenticate(ctx context.Context, token string) (*user.User, error)
}
