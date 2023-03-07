package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/forum-api-go/domain"
)

// ========================

type ResponseError struct {
	Message string `json:"message"`
}

type ResponseSuccess struct {
	CreatedUser domain.CreatedUserDto `json:"createdUser"`
}

// ========================

type UserHandler struct {
	UserService domain.UserService
}

// ========================

func (u *UserHandler) New(r *gin.Engine, userService domain.UserService) {
	handler := &UserHandler{
		UserService: userService,
	}
	r.POST("/users", handler.Create)
}

func (u *UserHandler) Create(c *gin.Context) {
	var user domain.CreateUserDto

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{err.Error()})
		return
	}

	createdUser, err := u.UserService.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ResponseSuccess{createdUser})
}
