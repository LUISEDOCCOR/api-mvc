package task_routes

import (
	task_controller "github.com/LUISEDOCCOR/api-mvc/contollers/task"
	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	router.Get("/", task_controller.GetALL)
	router.Post("/create", task_controller.Create)
}