package routes

import (
	"ariskaAdi/personal-digital-wallet/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, handler *handler.AuthHandler) {
    auth := app.Group("/auth")

    auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)
}
