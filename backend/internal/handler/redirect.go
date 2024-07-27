package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/m-Chetan/go-shawty/database"
	"github.com/m-Chetan/go-shawty/internal/model"
)

func RedirectUrl(c *fiber.Ctx) error {
	db := database.DB

	short_url := c.Params("url")

	var url model.Url

	res := db.Where("short_url=?", short_url).Find(&url)

	if res.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Url does not exists"})
	}
	url.Visits++
	fmt.Println(url.Visits)
	db.Save(&url)
	return c.Redirect(url.Original_Url, 301)
}
