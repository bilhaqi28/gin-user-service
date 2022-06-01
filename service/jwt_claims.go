package service

import (
	"errors"
	"gin-user-service/model/domain"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type M map[string]interface{}

var APPLICATION_NAME = "My Simple JWT App"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("belajarmicroservice")

type MyClaims struct {
	jwt.StandardClaims
	Nama  string `json:"Nama"`
	Email string `json:"Email"`
	Id    int    `Json:"Id"`
}

func newClaims(user domain.User) MyClaims {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Nama:  user.Nama,
		Email: user.Email,
		Id:    int(user.ID),
	}
	return claims
}

func GenerateTokenJwt(user domain.User) (string, error) {
	claims := newClaims(user)
	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)
	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return signedToken, nil
}
