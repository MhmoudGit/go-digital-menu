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
	Name   string `json:"name"`
	EnName string `json:"enName"`
	Price  int    `json:"price"`
}

// schemas
type PostProduct struct {
	Name       string    `json:"name"`
	EnName     string    `json:"enName"`
	Details    string    `json:"details"`
	EnDetails  string    `json:"enDetails"`
	Image      string    `json:"image"`
	Price      int       `json:"price"`
	Options1   []Options `json:"options1"`
	IsActive   bool      `json:"isActive"`
	CategoryID uint      `json:"categoryId"`
	ProviderID uint      `json:"providerId"`
}

type GetProduct struct{
	gorm.Model
	PostProduct
}
