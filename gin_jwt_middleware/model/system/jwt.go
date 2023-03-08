package model_system

import "github.com/golang-jwt/jwt/v5"

type TokenInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type CustomClaims struct {
	Info TokenInfo `json:"info"`
	jwt.RegisteredClaims
}
