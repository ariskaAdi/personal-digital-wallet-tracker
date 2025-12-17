package entity

import "time"

type Transactions struct {
	ID       uint      `json:"id" db:"id"`
	Amount   float64   `json:"amount" db:"amount"`
	Notes    string    `json:"notes" db:"notes"`
	Type     string    `json:"type" db:"type"`

	WalletID uint      `json:"wallet_id" db:"wallet_id"`
	UserID   uint      `json:"user_id" db:"user_id"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
