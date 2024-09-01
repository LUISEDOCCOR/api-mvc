package task_controller

import (
	task_model "github.com/LUISEDOCCOR/api-mvc/models/task"
	"github.com/LUISEDOCCOR/api-mvc/utils"
	"github.com/gofiber/fiber/v2"
)

func GetALL(c *fiber.Ctx) error {
	tokenData := utils.GetClaims(c)

	tasks, err := task_model.GetAll(tokenData.Id)

	if err != nil {
		return utils.CreateResponse(c, 500, "Database Error")
	}

	return c.JSON(tasks)
}
