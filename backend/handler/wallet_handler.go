package handler

import (
	"fiber/database"
	"fiber/model/entity"

	"github.com/gofiber/fiber/v2"
)

func WalletHandlerGetByUser(c *fiber.Ctx) error {
	userId := c.Locals("userId")

	var wallet []entity.Wallet
	if err := database.DB.Where("user_id = ?" ,userId).Find(&wallet).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "failed to get wallet",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data": wallet,
	})
}

