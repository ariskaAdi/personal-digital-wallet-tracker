package response

import (
	"ariskaAdi/personal-digital-wallet/internal/model/entity"
	"time"
)

type UserResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`	
}

func NewUserResponse(user entity.Users) UserResponse {
	return UserResponse{
		Id:   int(user.ID),
		Name: user.Name,
		Email: user.Email,
		CreatedAt: user.CreatedAt,

	}
}

func UpdateUserResponse(user entity.Users) UserResponse {
	return UserResponse{
		Name: user.Name,
		Email: user.Email,
		UpdatedAt: user.UpdatedAt,
	}
}

func NewUserResponses(users []entity.Users) []UserResponse {
	res := make([]UserResponse, 0, len(users))
	for _, u := range users {
		res = append(res, NewUserResponse(u))
	}
	return res
}