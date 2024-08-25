package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/m-Chetan/go-shorty/config"
	"github.com/m-Chetan/go-shorty/database"

	"github.com/m-Chetan/go-shorty/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	database.ConnectDB()
	database.ConnectCache()

	router.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		log.Info("Application")
		err := c.SendString("API is UP...")
		if err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(
				fiber.Map{"error": "Service Unavailable"})
		}
		return nil
	})

	app.Listen(":" + config.Config("DOMAIN"))
}
