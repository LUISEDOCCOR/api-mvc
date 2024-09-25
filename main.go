package main

import (
	"github.com/LUISEDOCCOR/api-mvc/database"
	"github.com/LUISEDOCCOR/api-mvc/middlewares"
	"github.com/LUISEDOCCOR/api-mvc/models"
	task_routes "github.com/LUISEDOCCOR/api-mvc/routes/task"
	user_routes "github.com/LUISEDOCCOR/api-mvc/routes/user"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

const rootEndpoint = "/api/v1"

func init() {
	_ = godotenv.Load()
}

func main() {
	database.Conn()
	database.DB.AutoMigrate(models.Task{})
	database.DB.AutoMigrate(models.User{})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	UserRouter := app.Group(rootEndpoint + "/user")
	user_routes.Router(UserRouter)

	TaskRouter := app.Group(rootEndpoint+"/task", middlewares.AuthMiddleware())
	task_routes.Router(TaskRouter)

	app.Listen(":4321")
}
