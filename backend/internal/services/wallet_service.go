package services

import (
	"ariskaAdi/personal-digital-wallet/internal/dto/request"
	"ariskaAdi/personal-digital-wallet/internal/dto/response"
	"ariskaAdi/personal-digital-wallet/internal/model/entity"
	"context"
)

type WalletService interface {
	Create(context context.Context,req request.CreateWalletRequest,userID int) (response.WalletResponse, error)
	FindAll(context context.Context, userID int) ([]entity.Wallets, error)
	FindById(context context.Context, id int, userID int,) (entity.Wallets, error)
	Update(context context.Context, req request.UpdateWalletRequest, userID int,) (response.WalletResponse, error)
	Delete(context context.Context, id int, userID int) (error)
}