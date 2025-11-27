package router

import (
	"fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(r *fiber.App) {
	r.Get("/user", handler.UserHandlerGetAll)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Post("/user", handler.UserHandlerAddUser)
	r.Patch("/user/:id", handler.UserHandlerUpdateUser)
	r.Delete("/user/:id", handler.UserHandlerDeleteUser)
}