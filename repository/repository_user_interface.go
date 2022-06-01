package repository

import (
	"context"

	"gin-user-service/model/domain"

	"gorm.io/gorm"
)

type RepositoryUser interface {
	Register(ctx context.Context, tx *gorm.DB, user domain.User) domain.User
	Show(ctx context.Context, tx *gorm.DB, email string) (domain.User, error)
	SetToken(ctx context.Context, tx *gorm.DB, user *domain.User, token string) error
	ShowByToken(ctx context.Context, tx *gorm.DB, token string) (domain.User, error)
}
