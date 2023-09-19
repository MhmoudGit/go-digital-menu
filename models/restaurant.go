package models

import (
	"time"

	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	UserID     uint       `gorm:"not null" json:"userId"`
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
	Discount   int        `gorm:"not null" json:"discount"`
	Tables     int        `gorm:"not null, default:0" json:"tables"`
	IsActive   bool       `gorm:"not null, default:true" json:"isActive"`
	Categories []Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:RestaurantID" json:"-"`
	Products   []Product  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:RestaurantID" json:"-"`
}

// initiate new Restaurant
func NewRestaurant(userId uint, name, enName, image, theme, cover, whatsapp, url, googleMap string, openedFrom, openedTo time.Time, discount, table int) *Restaurant {
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
		Discount:   discount,
		Tables:     table,
		IsActive:   true,
	}
}
