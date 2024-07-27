package model

import (
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	ID           uint64 `gorm:"primaryKey;autoIncrement"`
	Original_Url string `gorm:"not null"`
	Short_Url    string
	Visits       int `gorm:"default:0"`
}
