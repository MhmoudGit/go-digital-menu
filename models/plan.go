package models

import (
	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	Name     string   `gorm:"not null" json:"name"`
	EnName   string   `gorm:"not null" json:"enName"`
	Features []string `gorm:"not null" json:"features"`
	Price    float64  `gorm:"not null" json:"price"`
	Duration int      `gorm:"not null" json:"duration"`
	Discount float64  `gorm:"not null, default:0" json:"discount"`
	Users    []User   `gorm:"foreignKey:UserID" json:"-"`
}

// initiate new Plan
func NewPlan(name, enName string, features []string, price, discount float64, duration int) *Plan {
	return &Plan{
		Name:     name,
		EnName:   enName,
		Features: features,
		Price:    price,
		Duration: duration,
		Discount: discount,
	}
}
