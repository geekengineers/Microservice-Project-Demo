package ports

import (
	"context"
	"errors"

	"github.com/tahadostifam/go-hexagonal-architecture/internal/core/domain/user"
)

var (
	ErrUserNotFound = errors.New("user does not exist")
)

type AuthRepositorySecondaryPort interface {
	Create(ctx context.Context, user *user.User) (*user.User, error)
	Update(ctx context.Context, id int64, changes *user.User) (*user.User, error)
	Find(ctx context.Context, id int64) (*user.User, error)
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (*user.User, error)
}
