package http

import (
	"net/http"

	"github.com/tfkhdyt/forum-api-go/common"
	"github.com/tfkhdyt/forum-api-go/domain"

	"github.com/gin-gonic/gin"
)

// ========================

type UserHandler struct {
	userService domain.UserService
}

// ========================

func New(r *gin.Engine, userService domain.UserService) {
	handler := &UserHandler{userService}
	r.POST("/users", handler.Create)
}

// ========================

func (u *UserHandler) Create(c *gin.Context) {
	var user domain.CreateUserDto

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, common.ResponseWithMessage{
			Message: err.Error(),
		})
		return
	}

	createdUser, err := u.userService.Create(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ResponseWithMessage{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, common.CreateUserResponse{
		CreatedUser: createdUser,
	})
}
