package repositories

import "github.com/jmoiron/sqlx"

type TransactionRepository interface{
	Create()
	Update()
	Delete()
	FindAll()
	FindAllIncome()
	FindAllExpense()
}

type transactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{}
}