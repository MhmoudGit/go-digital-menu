package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name       string `gorm:"not null"`
	EnName     string `gorm:"not null"`
	Logo       string `gorm:"not null"`
	Products   []Product
	ProviderID uint
}
