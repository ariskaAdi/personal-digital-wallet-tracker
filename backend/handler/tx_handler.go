package handler

import (
	"fiber/database"
	"fiber/model/entity"
	"fiber/model/response"
	"fiber/repositories"

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
	return repositories.CreateTransaction(c, "INCOME")
}

func TxExpenseHandler(c *fiber.Ctx) error {
	return repositories.CreateTransaction(c, "EXPENSE")
}