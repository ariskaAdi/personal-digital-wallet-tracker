package response

import "time"

type TxResponse struct {
	ID        uint      `json:"id"`
	Amount    float64   `json:"amount"`
	Notes     string    `json:"notes"`
	Type      string    `json:"type"`
	WalletID  uint      `json:"wallet_id"`
	CreatedAt time.Time `json:"created_at"`
}