package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/m-Chetan/go-shorty/internal/handler"
)

func SetupRoutes(app *fiber.App) {

	app.Post("shorten", handler.ShortenUrl)
	app.Get("/:url", handler.RedirectUrl)
}
