package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tfkhdyt/forum-api-go/config"
	"github.com/tfkhdyt/forum-api-go/domain"
)

type jwtTokenService struct{}

func New() domain.TokenService {
	return &jwtTokenService{}
}

func (j *jwtTokenService) CreateAccessToken(userId uint) (string, error) {
	claims := domain.JwtClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "forumAPI",
			Subject:   "accessToken",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(config.GetJwtSecretKey())
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (j *jwtTokenService) CreateRefreshToken(userId uint) (string, error) {
	claims := domain.JwtClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(720 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "forumAPI",
			Subject:   "refreshToken",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(config.GetJwtSecretKey())
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (j *jwtTokenService) DecodePayload(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&domain.JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return config.GetJwtSecretKey(), nil
		},
	)

	if claims, ok := token.Claims.(*domain.JwtClaims); ok && token.Valid {
		return claims.ID, nil
	}

	return "", err
}
