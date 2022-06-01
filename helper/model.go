package helper

import (
	"gin-user-service/model/domain"
	"gin-user-service/model/web/response"
)

func ToUserLoginResponse(user domain.User) response.UserLogin {
	return response.UserLogin{
		Nama:  user.Nama,
		Email: user.Email,
		Token: user.Token,
	}
}

func ToUserRegisterResponse(user domain.User) response.UserRegister {
	return response.UserRegister{
		Nama:  user.Nama,
		Email: user.Email,
	}
}
