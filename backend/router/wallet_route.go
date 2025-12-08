package router

import (
	"fiber/handler"
	"fiber/middleware"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func WalletRoute(r *fiber.App) {
	 fmt.Println("WALLET ROUTE REGISTERED")
	wallet := r.Group("/wallet")

	
	wallet.Get("/all", middleware.UserMiddleware, handler.WalletHandlerGetByUser)
	wallet.Get("/total", middleware.UserMiddleware, handler.WalletHandlerGetAllBalance)
	wallet.Get("/:id", middleware.UserMiddleware, handler.WalletHandlerGetByIdWallet)
	wallet.Post("/create", middleware.UserMiddleware, handler.WalletHandlerCreate)
}