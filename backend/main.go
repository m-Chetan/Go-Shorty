package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/m-Chetan/go-shawty/database"

	"github.com/m-Chetan/go-shawty/router"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("API is UP...")
		if err != nil {
			return err
		}
		return nil
	})

	app.Listen(":3000")
}
