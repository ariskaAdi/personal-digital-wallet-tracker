package middleware

import "github.com/gofiber/fiber/v2"

func UserMiddleware(c *fiber.Ctx) error {
	token := c.Get("x-token")
	if token != "secret" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"messsage": "unauthorized",
		})
	}

	return c.Next()
}
