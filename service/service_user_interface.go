package service

import (
	"context"

	"gin-user-service/model/web/request"
	"gin-user-service/model/web/response"
	"gin-user-service/model/web/response/resgrpc"
)

type ServiceUser interface {
	Register(ctx context.Context, request request.CreateUser) response.UserRegister
	Login(ctx context.Context, request request.Login) (response.UserLogin, error)
	ShowByToken(ctx context.Context, tokenGrand string) (resgrpc.UserTokenJwt, error)
}
