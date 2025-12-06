package main

import (
	"fiber/database"
	"fiber/database/migration"
	"fiber/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	// INITIAL DATABASE
	database.DatabaseInit()
	migration.RunMigration()

	// INITIAL APP
	app := fiber.New()

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))


	// INITIAL ROUTE
	router.UserRoute(app)
	router.AuthRoute(app)
	router.WalletRoute(app)
	router.TxRoute(app)
	

	app.Listen(":4000")
}