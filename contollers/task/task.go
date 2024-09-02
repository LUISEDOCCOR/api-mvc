package task_controller

import (
	"errors"

	"github.com/LUISEDOCCOR/api-mvc/models"
	task_model "github.com/LUISEDOCCOR/api-mvc/models/task"
	"github.com/LUISEDOCCOR/api-mvc/utils"
	"github.com/gofiber/fiber/v2"
)

func ExistingTask(c *fiber.Ctx) (models.Task, error) {
	var err error
	var id uint
	var task models.Task

	id, err = utils.ParseUint(c.Params("id"))
	tokenData := utils.GetClaims(c)

	task, err = task_model.GetById(id, tokenData.Id)

	if task.ID == 0 {
		err = errors.New("Not Found")
	}

	return task, err
}

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
	task, err := ExistingTask(c)

	if err != nil {
		return utils.CreateResponse(c, 400, "Invalid Id")
	}

	return c.JSON(task)

}

func Delete(c *fiber.Ctx) error {
	task, err := ExistingTask(c)

	if err != nil {
		return utils.CreateResponse(c, 400, "Invalid Id")
	}

	err = task_model.Delete(task.ID)

	if err != nil {
		return utils.CreateResponse(c, 500, "Error deleting task")
	}

	return c.SendStatus(200)
}

func Finish(c *fiber.Ctx) error {
	task, err := ExistingTask(c)

	if err != nil {
		return utils.CreateResponse(c, 400, "Invalid Id")
	}

	task.IsDone = !task.IsDone

	err = task_model.Update(&task)

	if err != nil {
		return utils.CreateResponse(c, 500, "Error updating task")
	}

	return c.SendStatus(200)
}

func Update(c *fiber.Ctx) error {
	newTask := new(models.Task)

	if err := c.BodyParser(newTask); err != nil {
		return utils.CreateResponse(c, 200, "Invalid JSON")
	}

	task, err := ExistingTask(c)

	if err != nil {
		return utils.CreateResponse(c, 400, "Invalid Id")
	}

	title := func() string {
		if newTask.Title == "" {
			return task.Title
		}
		return newTask.Title
	}

	content := func() string {
		if newTask.Content == "" {
			return task.Content
		}
		return newTask.Content
	}

	task.Title = title()
	task.Content = content()

	err = task_model.Update(&task)

	if err != nil {
		return utils.CreateResponse(c, 500, "Error updating task")
	}

	return c.SendStatus(200)
}
