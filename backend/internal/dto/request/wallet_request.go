package request

type CreateWalletRequest struct {
	Name    string
	Balance int
}

type UpdateWalletRequest struct {
	Name    string
	Balance int
}