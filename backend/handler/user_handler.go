package handler

import (
	"fiber/database"
	"fiber/model/entity"
	"fiber/model/request"
	"fiber/model/response"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/go-playground/validator/v10"
)

func UserHandlerGetAll(c *fiber.Ctx) error {
		var users []entity.User
		err := database.DB.Debug().Find(&users).Error
		if err != nil {
			log.Println(err)
		}
		// RETURN USER RESPONSE
	var userResponses []response.UserResponse
		for _, user := range users {
			userResponse := response.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			}
			userResponses = append(userResponses, userResponse)
		}
		return c.JSON(fiber.Map{
			"message": "data user",
			"data": userResponses,
		})
	
}

func UserHandlerGetById(c *fiber.Ctx) error {
	userId := c.Params("id")

	// GET USER BY ID FROM DATABASE
	var user entity.User
	err := database.DB.Debug().First(&user, "id = ?", userId).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// RETURN USER RESPONSE
	userResponse := response.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,

	}
	return c.JSON(fiber.Map{
		"message": "data user",
		"data":    userResponse,
	})
}

func UserHandlerAddUser(c *fiber.Ctx) error {
	user := new(request.UserCreateRequest)
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
		Password: user.Password,
	}

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

func UserHandlerUpdateUser(c *fiber.Ctx) error {

	// VALIDATE TO CHECK REQUIRED FIELD
	userUpdate := new(request.UserUpdateRequest)
	if err := c.BodyParser(userUpdate); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "bad request",
		})
	}


	userId := c.Params("id")
	// GET USER BY ID FROM DATABASE
	var user entity.User
	err := database.DB.Debug().First(&user, "id = ?", userId).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "user not found",
		})
	}


	// UPDATE USER
	if userUpdate.Name != "" {
		user.Name = userUpdate.Name
	}

	if userUpdate.Email != "" {
		user.Email = userUpdate.Email
	}
	errUpdate := database.DB.Debug().Save(&user).Error
	if errUpdate != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": errUpdate.Error(),
		})
	}

		// RETURN USER RESPONSE
	userResponse := response.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,

	}

	// UPDATE USER
	return c.JSON(fiber.Map{
		"message": "data user",
		"data":    userResponse,
	})
}

func UserHandlerDeleteUser(c *fiber.Ctx) error {
		userId := c.Params("id")
	// GET USER BY ID FROM DATABASE
	var user entity.User
	err := database.DB.Debug().First(&user, "id = ?", userId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "user not found",
		})
	}

	// DELETE USER
	errDelete := database.DB.Debug().Delete(&user).Error
	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "error delete user",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success delete user",
	})
	
}