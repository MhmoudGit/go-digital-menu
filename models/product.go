package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name       string    `gorm:"not null" json:"name"`
	EnName     string    `gorm:"not null" json:"enName"`
	Details    string    `gorm:"not null" json:"details"`
	EnDetails  string    `gorm:"not null" json:"enDetails"`
	Image      string    `json:"image"`
	Price      int       `json:"price"`
	Options1   []Options `gorm:"type:jsonb" json:"options1"`
	IsActive   bool      `gorm:"not null, default:'true'" json:"isActive"`
	CategoryID uint      `gorm:"not null" json:"categoryId"`
	ProviderID uint      `gorm:"not null" json:"providerId"`
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

type UpdateProduct struct {
	Name       string    `json:"name"`
	EnName     string    `json:"enName"`
	Details    string    `json:"details"`
	EnDetails  string    `json:"enDetails"`
	Price      int       `json:"price"`
	Options1   []Options `json:"options1"`
	IsActive   bool      `json:"isActive"`
	CategoryID uint      `json:"categoryId"`
	ProviderID uint      `json:"providerId"`
}

type UpdateProductImage struct {
	Image string `json:"image"`
}
