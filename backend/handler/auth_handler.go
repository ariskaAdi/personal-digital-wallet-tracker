package handler

import (
	"fiber/database"
	"fiber/model/entity"
	"fiber/model/request"
	"fiber/model/response"
	"fiber/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

var validate = validator.New()

func LoginHandler(c *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	
	// VALIDATE TO CHECK REQUIRED FIELD 
	errRequest := validate.Struct(loginRequest)
	if errRequest != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": errRequest.Error(),
		})
	}

	// CHECK AVAILABLE USER BY EMAIL
	var user entity.User
	err := database.DB.Take(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "user not found",
		})
	}

	// VALIDATE HASH PASSWORD
	if  !utils.CheckedHashPassword(loginRequest.Password, user.Password) {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "invalid password or email",
		})
	}

	// GENERATE JWT TOKEN
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	
	token, errGenToken := utils.GenerateToken(&claims)
	if errGenToken != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": errGenToken.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success" : true,
		"message" : "success login",
		"data": fiber.Map{
			"token": token,
		},
	})
}

func RegisterHandler(c *fiber.Ctx) error {
	user := new(request.RegisterRequest)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	//VALIATE EXISTING USER BY EMAIL
	var existingUser entity.User
	errExistingUser := database.DB.Take(&existingUser, "email = ?", user.Email).Error
	if errExistingUser == nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "user already exist",
		})
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
		ID:    newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	// CREATE USER TO DATABASE
	database.DB.Create(&newUser)
	return c.JSON(fiber.Map{
		"message": "success add user",
		"data":    userResponse,
	})
}