package routes

import (
	"ariskaAdi/personal-digital-wallet/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, handler *handler.UserHandler) {
    user := app.Group("/users")

    user.Get("/", handler.FindAll)
    user.Get("/:id", handler.FindById)
    user.Post("/", handler.Create)
    user.Put("/", handler.Update)
    user.Delete("/:id", handler.Delete)
}
