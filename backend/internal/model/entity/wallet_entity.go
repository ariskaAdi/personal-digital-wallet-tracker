package entity

import "time"

type Wallets struct {
	ID     uint   `json:"id" db:"id"`
	UserID uint   `json:"user_id" db:"user_id"`

	Name    string  `json:"name" db:"name"`
	Balance float64 `json:"balance" db:"balance"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
