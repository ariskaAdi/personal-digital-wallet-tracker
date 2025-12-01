package middleware

import (
	"fiber/utils"

	"github.com/gofiber/fiber/v2"
)

func UserMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success" : false,
			"messsage": "unauthenticated",
		})
	}

	name, email, err := utils.GetUserInfoFromToken(token)
	if err != nil {
		return  c.Status(401).JSON(fiber.Map{
			"success" : false,
			"message" : "invalid token",
		})
	}

	c.Locals("name", name)
	c.Locals("email", email)
	return c.Next()
}
