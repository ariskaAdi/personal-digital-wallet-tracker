package router

import (
	"fiber/handler"
	"fiber/middleware"

	"github.com/gofiber/fiber/v2"
)



func UserRoute(r *fiber.App) {
	user := r.Group("/user")

	user.Get("/", handler.UserHandlerGetAll)
	user.Get("/me",middleware.UserMiddleware, handler.UserHandlerGetMe)
	user.Get("/:id", handler.UserHandlerGetById)
	user.Patch("/:id", handler.UserHandlerUpdateUser)
	user.Delete("/:id", handler.UserHandlerDeleteUser)
}