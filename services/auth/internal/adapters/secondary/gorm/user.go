package gorm_adapter

import (
	"context"

	"github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/core/domain/user"
	"github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/ports"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(dialector gorm.Dialector) (ports.AuthRepositorySecondaryPort, error) {
	db, err := GORM(dialector)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&user.User{})

	return &userRepository{db}, nil
}

func (a *userRepository) Create(ctx context.Context, u *user.User) (*user.User, error) {
	tx := a.db.Model(&user.User{}).Create(&u)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return u, nil
}

func (a *userRepository) Find(ctx context.Context, id int64) (*user.User, error) {
	var u *user.User
	tx := a.db.Model(&user.User{}).Where("id = ?", id).First(&u)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return u, nil
}

func (a *userRepository) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*user.User, error) {
	var u *user.User
	tx := a.db.Model(&user.User{}).Where("phone_number = ?", phoneNumber).First(&u)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return u, nil
}

func (a *userRepository) Update(ctx context.Context, id int64, changes *user.User) (*user.User, error) {
	u, err := a.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	if changes.Name != "" {
		u.Name = changes.Name
	}

	if changes.PhoneNumber != "" {
		u.PhoneNumber = changes.PhoneNumber
	}

	tx := a.db.Model(&user.User{}).Save(&u)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return u, nil
}
