package models

import (
	"time"

	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	UserID     uint       `gorm:"not null;unique" json:"userId"`
	Name       string     `gorm:"not null" json:"name"`
	EnName     string     `gorm:"not null" json:"enName"`
	Image      string     `json:"image"`
	Theme      string     `gorm:"not null" json:"theme"`
	Cover      string     `gorm:"not null" json:"cover"`
	Whatsapp   string     `gorm:"not null" json:"whatsapp"`
	OpenedFrom time.Time  `gorm:"not null" json:"openedFrom"`
	OpenedTo   time.Time  `gorm:"not null" json:"openedTo"`
	Url        string     `json:"url"`
	GoogleMap  string     `json:"googleMap"`
	Discount   int        `gorm:"not null;default:0" json:"discount"`
	Tables     int        `gorm:"not null;default:0" json:"tables"`
	IsActive   bool       `gorm:"not null;default:true" json:"isActive"`
	Categories []Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:RestaurantID" json:"-"`
	Products   []Product  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:RestaurantID" json:"-"`
}

// initiate new Restaurant
func NewRestaurant(userId uint, name, enName, image, theme, cover, whatsapp, url, googleMap string, openedFrom, openedTo time.Time, tables int) *Restaurant {
	return &Restaurant{
		UserID:     userId,
		Name:       name,
		EnName:     enName,
		Image:      image,
		Theme:      theme,
		Cover:      cover,
		Whatsapp:   whatsapp,
		OpenedFrom: openedFrom,
		OpenedTo:   openedTo,
		Url:        url,
		GoogleMap:  googleMap,
		Discount:   0,
		Tables:     tables,
		IsActive:   true,
	}
}

type UpdateRestaurant struct {
	Name       string    `json:"name"`
	EnName     string    `json:"enName"`
	Whatsapp   string    `json:"whatsapp"`
	OpenedFrom time.Time `json:"openedFrom"`
	OpenedTo   time.Time `json:"openedTo"`
	Url        string    `json:"url"`
	GoogleMap  string    `json:"googleMap"`
	Discount   int       `json:"discount"`
	Tables     int       `json:"tables"`
	IsActive   bool      `json:"isActive"`
}

type UpdateRestaurantImage struct {
	Image string `json:"image"`
}

type UpdateRestaurantTheme struct {
	Theme string `json:"theme"`
}

type UpdateRestaurantCover struct {
	Cover string `json:"cover"`
}
