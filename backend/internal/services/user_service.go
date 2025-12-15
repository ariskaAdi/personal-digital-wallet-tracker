package services

import (
	"ariskaAdi/personal-digital-wallet/internal/dto/request"
	"ariskaAdi/personal-digital-wallet/internal/model/entity"
	"ariskaAdi/personal-digital-wallet/internal/repositories"
	"context"
	"errors"
)

type UserService interface {
	Create(ctx context.Context, req  request.CreateUserRequest) (entity.Users, error)
	Update(ctx context.Context, req request.UpdateUserRequest) (entity.Users, error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]entity.Users, error)
	FindById(ctx context.Context, id int) (entity.Users, error)
}


type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}


func (s *userService) Create(ctx context.Context, req request.CreateUserRequest) (entity.Users, error) {
    if req.Name == "" {
        return entity.Users{}, errors.New("name tidak boleh kosong")
    }
	 if req.Email == "" {
        return entity.Users{}, errors.New("email tidak boleh kosong")
    }
	 if req.Password == "" {
        return entity.Users{}, errors.New("password tidak boleh kosong")
    }

    user := entity.Users{
        Name: req.Name,
		Email: req.Email,
		Password: req.Password,
    }

    return s.repo.Create(ctx, user)
}

func (s *userService) Update(ctx context.Context, req request.UpdateUserRequest) (entity.Users, error) {
    // if req. <= 0 {
    //     return entity.Users{}, errors.New("id tidak valid")
    // }
    if req.Name == "" {
        return entity.Users{}, errors.New("name tidak boleh kosong")
    }
	if req.Email == "" {
        return entity.Users{}, errors.New("email tidak boleh kosong")
    }
	if req.Password == "" {
        return entity.Users{}, errors.New("password tidak boleh kosong")
    }

    user := entity.Users{
        Name: req.Name,
		Email: req.Email,
		Password: req.Password,
    }
    return s.repo.Update(ctx, user)
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
