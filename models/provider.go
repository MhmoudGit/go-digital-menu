package models

import (
	"time"

	"gorm.io/gorm"
)

type Provider struct {
	gorm.Model
	Email       string `gorm:"not null;index;unique"`
	Password    string `gorm:"not null"`
	Image       string
	Name        string `gorm:"not null"`
	EnName      string `gorm:"not null"`
	ServiceType string `gorm:"not null"`
	Whatsapp    string `gorm:"not null"`
	Phone       string `gorm:"not null"`
	Address     string `gorm:"not null"`
	EnAddress   string `gorm:"not null"`
	Facebook    string
	Theme       string    `gorm:"not null"`
	OpenedFrom  time.Time `gorm:"not null"`
	OpenedTo    time.Time `gorm:"not null"`
	Url         string
	IsActive    bool `gorm:"not null, default:'true'"`
	Categories  []Category
	Products    []Product
}
