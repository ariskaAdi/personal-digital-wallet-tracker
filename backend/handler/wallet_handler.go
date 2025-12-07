package handler

import (
	"fiber/database"
	"fiber/model/entity"
	"fiber/model/response"

	"github.com/gofiber/fiber/v2"
)

func WalletHandlerGetByUser(c *fiber.Ctx) error {
    userRaw := c.Locals("userId").(entity.User)
    userId := userRaw.ID

    var wallets []entity.Wallet
    if err := database.DB.Where("user_id = ?", userId).Find(&wallets).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{
            "success": false,
            "message": "failed to get wallet",
        })
    }

    var responseData []response.WalletResponse
    for _, w := range wallets {
        responseData = append(responseData, response.WalletResponse{
            ID:        w.ID,
            UserID:    w.UserId, 
            Name:      w.Name,
            Balance:   w.Balance,
            CreatedAt: w.CreatedAt,
            UpdatedAt: w.UpdatedAt,
        })
    }

    return c.JSON(fiber.Map{
        "success": true,
        "data":    responseData,
    })
}


func WalletHandlerCreate(c *fiber.Ctx) error {
	  // Ambil userId dari middleware
    userId := c.Locals("userId").(entity.User)
    
	var body struct {
		Name string `json:"name"`
	}
	
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "bad request",
		})
	}

	wallet := entity.Wallet{
		UserId: userId.ID,
		Name: body.Name,
		Balance: 0,
	}

	if err := database.DB.Create(&wallet).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "failed to create wallet",
		})
	}

	return  c.JSON(fiber.Map{
		"success": true,
		"data": wallet,
	})
}
