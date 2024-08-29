package main

import (
	"github.com/LUISEDOCCOR/api-mvc/database"
	"github.com/LUISEDOCCOR/api-mvc/models"
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

	UserRouter := app.Group(rootEndpoint + "/user")
	user_routes.Router(UserRouter)

	app.Listen(":3000")
}
