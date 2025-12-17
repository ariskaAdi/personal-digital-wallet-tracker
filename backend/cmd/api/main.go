package main

import (
	"ariskaAdi/personal-digital-wallet/external/database"
	"ariskaAdi/personal-digital-wallet/internal/config"
	"ariskaAdi/personal-digital-wallet/internal/handler"
	"ariskaAdi/personal-digital-wallet/internal/repositories"
	"ariskaAdi/personal-digital-wallet/internal/routes"
	"ariskaAdi/personal-digital-wallet/internal/services"
	"ariskaAdi/personal-digital-wallet/internal/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// INITIAL DATABASE
	filename := "../../cmd/api/config.yaml"
	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}
	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected")
	}

	// USERS
    userRepo := repositories.NewUserRepository(db)
    userService := services.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userService)

	// AUTH
	jwtService := utils.NewJWTService(config.Cfg.App.JWTSecret)
	authService := services.NewAuthService(userRepo, jwtService)
	authHandler := handler.NewAuthHandler(authService)

	// WALLET
	walletRepo := repositories.NewWalletRepository(db)
	walletService := services.NewWalletService(walletRepo)
	walletHandler := handler.NewWalletHandler(walletService)


	// Routes
    routes.UserRoutes(app, userHandler)
	routes.AuthRoutes(app, authHandler)
	routes.WalletRoutes(app, walletHandler, jwtService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Personal Digital Wallet API")
	})

	

	app.Listen(":4000")
}