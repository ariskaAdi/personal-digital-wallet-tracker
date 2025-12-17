package services

import (
	"ariskaAdi/personal-digital-wallet/internal/dto/request"
	"ariskaAdi/personal-digital-wallet/internal/dto/response"
	"ariskaAdi/personal-digital-wallet/internal/model/entity"
	"ariskaAdi/personal-digital-wallet/internal/repositories"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}


func (s *userService) Update(ctx context.Context, req request.UpdateUserRequest, id int) (response.AuthResponse, error) {
	if id <= 0 {
	    return response.AuthResponse{}, errors.New("id tidak valid")
	}

	if req.Name == "" || req.Email == "" || req.Password == ""{
		return response.AuthResponse{}, errors.New("field tidak boleh kosong")
	}

		// 3. Hash password
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return response.AuthResponse{}, err
	}

	user := entity.Users{
		ID:       uint(id),
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashed),
	}
	updated, err := s.repo.Update(ctx, user)
	if err != nil {
		return response.AuthResponse{}, err
	}

	return response.AuthResponse{
		Name:  updated.Name,
		Email: updated.Email,
	}, nil
}

func (s *userService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("id tidak valid")
	}
	return s.repo.Delete(ctx, id)
}

func (s *userService) FindAll(ctx context.Context) ([]entity.Users, error) {
	return s.repo.FindAll(ctx)
}

func (s *userService) FindById(ctx context.Context, id int) (entity.Users, error) {
	if id <= 0 {
		return entity.Users{}, errors.New("id tidak valid")
	}

	user, err := s.repo.FindById(ctx, id)
	if err != nil {
		return entity.Users{}, errors.New("user tidak ditemukan")
	}

	return user, nil
}
