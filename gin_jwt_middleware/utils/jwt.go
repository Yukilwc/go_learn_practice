package utils

import (
	"fmt"
	"time"

	"gjm/global"
	model_system "gjm/model/system"

	"github.com/golang-jwt/jwt/v5"
)

func GetToken(info model_system.TokenInfo, expiresTime int64) (string, error) {
	expiresAt := jwt.NewNumericDate(time.Now().Add(time.Duration(expiresTime) * time.Second))
	fmt.Println("get token设定的expiredAt:", expiresAt.Time.String())
	c := model_system.CustomClaims{
		Info: info,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    global.CONFIG.JWT.Issuer,
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(global.CONFIG.JWT.Secret))
}

func ParseToken(tokenString string) (*model_system.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model_system.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(global.CONFIG.JWT.Secret), nil
	})
	// if err != nil {
	// 	fmt.Println("parse token error:", err, token)
	// 	return nil, err
	// }
	if claims, ok := token.Claims.(*model_system.CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		// 判断下是否是过期导致的
		fmt.Println("token valid error:", err)
		return claims, err
	}

}

func TokenIsExpired(err error) {
}
