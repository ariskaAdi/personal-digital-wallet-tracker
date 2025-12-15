package handler

import (
	"ariskaAdi/personal-digital-wallet/internal/dto/request"
	"ariskaAdi/personal-digital-wallet/internal/services"
	"ariskaAdi/personal-digital-wallet/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{service}
}


func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req request.RegisterUserRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorMessage(c, 400, err)
	}

	res, err := h.service.Register(c.Context(), req)
	if err != nil {
		return utils.ErrorMessage(c, 400, err)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    res,
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req request.LoginUserRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorMessage(c, 400, err)
	}

	res, err := h.service.Login(c.Context(), req)
	if err != nil {
		return utils.ErrorMessage(c, 400, err)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    res,
	})
}
