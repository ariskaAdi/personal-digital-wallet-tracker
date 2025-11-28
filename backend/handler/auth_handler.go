package handler

import (
	"fiber/database"
	"fiber/model/entity"
	"fiber/model/request"
	"fiber/model/response"
	"fiber/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	
	// VALIDATE TO CHECK REQUIRED FIELD 
	validate := validator.New()
	errRequest := validate.Struct(loginRequest)
	if errRequest != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": errRequest.Error(),
		})
	}

	// CHECK AVAILABLE USER BY EMAIL
	var user entity.User
	err := database.DB.Debug().First(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "user not found",
		})
	}

	// VALIDATE HASH PASSWORD
	isValid := utils.CheckedHashPassword(loginRequest.Password, user.Password)
	if !isValid {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "invalid password or email",
		})
	}

	return c.JSON(fiber.Map{
		"token" : "jsbdjs<JWT_TOKEN>",
	})
}

func RegisterHandler(c *fiber.Ctx) error {
	user := new(request.RegisterRequest)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	// VALIDATE TO CHECK REQUIRED FIELD 
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	
	// CREATE NEW USER
	newUser := entity.User{
		Name:     user.Name,
		Email:    user.Email,
	}

	// HASH PASSWORD
	hashedPassword, err := utils.HashedPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	newUser.Password = hashedPassword

	// RETURN USER RESPONSE
	userResponse := response.UserResponse{
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	// CREATE USER TO DATABASE
	database.DB.Debug().Create(&newUser)
	return c.JSON(fiber.Map{
		"message": "success add user",
		"data":    userResponse,
	})
}