package user_controller

import (
	"github.com/LUISEDOCCOR/api-mvc/models"
	user_model "github.com/LUISEDOCCOR/api-mvc/models/user"
	"github.com/LUISEDOCCOR/api-mvc/utils"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return utils.CreateResponse(c, 400, "Invalid JSON")
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		return utils.CreateResponse(c, 400, "All data is required")
	}

	existingUser, err := user_model.GetByEmail(user.Email)

	if err != nil {
		return utils.CreateResponse(c, 500, "Database error")
	}

	if existingUser.ID != 0 {
		return utils.CreateResponse(c, 400, "The email already exists")
	}

	// Modify the user (password) and encrypt it
	utils.HashPassword(user)

	user_model.Create(user)

	token, _ := utils.CreateToken(user)

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func Login(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return utils.CreateResponse(c, 400, "Invalid JSON")
	}

	if user.Email == "" || user.Password == "" {
		return utils.CreateResponse(c, 400, "All data is required")
	}

	existingUser, err := user_model.GetByEmail(user.Email)

	if err != nil {
		return utils.CreateResponse(c, 500, "Database error")
	}

	if existingUser.ID == 0 {
		return utils.CreateResponse(c, 400, "Incorrect or non-existent data")
	}

	mach, _ := utils.VerifyPassword(user.Password, existingUser.Password)

	if !mach {
		return utils.CreateResponse(c, 400, "Incorrect or non-existent data")
	}

	token, _ := utils.CreateToken(&existingUser)

	return c.JSON(fiber.Map{
		"token": token,
	})
}
