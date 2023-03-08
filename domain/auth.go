package domain

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ==================

type Auth struct {
	gorm.Model
	RefreshToken string `json:"refreshToken" gorm:"unique"`
}

type LoginLogoutDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type Credentials struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type DeleteRefreshAuthDto struct {
	RefreshToken string `json:"refreshToken" binding:"required,jwt"`
}

// ==================

type AuthHandler interface {
	Post(r *gin.Context)
	Patch(r *gin.Context)
	Delete(r *gin.Context)
}

type AuthService interface {
	Login(loginDto LoginLogoutDto) (Credentials, error)
	Logout(logoutDto LoginLogoutDto) error
	RefreshToken(refreshDto DeleteRefreshAuthDto) (string, error)
	DeleteToken(deleteAuthDto DeleteRefreshAuthDto) error
}

type AuthRepository interface {
	CreateToken(token string) error
	CheckTokenAvailability(token string) error
	DeleteToken(token string) error
}
