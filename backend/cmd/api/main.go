package main

import (
	"ariskaAdi/personal-digital-wallet/internal/config"
	"ariskaAdi/personal-digital-wallet/internal/handler"
	"ariskaAdi/personal-digital-wallet/internal/repositories"
	"ariskaAdi/personal-digital-wallet/internal/routes"
	"ariskaAdi/personal-digital-wallet/internal/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// INITIAL DATABASE
	db := config.OpenConnection()

		 // USERS
    userRepo := repositories.NewUserRepository(db)
    userService := services.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userService)

	// Routes
    routes.UserRoutes(app, userHandler)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":4000")
}