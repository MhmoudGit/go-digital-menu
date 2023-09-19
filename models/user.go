package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email      string    `gorm:"not null;index;unique" json:"email"`
	Password   string    `gorm:"not null" json:"-"`
	Name       string    `gorm:"not null" json:"name"`
	Phone      string    `gorm:"not null" json:"phone"`
	StartDate  time.Time `gorm:"not null" json:"startDate"`
	EndDate    time.Time `gorm:"not null" json:"endDate"`
	Paid       bool      `gorm:"not null, default:'false'" json:"paid"`
	IsVerified bool      `gorm:"not null, default:'false'" json:"isVerified"`
	IsActive   bool      `gorm:"not null, default:'true'" json:"isActive"`
	PlanID     uint      `gorm:"not null" json:"planId"`
	// Restaurant  Restaurant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID" json:"-"`
}

// Verify Password.
func (u *User) VerifyPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err
}

// HashPassword securely hashes the provided password and sets it in the PasswordHash field.
func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// initiate new user
func NewUser(duration int, planId uint, email, password, name, phone string) *User {
	return &User{
		Email:     email,
		Password:  password,
		Name:      name,
		Phone:     phone,
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, duration, 0),
		Paid:      true,
		IsVerified: false,
		IsActive:  true,
		PlanID:    planId,
	}
}
