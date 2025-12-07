package handler

import (
	"fiber/database"
	"fiber/model/entity"
	"fiber/model/response"

	"github.com/gofiber/fiber/v2"
)


func TxHistoryAllHandler(c *fiber.Ctx) error {
	user := c.Locals("userId").(entity.User)

	var txs []entity.Transaction
	if err := database.DB.Where("user_id = ?", user.ID).Order("created_at DESC").Find(&txs).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "failed to get transaction history",
		})
	}

	// RESPONSE DATA 
	var responseData []response.TxResponse
	for _, tx := range txs {
		responseData = append(responseData, response.TxResponse{
			ID: tx.ID,
			Amount: tx.Amount,
			Notes: tx.Notes,
			Type: tx.Type,
			WalletID: tx.WalletID,
			CreatedAt: tx.CreatedAt,
		})
	}

	return c.JSON(fiber.Map{
		"success" : true,
		"data": responseData,
	})
}

func TxIncomeHandler(c *fiber.Ctx) error {
	user := c.Locals("userId").(entity.User)

	// parse body
	var body struct {
		Amount int `json:"amount"`
		Notes string `json:"notes"`
		WalletID uint `json:"wallet_id"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success" : false,
			"message" : "invalid body",
		})
	}

	// get wallet user
	var wallet entity.Wallet
	if err := database.DB.Where("id = ? AND user_id = ?" , body.WalletID, user.ID).First(&wallet).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "failed to get wallet",
		})
	}

	// create new transaction
	tx := entity.Transaction{
		WalletID: wallet.ID,
		UserID:   user.ID,
		Amount : float64(body.Amount),
		Type: "INCOME",
		Notes: body.Notes,
	}

	if err := database.DB.Create(&tx).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "failed to create transaction",
		})
	}

	// update wallet
	wallet.Balance += float64(body.Amount)
	database.DB.Save(&wallet)

	return c.JSON(fiber.Map{
		"success": true,
		"data" : fiber.Map{
			"transaction" : tx,
			"wallet" : wallet,
		},
	})
}

func TxExpenseHandler(c *fiber.Ctx) error {
	user := c.Locals("userId").(entity.User)

	// parse body
	var body struct {
		Amount int `json:"amount"`
		Notes string `json:"notes"`
		WalletID uint `json:"wallet_id"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success" : false,
			"message" : "invalid body",
		})
	}

	// get wallet user
	var wallet entity.Wallet
	if err := database.DB.Where("id = ? AND user_id = ?" , body.WalletID, user.ID).First(&wallet).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "failed to get wallet",
		})
	}

	// create new transaction
	tx := entity.Transaction{
		WalletID: wallet.ID,
		UserID:   user.ID,
		Amount : float64(body.Amount),
		Type: "EXPENSE",
		Notes: body.Notes,
	}

	if err := database.DB.Create(&tx).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "failed to create transaction",
		})
	}

	// update wallet
	wallet.Balance -= float64(body.Amount)
	database.DB.Save(&wallet)
	

	// RESPONSE DATA 
	var responseData []response.TxResponse
	for _, tx := range wallet.Transactions {
		responseData = append(responseData, response.TxResponse{
			ID: tx.ID,
			Amount: tx.Amount,
			Notes: tx.Notes,
			Type: tx.Type,
			WalletID: tx.WalletID,
			CreatedAt: tx.CreatedAt,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data" : fiber.Map{
			"transaction" : tx,
			"wallet" : responseData,
		},
	})
}