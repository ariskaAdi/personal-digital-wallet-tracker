package handler

import (
	"ariskaAdi/personal-digital-wallet/internal/dto/request"
	"ariskaAdi/personal-digital-wallet/internal/dto/response"
	"ariskaAdi/personal-digital-wallet/internal/services"
	"ariskaAdi/personal-digital-wallet/internal/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type WalletHandler struct {
	service services.WalletService
}

func NewWalletHandler(service services.WalletService) *WalletHandler {
	return &WalletHandler{service}
}

func (h *WalletHandler) Create(c *fiber.Ctx) error {
	userID, ok := c.Locals("userId").(int)
	if !ok {
		return utils.ErrorMessage(c, fiber.StatusUnauthorized, errors.New("unauthorized"))
	}

	var req request.CreateWalletRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorMessage(c, fiber.StatusBadRequest, err)
	}

	res, err := h.service.Create(c.Context(), req, userID)
	if err != nil {
		return utils.ErrorMessage(c, fiber.StatusBadRequest, err)
	}

	if res.Id == 0 {
		return utils.ErrorMessage(c, fiber.StatusBadRequest, errors.New("wallet name already exist"))
	}

	return  c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success" : true,
		"message": "wallet suscesfully created",
	})
}

func (h *WalletHandler) Update(c *fiber.Ctx) error {
	userID, ok := c.Locals("userId").(int)
	if !ok {
		return utils.ErrorMessage(c, fiber.StatusUnauthorized, errors.New("unauthorized"))
	}

	var req request.UpdateWalletRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorMessage(c, fiber.StatusBadRequest, err)
	}

	res, err := h.service.Update(c.Context(), req, userID)
	if err != nil {
		return utils.ErrorMessage(c, fiber.StatusBadRequest, err)
	}

	if res.Id == 0 {
		return utils.ErrorMessage(c, fiber.StatusBadRequest, errors.New("wallet name already exist"))
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "wallet updated",
	})
}

func (h *WalletHandler) Delete(c *fiber.Ctx) error {
	userID, ok := c.Locals("userId").(int)
	if !ok {
		return utils.ErrorMessage(c, fiber.StatusUnauthorized, errors.New("unauthorized"))
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.ErrorMessage(c, fiber.StatusBadRequest, err)
	}

	err = h.service.Delete(c.Context(), id, userID)
	if err != nil {
		return utils.ErrorMessage(c, fiber.StatusBadRequest, err)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "wallet deleted",
	})
}

func (h *WalletHandler) FindAll(c *fiber.Ctx) error {
	userID, ok := c.Locals("userId").(int)
	if !ok {
		return utils.ErrorMessage(c, fiber.StatusUnauthorized, errors.New("unauthorized"))
	}

	wallets, err := h.service.FindAll(c.Context(), userID)
	if err != nil {
		return utils.ErrorMessage(c, fiber.StatusBadRequest, err)
	}

	response := response.GetAllWalletResponse(wallets)

	return c.JSON(fiber.Map{
		"success": true,
		"data": response,
	})
}

func (h *WalletHandler) FindById(c *fiber.Ctx) error {
	userID, ok := c.Locals("userId").(int)
	if !ok {
		return utils.ErrorMessage(c, fiber.StatusUnauthorized, errors.New("unauthorized"))
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.ErrorMessage(c, fiber.StatusBadRequest, err)
	}

	wallet, err := h.service.FindById(c.Context(), id, userID)
	if err != nil {
		return utils.ErrorMessage(c, fiber.StatusBadRequest, err)
	}

	response := response.NewWalletResponse(wallet)

	return c.JSON(fiber.Map{
		"success": true,
		"data" : response,
	})
}