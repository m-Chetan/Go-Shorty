package database

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2/log"
	"github.com/m-Chetan/go-shorty/config"
	"github.com/m-Chetan/go-shorty/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	p := config.Config("DB_PORT")

	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Error("Invalid port number!")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	log.Info("Connection Opened to Database")

	DB.AutoMigrate(&model.Url{})
	log.Info("Database Migrated")

}
