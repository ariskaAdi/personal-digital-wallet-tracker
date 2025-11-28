package router

import (
	"fiber/handler"
	"fiber/middleware"

	"github.com/gofiber/fiber/v2"
)



func UserRoute(r *fiber.App) {
	r.Get("/user",middleware.UserMiddleware, handler.UserHandlerGetAll)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Patch("/user/:id", handler.UserHandlerUpdateUser)
	r.Delete("/user/:id", handler.UserHandlerDeleteUser)
}