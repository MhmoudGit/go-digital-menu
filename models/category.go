package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name       string    `json:"name"`
	EnName     string    `json:"enName"`
	Logo       string    `json:"logo"`
	ProviderID uint      `gorm:"not null" json:"providerId"`
	Products   []Product `json:"-"`
}

// Schema
type UpdateCategory struct {
	Name   string `json:"name"`
	EnName string `json:"enName"`
	Logo   string `json:"logo"`
}

type PostCategory struct {
	Name       string `json:"name"`
	EnName     string `json:"enName"`
	Logo       string `json:"logo"`
	ProviderID uint   `json:"providerId"`
}

// type GetCategory struct {
// 	gorm.Model
// 	Name       string `json:"name"`
// 	EnName     string `json:"enName"`
// 	Logo       string `json:"logo"`
// 	ProviderID uint   `json:"providerId"`
// }
