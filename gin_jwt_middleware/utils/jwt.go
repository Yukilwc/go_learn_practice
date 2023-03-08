package utils

import (
	"errors"
	"time"

	model_system "gjm/model/system"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("mysecret")

func GetToken(info model_system.TokenInfo, expiresTime time.Duration) (string, error) {
	c := model_system.CustomClaims{
		Info: info,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresTime)),
			Issuer:    "project name",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(secret)
}

func ParseToken(tokenString string) (*model_system.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model_system.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model_system.CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
