package entity

import "time"

type Transaction struct {
	ID     uint      `json:"id" gorm:"primaryKey"`
	Amount float64   `json:"amount"`
	Notes  string    `json:"notes"`
	Type string `json:"type"`

	WalletID uint   `json:"wallet_id"`
	Wallet   Wallet `json:"wallet" gorm:"foreignKey:WalletID"`

	UserID uint `json:"user_id"`
	User   Users `json:"user"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}