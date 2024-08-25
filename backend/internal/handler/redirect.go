package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/m-Chetan/go-shorty/database"
	"github.com/m-Chetan/go-shorty/internal/model"
	"github.com/redis/go-redis/v9"
)

func RedirectUrl(c *fiber.Ctx) error {
	shortUrl := c.Params("url")
	client := database.Client
	ctx := database.Ctx

	actualUrl, err := checkCache(shortUrl, client, ctx)

	if err != nil {
		log.Error("Url not found in cache!")
	} else {
		log.Infof("Redirecting to %s from cache", actualUrl)
		return c.Redirect(actualUrl, 301)
	}

	actualUrl, err = fetchFromDB(shortUrl, client, ctx)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Url does not exists"})
	}

	log.Info("Redirecting to ", actualUrl)

	return c.Redirect(actualUrl, 301)
}

func fetchFromDB(shortUrl string, client *redis.Client, ctx context.Context) (string, error) {
	var url model.Url
	db := database.DB

	db.Where("short_url=?", shortUrl).Find(&url)
	if url.Original_Url == "" {
		log.Error("No record found in db")
		return "", fiber.NewError(404, "No record found in db")
	}

	response := client.Set(ctx, shortUrl, url.Original_Url, 0)
	if response.Val() != "OK" {
		log.Error("Cache error while storing url")
	}

	url.Visits++
	db.Save(&url)

	return url.Original_Url, nil
}

func checkCache(shortUrl string, client *redis.Client, ctx context.Context) (string, error) {
	value, err := client.Get(ctx, shortUrl).Result()

	if err != nil {
		return "", err
	}

	return value, nil
}
