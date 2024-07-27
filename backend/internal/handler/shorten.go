package handler

import (
	"log"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/m-Chetan/go-shawty/config"
	"github.com/m-Chetan/go-shawty/database"
	"github.com/m-Chetan/go-shawty/internal/helpers"
	"github.com/m-Chetan/go-shawty/internal/model"
)

func createShortUrl() string {
	db := database.DB
	var url model.Url
	randNum := rand.Uint64()

	short_url := helpers.Base62Encode(randNum)

	res := db.Where("short_url=?", short_url).First(&url)

	if res.Error != nil {
		return short_url
	}

	return ""
}

func ShortenUrl(c *fiber.Ctx) error {
	db := database.DB

	url := new(model.Url)
	err := c.BodyParser(url)
	if err != nil {
		log.Fatal(err)
	}

	short_url := createShortUrl()

	check_url := db.Where("original_url=?", url.Original_Url).Find(&url)

	if short_url == "" || check_url.Error != nil {
		return c.Status(fiber.StatusForbidden).JSON(
			fiber.Map{"error": "Short Url already exists try again"})
	}

	url.Short_Url = short_url

	res := db.Create(&url)

	if res.Error != nil {
		return res.Error
	}

	response := model.Url{
		ID:           url.ID,
		Original_Url: url.Original_Url,
		Short_Url:    config.Config("DOMAIN") + "/" + url.Short_Url,
		Visits:       url.Visits,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
