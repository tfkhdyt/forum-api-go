package common

import "github.com/tfkhdyt/forum-api-go/domain"

type ResponseWithMessage struct {
	Message string `json:"message"`
}

type CreateUserResponse struct {
	CreatedUser domain.CreatedUserDto `json:"createdUser"`
}
