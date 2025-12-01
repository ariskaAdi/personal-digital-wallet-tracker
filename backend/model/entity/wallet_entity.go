package entity

import "time"

type Wallet struct {
	ID     uint `json:"id" gorm:"primary_key"`
	UserId uint `json:"user_id"`
	User   User `json:"user" gorm:"foreignKey:UserId"`

	Name         string        `json:"name"`
	Balance      float64       `json:"balance"`
	Transactions []Transaction `json:"transactions" gorm:"constraint:OnDelete:CASCADE"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}