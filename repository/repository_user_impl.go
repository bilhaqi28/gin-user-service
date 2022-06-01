package repository

import (
	"context"
	"errors"
	"gin-user-service/helper"
	"gin-user-service/model/domain"

	"gorm.io/gorm"
)

type RepositoryUserImpl struct {
}

// ShowByToken implements RepositoryUser
func (*RepositoryUserImpl) ShowByToken(ctx context.Context, tx *gorm.DB, token string) (domain.User, error) {
	var user domain.User
	result := tx.First(&user, "token = ?", token)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, errors.New("Token Is Not Match")
	} else {
		return user, nil
	}
}

// SetToken implements RepositoryUser
func (*RepositoryUserImpl) SetToken(ctx context.Context, tx *gorm.DB, user *domain.User, token string) error {
	user.Token = token
	result := tx.Model(&user).Select("token").Updates(user)
	if result.Error != nil {
		return errors.New("Failed Set Token")
	} else {
		return nil
	}
}

// Show implements RepositoryUser
func (*RepositoryUserImpl) Show(ctx context.Context, tx *gorm.DB, email string) (domain.User, error) {
	var user domain.User
	result := tx.First(&user, "email = ?", email)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, errors.New("User Not Found")
	} else {
		return user, nil
	}
}

// Login implements RepositoryUser

// Register implements RepositoryUser
func (*RepositoryUserImpl) Register(ctx context.Context, tx *gorm.DB, user domain.User) domain.User {
	result := tx.Create(&user)
	helper.PanicIfError(result.Error)
	return user
}

func NewRepositoryUser() RepositoryUser {
	return &RepositoryUserImpl{}
}
