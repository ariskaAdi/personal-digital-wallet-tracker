package services

import (
	"ariskaAdi/personal-digital-wallet/internal/dto/request"
	"ariskaAdi/personal-digital-wallet/internal/dto/response"
	"ariskaAdi/personal-digital-wallet/internal/model/entity"
	"ariskaAdi/personal-digital-wallet/internal/repositories"
	"ariskaAdi/personal-digital-wallet/internal/utils"
	"context"
	"database/sql"
	"errors"
)

type walletService struct {
	repo repositories.WalletRepository
}

func NewWalletService(repo repositories.WalletRepository) WalletService {
	return &walletService{repo}
}


func (s *walletService) Create(context context.Context,req request.CreateWalletRequest, userID int) (response.WalletResponse, error) {
	// VALIDASI REQUEST
	if req.Name == "" || req.Balance <= 0 {
		return response.WalletResponse{}, errors.New("field tidak boleh kosong")
	}

	wallet := entity.Wallets{
		UserID:    uint(userID),
		Name:      req.Name,
		Balance:   float64(req.Balance),
	}
	
	// SAVE WALLET
	saved, err := s.repo.Create(context, wallet)
	if err != nil {
		if utils.IsUniqueViolation(err) {
			return response.WalletResponse{}, errors.New("wallet name already exist")
		}
		return response.WalletResponse{}, err
	}

	return  response.NewWalletResponse(saved), nil

}
func (s *walletService) Update(context context.Context, req request.UpdateWalletRequest, userID int,) (response.WalletResponse, error) {
	// VALIDASI REQUEST
	if req.Name == "" || req.Balance <= 0 {
		return response.WalletResponse{}, errors.New("field tidak boleh kosong")
	}

	wallet := entity.Wallets{
		UserID:    uint(userID),
		Name:      req.Name,
		Balance:   float64(req.Balance),
	}
	
	// UPDATE WALLET
	updated, err := s.repo.Update(context, wallet)
	if err != nil {
		if utils.IsUniqueViolation(err) {
			return response.WalletResponse{}, errors.New("wallet name already exist")
		}
		return response.WalletResponse{}, err
	}

	return  response.UpdateWalletResponse(updated), nil
}
func (s *walletService) Delete(context context.Context, id int, userID int) (error) {
		if id <= 0 {
		return errors.New("id tidak valid")
	}
	
	err := s.repo.Delete(context, id, userID)
	if err == sql.ErrNoRows {
		return errors.New("wallet tidak ditemukan")
	}

	return err
}
func (s *walletService) FindAll(context context.Context, userID int) ([]entity.Wallets, error) {
	return  s.repo.FindAll(context, userID)
}

func (s *walletService) FindById(context context.Context, id int, userID int,) (entity.Wallets, error) {
	if id <= 0 {
		return entity.Wallets{}, errors.New("id tidak valid")
	}
	return  s.repo.FindById(context, id, userID)
	
}