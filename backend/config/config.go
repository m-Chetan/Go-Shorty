package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Sprint("Error Loading .env file!")
	}
	return os.Getenv(key)
}
