package services

import (
	"ariskaAdi/personal-digital-wallet/internal/dto/request"
	"ariskaAdi/personal-digital-wallet/internal/model/entity"
	"context"
)

type UserService interface {
	Create(ctx context.Context, req  request.CreateUserRequest) (entity.Users, error)
	Update(ctx context.Context, req request.UpdateUserRequest) (entity.Users, error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]entity.Users, error)
	FindById(ctx context.Context, id int) (entity.Users, error)
}


