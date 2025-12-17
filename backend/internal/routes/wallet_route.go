package routes

import (
	"ariskaAdi/personal-digital-wallet/internal/handler"
	"ariskaAdi/personal-digital-wallet/internal/middleware"
	"ariskaAdi/personal-digital-wallet/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func WalletRoutes(app *fiber.App, handler *handler.WalletHandler, jwt utils.JWTService) {
	wallet := app.Group("/wallets", middleware.AuthMiddleware(jwt))

	wallet.Post("/", handler.Create)
	wallet.Patch("/:id", handler.Update)
	wallet.Delete("/:id", handler.Delete)
	wallet.Get("/all", handler.FindAll)
	wallet.Get("/:id", handler.FindById)
}