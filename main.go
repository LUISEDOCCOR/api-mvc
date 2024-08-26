package main

import (
	"github.com/LUISEDOCCOR/api-mvc/database"
	"github.com/LUISEDOCCOR/api-mvc/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Conn()
	database.DB.AutoMigrate(models.Task{})
	database.DB.AutoMigrate(models.User{})

	app := fiber.New()
	router := app.Group("/api/v1")

	router.Get("/ok", func(c *fiber.Ctx) error {
		return c.SendString("Is Ok")
	})

	app.Listen(":3000")
}
