package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name       string `gorm:"not null"`
	EnName     string `gorm:"not null"`
	Details    string `gorm:"not null"`
	EnDetails  string `gorm:"not null"`
	Image      string
	Price      int
	Options1   []Options `gorm:"type:jsonb"`
	IsActive   bool      `gorm:"not null, default:'true'"`
	CategoryID uint      `gorm:"not null"`
	ProviderID uint      `gorm:"not null"`
}

type Options struct {
	Name   string
	EnName string
	Price  int
}
