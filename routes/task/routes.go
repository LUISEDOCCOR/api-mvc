package task_routes

import (
	task_controller "github.com/LUISEDOCCOR/api-mvc/contollers/task"
	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	router.Get("/:id", task_controller.GetById)
	router.Get("/", task_controller.GetALL)
	router.Get("/finish/:id", task_controller.Finish)
	router.Delete("/:id", task_controller.Delete)
	router.Post("/create", task_controller.Create)
	router.Put("/:id", task_controller.Update)
}
