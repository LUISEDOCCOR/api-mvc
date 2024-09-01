package task_controller

import (
	"github.com/LUISEDOCCOR/api-mvc/models"
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

func Create(c *fiber.Ctx) error {
	tokenData := utils.GetClaims(c)
	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		return utils.CreateResponse(c, 200, "Invalid JSON")
	}

	if task.Title == "" {
		return utils.CreateResponse(c, 200, "The title is a required field")
	}

	task.UserId = tokenData.Id

	err := task_model.Create(task)

	if err != nil {
		return utils.CreateResponse(c, 200, "Database Error")
	}

	return c.JSON(task)
}
