package user_routes

import (
	user_controller "github.com/LUISEDOCCOR/api-mvc/contollers/user"
	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	router.Post("/create", user_controller.Create)
	router.Post("/login", user_controller.Login)
}
