package repositories

import (
	"fiber/database"
	"fiber/model/entity"
	"fiber/model/response"

	"github.com/gofiber/fiber/v2"
)

func CreateTransaction(c *fiber.Ctx, txType string) error {
	user := c.Locals("userId").(entity.User)

	// parse body
	var body struct {
		Amount int `json:"amount"`
		Notes string `json:"notes"`
		WalletID uint `json:"wallet_id"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "invalid body request",
		})
	}

	// connect database
	var wallet entity.Wallet
	if err := database.DB.Where("id = ? AND user_id = ?", body.WalletID, user.ID).First(&wallet).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message" : "wallet not found",
		})
	}

	// make transaction 
	tx := entity.Transaction{
		WalletID: wallet.ID,
		UserID: user.ID,
		Amount: float64(body.Amount),
		Type: txType,
		Notes: body.Notes,
	}

	// check database
	if err := database.DB.Create(&tx).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "failed to create transaction",
		})
	}

	// update balance
	if txType == "INCOME" {
		wallet.Balance += float64(body.Amount)
	} else {
		wallet.Balance -= float64(body.Amount)
	}

	database.DB.Save(&wallet)

	// response DTO

	responseData := response.TxResponse{
		ID: tx.ID,
		Amount: tx.Amount,
		Notes: tx.Notes,
		Type: tx.Type,
		WalletID: tx.WalletID,
		CreatedAt: tx.CreatedAt,
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data": responseData,
	})
}