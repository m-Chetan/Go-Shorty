package database

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/m-Chetan/go-shorty/config"
	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

var Client *redis.Client

func ConnectCache() {
	Client = redis.NewClient(
		&redis.Options{
			Addr:     config.Config("REDIS_ADDR"),
			Password: config.Config("REDIS_PASSWORD"),
			DB:       0,
		},
	)

	pong, err := Client.Ping(Ctx).Result()
	if err != nil {
		log.Error("Redis unable to connect ", err)
	} else {
		log.Info("Redis connected ", pong)
	}
}
