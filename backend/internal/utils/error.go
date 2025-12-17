package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func ErrorMessage(c *fiber.Ctx, status int, err error ) error {
	if err == nil {
		return nil
	}

	return c.Status(status).JSON(fiber.Map{
		"success": false,
		"message": err.Error(),
	})
}

func IsUniqueViolation(err error) bool {
	if pqErr, ok := err.(*pq.Error); ok {
		return pqErr.Code == "23505"
	}
	return false
}