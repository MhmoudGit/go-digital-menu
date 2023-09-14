package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Provider struct {
	gorm.Model
	Email       string     `gorm:"not null;index;unique" json:"email"`
	Password    string     `gorm:"not null" json:"-"`
	Image       string     `json:"image"`
	Name        string     `gorm:"not null" json:"name"`
	EnName      string     `gorm:"not null" json:"enName"`
	ServiceType string     `gorm:"not null" json:"serviceType"`
	Whatsapp    string     `gorm:"not null" json:"whatsapp"`
	Phone       string     `gorm:"not null" json:"phone"`
	Address     string     `gorm:"not null" json:"adress"`
	EnAddress   string     `gorm:"not null" json:"enAdress"`
	Facebook    string     `json:"facebook"`
	Theme       string     `gorm:"not null" json:"theme"`
	OpenedFrom  time.Time  `gorm:"not null" json:"openedFrom"`
	OpenedTo    time.Time  `gorm:"not null" json:"openedTo"`
	Url         string     `json:"url"`
	IsActive    bool       `gorm:"not null;default:true" json:"isActive"`
	Categories  []Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ProviderID" json:"-"`
	Products    []Product  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ProviderID" json:"-"`
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

type PostProvider struct {
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Image       string    `json:"image"`
	Name        string    `json:"name"`
	EnName      string    `json:"enName"`
	ServiceType string    `json:"serviceType"`
	Whatsapp    string    `json:"whatsapp"`
	Phone       string    `json:"phone"`
	Address     string    `json:"adress"`
	EnAddress   string    `json:"enAdress"`
	Facebook    string    `json:"facebook"`
	Theme       string    `json:"theme"`
	OpenedFrom  time.Time `json:"openedFrom"`
	OpenedTo    time.Time `json:"openedTo"`
	Url         string    `json:"url"`
}

type UpdateProvider struct {
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	EnName      string    `json:"enName"`
	ServiceType string    `json:"serviceType"`
	Whatsapp    string    `json:"whatsapp"`
	Phone       string    `json:"phone"`
	Address     string    `json:"adress"`
	EnAddress   string    `json:"enAdress"`
	Facebook    string    `json:"facebook"`
	Theme       string    `json:"theme"`
	OpenedFrom  time.Time `json:"openedFrom"`
	OpenedTo    time.Time `json:"openedTo"`
	Url         string    `json:"url"`
}

type UpdateProviderImage struct {
	Image string `json:"image"`
}
