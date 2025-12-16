package repositories

import (
	"ariskaAdi/personal-digital-wallet/internal/model/entity"
	"context"

	"github.com/jmoiron/sqlx"
)

type WalletRepository interface {
	Create(context context.Context, wallet entity.Wallet) (entity.Wallet, error)
	Update(context context.Context, wallet entity.Wallet) (entity.Wallet, error)
	Delete(context context.Context, id int) error
	FindAll(context context.Context) ([]entity.Wallet, error)
	FindById(context context.Context, id int) (entity.Wallet, error)
}

type walletRepository struct {
	db *sqlx.DB
}

func NewWalletRepository(db *sqlx.DB) WalletRepository {
	return &walletRepository{db: db}
}

func(r *walletRepository) Create(context context.Context, wallet entity.Wallet) (entity.Wallet, error) {
	return entity.Wallet{}, nil	
}

func (r *walletRepository) Update(context context.Context, wallet entity.Wallet) (entity.Wallet, error) {
	return entity.Wallet{}, nil	
}

func (r *walletRepository) Delete(context context.Context, id int) error {
	return nil	
}

func (r *walletRepository) FindAll(context context.Context) ([]entity.Wallet, error) {
	return []entity.Wallet{}, nil	
}

func (r *walletRepository) FindById(context context.Context, id int) (entity.Wallet, error) {
	return entity.Wallet{}, nil
}

