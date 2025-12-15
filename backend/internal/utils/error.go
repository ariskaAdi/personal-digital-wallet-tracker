package utils

import "github.com/gofiber/fiber/v2"

func ErrorMessage(c *fiber.Ctx, status int, err error ) error {
	if err == nil {
		return nil
	}

	return c.Status(status).JSON(fiber.Map{
		"success": false,
		"message": err.Error(),
	})
}