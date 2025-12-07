package middleware

import (
	"fiber/database"
	"fiber/model/entity"
	"fiber/utils"

	"github.com/gofiber/fiber/v2"
)

func UserMiddleware(c *fiber.Ctx) error {
	// token := c.Get("Authorization")
	token := c.Cookies("token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success" : false,
			"messsage": "unauthenticated",
		})
	}

	id, err := utils.GetUserInfoFromToken(token)
	if err != nil {
		return  c.Status(401).JSON(fiber.Map{
			"success" : false,
			"message" : "invalid token",
		})
	}

	var user entity.User
	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success" : false,
			"message": "invalid token",
		})
	}

	c.Locals("userId", user)
	return c.Next()
}
