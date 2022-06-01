package service

import (
	"context"
	"errors"
	"gin-user-service/helper"
	"gin-user-service/model/domain"
	"gin-user-service/model/web/request"
	"gin-user-service/model/web/response"
	"gin-user-service/model/web/response/resgrpc"
	"gin-user-service/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ServiceUserImpl struct {
	repositoryUser repository.RepositoryUser
	DB             *gorm.DB
}

// ShowByToken implements ServiceUser
func (service *ServiceUserImpl) ShowByToken(ctx context.Context, tokenGrand string) (resgrpc.UserTokenJwt, error) {
	var resgrpc resgrpc.UserTokenJwt
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.repositoryUser.ShowByToken(ctx, tx, tokenGrand)
	if err != nil {
		return resgrpc, errors.New(err.Error())
	}
	tokenJwt, err := GenerateTokenJwt(user)
	helper.PanicIfError(err)
	resgrpc.TokenJwt = tokenJwt
	return resgrpc, nil

}

// Login implements ServiceUser
func (service *ServiceUserImpl) Login(ctx context.Context, request request.Login) (response.UserLogin, error) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.repositoryUser.Show(ctx, tx, request.Email)
	if err != nil {
		return helper.ToUserLoginResponse(user), errors.New(err.Error())
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
		if err != nil {
			return helper.ToUserLoginResponse(user), errors.New("Username Atau Password Salah")
		}
		// generate token
		var tokenGrand = helper.DoHashUsingSalt(request.Email)
		err := service.repositoryUser.SetToken(ctx, tx, &user, tokenGrand)
		if err != nil {
			return helper.ToUserLoginResponse(user), errors.New(err.Error())
		}
		return helper.ToUserLoginResponse(user), nil
	}
}

// Register implements ServiceUser
func (service *ServiceUserImpl) Register(ctx context.Context, request request.CreateUser) response.UserRegister {
	// begin a transaction
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	// hash password
	password := []byte(request.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	helper.PanicIfError(err)

	user := domain.User{
		Nama:     request.Nama,
		Email:    request.Email,
		Password: string(hashedPassword),
	}
	result := service.repositoryUser.Register(ctx, tx, user)
	return helper.ToUserRegisterResponse(result)
}

func NewServiceUser(repositoryUser repository.RepositoryUser, DB *gorm.DB) ServiceUser {
	return &ServiceUserImpl{
		repositoryUser: repositoryUser,
		DB:             DB,
	}
}
