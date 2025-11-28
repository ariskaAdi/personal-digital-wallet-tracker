package router

import (
	"fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(r *fiber.App) {
	r.Post("/auth/login", handler.LoginHandler)
	r.Post("/auth/register", handler.RegisterHandler)
}