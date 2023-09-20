package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name         string    `json:"name"`
	EnName       string    `json:"enName"`
	Logo         string    `json:"logo"`
	RestaurantID uint      `gorm:"not null" json:"restaurantID"`
	UserID       uint      `gorm:"not null" json:"-"`
	Products     []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CategoryID" json:"-"`
}

// Schema
type UpdateCategory struct {
	Name   string `json:"name"`
	EnName string `json:"enName"`
}

type UpdateCategoryImage struct {
	Logo string `json:"logo"`
}

type PostCategory struct {
	Name         string `json:"name"`
	EnName       string `json:"enName"`
	Logo         string `json:"logo"`
	RestaurantID uint   `json:"restaurantID"`
	UserID       uint   `json:"userID"`
}
