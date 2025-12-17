package repositories

import (
	"ariskaAdi/personal-digital-wallet/internal/model/entity"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type WalletRepository interface {
	Create(context context.Context, wallet entity.Wallets) (entity.Wallets, error)
	Update(context context.Context, wallet entity.Wallets) (entity.Wallets, error)
	Delete(context context.Context, id int, userID int) error
	FindAll(context context.Context, userID int) ([]entity.Wallets, error)
	FindById(context context.Context, id int, userID int) (entity.Wallets, error)
}

type walletRepository struct {
	db *sqlx.DB
}

func NewWalletRepository(db *sqlx.DB) WalletRepository {
	return &walletRepository{db: db}
}

func(r *walletRepository) Create(context context.Context, wallet entity.Wallets) (entity.Wallets, error) {
	SQL := `
		INSERT INTO wallets (user_id, name, balance, created_at)
		VALUES ($1, $2, $3, now())
		RETURNING id
	`

	err := r.db.QueryRowContext(
		context,
		SQL,
		wallet.UserID,
		wallet.Name,
		wallet.Balance,
	).Scan(&wallet.ID)

	
	if err != nil {
		return entity.Wallets{}, err
	}

	return wallet, err
}

func (r *walletRepository) Update(context context.Context, wallet entity.Wallets) (entity.Wallets, error) {
	SQL := `
		UPDATE wallets
		SET name = $1, balance = $2
		WHERE id = $3 AND user_id = $4
	`
	res, err := r.db.ExecContext(
		context,
		SQL,
		wallet.Name,
		wallet.Balance,
		wallet.ID,
		wallet.UserID,
	)

	if err != nil {
		return entity.Wallets{}, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return entity.Wallets{}, err
	}

	if rows == 0 {
		return entity.Wallets{}, err
	}

	return wallet, nil

}

func (r *walletRepository) Delete(context context.Context, id int, userID int) error {
	SQL := `
		DELETE FROM wallets
		WHERE id = $1 AND user_id = $2
	`

	res, err := r.db.ExecContext(
		context,
		SQL,
		id,
		userID,
	)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return  nil
}

func (r *walletRepository) FindAll(context context.Context, userID int) ([]entity.Wallets, error) {
	var wallets []entity.Wallets
	SQL := `
		SELECT id, user_id, name, balance FROM wallets
		WHERE user_id = $1
	`

	err := r.db.SelectContext(
		context,
		&wallets,
		SQL,
		userID,
	)

	return wallets, err
}

func (r *walletRepository) FindById(context context.Context, id int, userID int) (entity.Wallets, error) {
	var wallet entity.Wallets
	SQL := `
		SELECT id, user_id, name, balance FROM wallets
		WHERE id = $1 AND user_id = $2
	`

	err := r.db.GetContext(
		context,
		&wallet,
		SQL,
		id,
		userID,
	)

	return wallet, err
}

