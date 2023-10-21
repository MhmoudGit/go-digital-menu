package helpers

import (
	"log"

	"github.com/MhmoudGit/go-digital-menu/models"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func GetUserById(db *gorm.DB, userID uint) ([]models.User, error) {
	var users []models.User
	result := db.Where("user_id = ?", userID).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func GetUser(db *gorm.DB, id uint) (models.User, error) {
	var User models.User
	result := db.Preload("Restaurant").First(&User, id)
	if result.Error != nil {
		return User, result.Error
	}
	return User, nil
}

func GetUserByEmail(db *gorm.DB, email string) (models.User, error) {
	var user models.User
	result := db.Preload("Restaurant").Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func CreateUser(db *gorm.DB, User *models.User) error {
	User.HashPassword(User.Password)
	// Create the User in the database
	result := db.Create(User)
	if result.Error != nil {
		return result.Error
	}
	log.Println("User was created successfully....")
	return nil
}

func UpdateUser(db *gorm.DB, updateUser *models.User, id uint) error {
	var UserToUpdate models.User
	result := db.First(&UserToUpdate, id).Save(updateUser)
	if result.Error != nil {
		return result.Error
	}
	log.Println("User was updated successfully....")
	return nil
}

func DeleteUser(db *gorm.DB, id uint) error {
	var User models.User
	result := db.Delete(&User, id)
	if result.Error != nil {
		return result.Error
	}
	log.Println("User was deleted successfully....")
	return nil
}
