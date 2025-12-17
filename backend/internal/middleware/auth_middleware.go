package middleware

import (
	"ariskaAdi/personal-digital-wallet/internal/utils"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(jwtService utils.JWTService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.ErrorMessage(c, fiber.StatusUnauthorized, errors.New("unauthorized"))
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return utils.ErrorMessage(c, fiber.StatusUnauthorized, errors.New("invalid token format"))
		}

		claims, err := jwtService.Decode(parts[1])
		if err != nil {
			return utils.ErrorMessage(c, fiber.StatusUnauthorized, err)
		}

		userID, ok := claims["user_id"].(float64)
		if !ok {
			return utils.ErrorMessage(c, fiber.StatusUnauthorized, errors.New("invalid token payload"))
		}

		c.Locals("userId", int(userID))
		return c.Next()
	}
}
