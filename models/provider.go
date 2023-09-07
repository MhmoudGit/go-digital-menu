package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
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

// Verify Password.
func (p *Provider) VerifyPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password))
	return err
}

// HashPassword securely hashes the provided password and sets it in the PasswordHash field.
func (p *Provider) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.Password = string(hashedPassword)
	return nil
}
