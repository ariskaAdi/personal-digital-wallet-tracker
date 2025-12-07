package router

import (
	"fiber/handler"
	"fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func TxRoute(r *fiber.App) {
	tx := r.Group("/tx")

	tx.Get("/", middleware.UserMiddleware, handler.TxHistoryAllHandler)
	tx.Post("/income", middleware.UserMiddleware, handler.TxIncomeHandler)
	tx.Post("/expense", middleware.UserMiddleware, handler.TxExpenseHandler)

}