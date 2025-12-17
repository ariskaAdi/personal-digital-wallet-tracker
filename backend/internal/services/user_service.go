package services

import (
	"ariskaAdi/personal-digital-wallet/internal/dto/request"
	"ariskaAdi/personal-digital-wallet/internal/dto/response"
	"ariskaAdi/personal-digital-wallet/internal/model/entity"
	"context"
)

type UserService interface {
	Update(ctx context.Context, req request.UpdateUserRequest, id int) (response.AuthResponse, error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]entity.Users, error)
	FindById(ctx context.Context, id int) (entity.Users, error)
}


