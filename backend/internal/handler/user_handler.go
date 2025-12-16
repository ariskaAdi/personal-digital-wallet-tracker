package handler

import (
	"ariskaAdi/personal-digital-wallet/internal/dto/request"
	"ariskaAdi/personal-digital-wallet/internal/dto/response"
	"ariskaAdi/personal-digital-wallet/internal/services"
	"ariskaAdi/personal-digital-wallet/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service}
}


func (h *UserHandler) Update(c *fiber.Ctx) error {
	  var req request.UpdateUserRequest
    if err := c.BodyParser(&req); err != nil {
       return utils.ErrorMessage(c, fiber.StatusBadRequest, err)
    }

    res, err := h.service.Update(c.Context(), req)
    if err != nil {
        return utils.ErrorMessage(c, fiber.StatusBadRequest, err)
    }

    response := response.NewUserResponse(res)

    return c.JSON(fiber.Map{
        "success": true,
        "data": response,
    })
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	  id, err := c.ParamsInt("id")
    if err != nil {
       return utils.ErrorMessage(c, fiber.StatusBadRequest, err)
    }

    err = h.service.Delete(c.Context(), id)
    if err != nil {
        return utils.ErrorMessage(c, fiber.StatusBadRequest, err)
    }

    return c.JSON(fiber.Map{"message": "user deleted"})
}

func (h *UserHandler) FindAll(c *fiber.Ctx) error {
	users, err := h.service.FindAll(c.Context())
    if err != nil {
        return utils.ErrorMessage(c, fiber.StatusBadRequest, err)
    }

	response := response.NewUserResponses(users)

    return c.JSON(fiber.Map{
		"success": true,
		"data": response,
	})
}

func (h *UserHandler) FindById(c *fiber.Ctx) error {
	 id, err := c.ParamsInt("id")
    if err != nil {
        return utils.ErrorMessage(c, fiber.StatusBadRequest, err)
    }

    user, err := h.service.FindById(c.Context(), id)
    if err != nil {
        return utils.ErrorMessage(c, fiber.StatusBadRequest, err)
    }

    response := response.NewUserResponse(user)

    return c.JSON(fiber.Map{
        "success": true,
        "data" : response,
    })
}