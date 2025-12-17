package repositories

import (
	"ariskaAdi/personal-digital-wallet/internal/model/entity"
	"context"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface{
	Create(context context.Context, transaction entity.Transactions) (entity.Transactions, error)
	Update(context context.Context, transation entity.Transactions) (entity.Transactions, error)
	Delete(context context.Context, id int) error
	FindAll(context context.Context) ([]entity.Transactions, error)
	FindAllIncome( context.Context, )
	FindAllExpense()
	FindByDate()
}

type transactionRepository struct {
	db *sqlx.DB
}

// func NewTransactionRepository() TransactionRepository {
// 	return &transactionRepository{}
// }