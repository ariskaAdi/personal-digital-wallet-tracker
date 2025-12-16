package services

import (
	"ariskaAdi/personal-digital-wallet/internal/dto/request"
	"ariskaAdi/personal-digital-wallet/internal/dto/response"
	"ariskaAdi/personal-digital-wallet/internal/model/entity"
	"ariskaAdi/personal-digital-wallet/internal/repositories"
	"ariskaAdi/personal-digital-wallet/internal/utils"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepo repositories.UserRepository
	jwt      utils.JWTService
}

func NewAuthService(
	userRepo repositories.UserRepository,
	jwt utils.JWTService,
) AuthService {
	return &authService{
		userRepo: userRepo,
		jwt:      jwt,
	}
}

func (s *authService) Register(
	ctx context.Context,
	req request.RegisterUserRequest,
) (response.AuthResponse, error) {

	// 1. Validasi
	if req.Email == "" || req.Password == "" || req.Name == "" {
		return response.AuthResponse{}, errors.New("field tidak boleh kosong")
	}

	// 2. Cek email sudah ada
	_, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err == nil {
		return response.AuthResponse{}, errors.New("email sudah terdaftar")
	}

	// 3. Hash password
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return response.AuthResponse{}, err
	}

	// 4. Simpan user
	user := entity.Users{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashed),
	}

	created, err := s.userRepo.Create(ctx, user)
	if err != nil {
		return response.AuthResponse{}, err
	}

	return response.AuthResponse{
		Name:  created.Name,
		Email: created.Email,
	}, nil
}


func (s *authService) Login(
	ctx context.Context,
	req request.LoginUserRequest,
) (response.LoginResponse, error) {

	// 1. Validasi
	if req.Email == "" || req.Password == "" {
		return response.LoginResponse{}, errors.New("email & password wajib diisi")
	}

	// 2. Cari user
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return response.LoginResponse{}, errors.New("email tidak ditemukan")
	}

	// 3. Compare password
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	); err != nil {
		return response.LoginResponse{}, errors.New("password salah")
	}

	// 4. Generate JWT
	token, err := s.jwt.Generate(int(user.ID), user.Email)
	if err != nil {
		return response.LoginResponse{}, err
	}

	return response.LoginResponse{
	
		Token: token,
	}, nil
}
