package entity

import "time"

type Category struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" `
	Type string `json:"type" `

	UserID uint `json:"user_id" `
	User   User `json:"user" gorm:"foreignKey:UserID"`

	Transaction []Transaction `json:"transaction"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	
}