package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	Name     string      `gorm:"not null" json:"name"`
	EnName   string      `gorm:"not null" json:"enName"`
	Features StringSlice `gorm:"not null" json:"features"`
	Price    float64     `gorm:"not null" json:"price"`
	Duration int         `gorm:"not null" json:"duration"`
	Discount float64     `gorm:"not null, default:0" json:"discount"`
	Users    []User      `gorm:"foreignKey:PlanID" json:"-"`
}

// StringSlice is a custom data type to represent a slice of strings in the database.
type StringSlice []string

// Value converts the StringSlice to a JSON-encoded string to be stored in the database.
func (s StringSlice) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// Scan converts the JSON-encoded string from the database into a StringSlice.
func (s *StringSlice) Scan(value interface{}) error {
	if value == nil {
		*s = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("invalid type")
	}
	return json.Unmarshal(bytes, s)
}

// initiate new Plan
func NewPlan(name, enName string, features StringSlice, price, discount float64, duration int) *Plan {
	return &Plan{
		Name:     name,
		EnName:   enName,
		Features: features,
		Price:    price,
		Duration: duration,
		Discount: discount,
	}
}
