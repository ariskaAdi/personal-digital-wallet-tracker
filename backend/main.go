package main

import (
	"fiber/database"
	"fiber/database/migration"
	"fiber/router"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// INITIAL DATABASE
	database.DatabaseInit()
	migration.RunMigration()

	// INITIAL APP
	app := fiber.New()

	// INITIAL ROUTE
	router.UserRoute(app)

	app.Listen(":4000")
}