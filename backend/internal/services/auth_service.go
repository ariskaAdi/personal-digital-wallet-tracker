package services

import (
	"ariskaAdi/personal-digital-wallet/internal/dto/request"
	"ariskaAdi/personal-digital-wallet/internal/dto/response"
	"context"
)

type AuthService interface {
	Register(ctx context.Context, req request.RegisterUserRequest) (response.AuthResponse, error)
	Login(ctx context.Context, req request.LoginUserRequest) (response.LoginResponse, error)
}
