package router

import (
	"fiber/handler"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(r *fiber.App) {
	fmt.Println("tx route run")
	r.Post("/auth/login", handler.LoginHandler)
	r.Post("/auth/register", handler.RegisterHandler)
}