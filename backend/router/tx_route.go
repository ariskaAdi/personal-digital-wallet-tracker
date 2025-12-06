package router

import (
	"fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func TxRoute(r *fiber.App) {
	tx := r.Group("/tx")

	tx.Post("/income", handler.TxIncomeHandler)
	tx.Post("/expense", handler.TxExpenseHandler)

}