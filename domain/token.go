package domain

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	UserId uint `json:"userId"`
	jwt.RegisteredClaims
}

type TokenService interface {
	CreateAccessToken(userId uint) (string, error)
	CreateRefreshToken(userId uint) (string, error)
	DecodePayload(token string) (uint, error)
}
