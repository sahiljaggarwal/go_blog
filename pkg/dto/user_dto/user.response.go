package user_dto

import (
	"blog-api/models"
)

type UserResponse struct {
	Id uint `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password,omitempty"`
}

func ToPublicResponse(user *models.User) UserResponse {
	return UserResponse{
		Id: user.ID,
		Email: user.Email,
		Username: user.Username,
		Password: user.Password,
	}
}

func ToPrivateResponse(user *models.User) UserResponse {
	return UserResponse{
		Id: user.ID,
		Email: user.Email,
		Username: user.Username,
	}
}