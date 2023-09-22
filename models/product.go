package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"log"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string     `gorm:"not null" json:"name"`
	EnName       string     `gorm:"not null" json:"enName"`
	Details      string     `gorm:"not null" json:"details"`
	EnDetails    string     `gorm:"not null" json:"enDetails"`
	Image        string     `json:"image"`
	Price        int        `json:"price"`
	Options1     OptionsArr `json:"options1"`
	Options2     OptionsArr `json:"options2"`
	IsActive     bool       `gorm:"not null;default:true" json:"isActive"`
	CategoryID   uint       `gorm:"not null" json:"categoryId"`
	RestaurantID uint       `gorm:"not null" json:"restaurantID"`
}

type Options struct {
	Name   string `json:"name"`
	EnName string `json:"enName"`
	Price  int    `json:"price"`
}

// schemas
type PostProduct struct {
	Name         string     `json:"name"`
	EnName       string     `json:"enName"`
	Details      string     `json:"details"`
	EnDetails    string     `json:"enDetails"`
	Image        string     `json:"image"`
	Price        int        `json:"price"`
	Options1     OptionsArr `json:"options1"`
	Options2     OptionsArr `json:"options2"`
	IsActive     bool       `json:"isActive"`
	CategoryID   uint       `json:"categoryId"`
	RestaurantID uint       `json:"restaurantID"`
}

type UpdateProduct struct {
	Name      string     `json:"name"`
	EnName    string     `json:"enName"`
	Details   string     `json:"details"`
	EnDetails string     `json:"enDetails"`
	Price     int        `json:"price"`
	Options1  OptionsArr `json:"options1"`
	Options2  OptionsArr `json:"options2"`
	IsActive  bool       `json:"isActive"`
}

type UpdateProductImage struct {
	Image string `json:"image"`
}

func CreateOptions(o string) []Options {
	var options []Options
	// Unmarshal the JSON string into the struct
	err := json.Unmarshal([]byte(o), &options)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return options
}

type OptionsArr []Options

// Value converts the Options to a JSON-encoded string to be stored in the database.
func (oa OptionsArr) Value() (driver.Value, error) {
	return json.Marshal(oa)
}

// Scan converts the JSON-encoded string from the database into a Options.
func (oa *OptionsArr) Scan(value interface{}) error {
	if value == nil {
		*oa = []Options{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("invalid type")
	}
	return json.Unmarshal(bytes, oa)
}
