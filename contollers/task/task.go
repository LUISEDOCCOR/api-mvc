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

func GetById(c *fiber.Ctx) error {
	id, err := utils.ParseUint(c.Params("id"))
	tokenData := utils.GetClaims(c)

	if err != nil {
		return utils.CreateResponse(c, 400, "Invalid Task ID")
	}

	task, err := task_model.GetById(id, tokenData.Id)

	if task.ID == 0 {
		return utils.CreateResponse(c, 404, "No task was found with that id")
	}

	return c.JSON(task)

}

func Delete(c *fiber.Ctx) error {
	id, err := utils.ParseUint(c.Params("id"))
	tokenData := utils.GetClaims(c)

	if err != nil {
		return utils.CreateResponse(c, 400, "Invalid Task ID")
	}

	task, err := task_model.GetById(id, tokenData.Id)

	if task.ID == 0 {
		return utils.CreateResponse(c, 404, "No task was found with that id")
	}

	err = task_model.Delete(task.ID)

	if err != nil {
		return utils.CreateResponse(c, 500, "Error deleting task")
	}

	return c.SendStatus(200)
}
