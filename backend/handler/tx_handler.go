package handler

import (
	"fiber/database"
	"fiber/model/entity"

	"github.com/gofiber/fiber/v2"
)

func TxIncomeHandler(c *fiber.Ctx) error {
	rawId := c.Locals("userId")
	if rawId == nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "unauthenticated",
		})
	}
	uid, ok := rawId.(uint)
	if !ok {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "unauthenticated",
		})
	}

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
	if err := database.DB.Where("id = ? AND user_id = ?" , body.WalletID, uid).First(&wallet).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "failed to get wallet",
		})
	}

	// create new transaction
	tx := entity.Transaction{
		WalletID: wallet.ID,
		UserID:   uid,
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

	return nil
}