package handler

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/m-Chetan/go-shorty/config"
	"github.com/m-Chetan/go-shorty/database"
	"github.com/m-Chetan/go-shorty/internal/helpers"
	"github.com/m-Chetan/go-shorty/internal/model"
)

func createShortUrl() string {
	db := database.DB
	var url model.Url
	randNum := rand.Uint64()
	short_url := helpers.Base62Encode(randNum)

	log.Info("Search if short url already exists.")

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

	if err != nil || url.Original_Url == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"error": "Enter a valid url"})
	}

	log.Info("Generate short url")
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
		Short_Url:    "localhost:" + config.Config("DOMAIN") + "/" + url.Short_Url,
		Visits:       url.Visits,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
