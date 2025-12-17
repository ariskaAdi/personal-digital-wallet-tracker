package response

import (
	"ariskaAdi/personal-digital-wallet/internal/model/entity"
	"time"
)

type WalletResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewWalletResponse(wallet entity.Wallets) WalletResponse {
	return WalletResponse{
		Id:        int(wallet.ID),
		Name:      wallet.Name,
		Balance:   int(wallet.Balance),
		CreatedAt: wallet.CreatedAt,
		UpdatedAt: wallet.UpdatedAt,
	}
}

func UpdateWalletResponse(wallet entity.Wallets) WalletResponse {
	return WalletResponse{
		Name:      wallet.Name,
		Balance:   int(wallet.Balance),
		UpdatedAt: wallet.UpdatedAt,
	}
}

func GetAllWalletResponse (wallet []entity.Wallets) ([]WalletResponse) {
	res := make([]WalletResponse,0, len(wallet) )
	for _, w := range wallet {
		res = append(res, NewWalletResponse(w))
	}
	return res
}