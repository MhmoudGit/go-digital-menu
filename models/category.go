package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name       string `gorm:"not null"`
	EnName     string `gorm:"not null"`
	Logo       string `gorm:"not null"`
	ProviderID uint   `gorm:"not null"`
	Products   []Product
}

// Schema
type UpdateCategory struct {
	Name   string `json:"name"`
	EnName string `json:"enName"`
	Logo   string `json:"logo"`
}

type PostCategory struct {
	UpdateCategory
	ProviderID uint `json:"providerId"`
}

type GetCategory struct {
	gorm.Model
	PostCategory
}
