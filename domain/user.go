package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
}

type CreateUserDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
	FullName string `json:"fullname" binding:"required"`
}

type CreatedUserDto struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
}

type UserService interface {
	CreateUser(createUserDto CreateUserDto) (CreatedUserDto, error)
}

type UserRepository interface {
	CreateUser(createUserDto CreateUserDto) (CreatedUserDto, error)
	VerifyAvailableUsername(username string) error
	FindPasswordByUsername(username string) (string, error)
	FindIdByUsername(username string) (uint, error)
}
