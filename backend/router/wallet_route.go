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

	wallet.Get("/", middleware.UserMiddleware, handler.WalletHandlerGetByUser)
}