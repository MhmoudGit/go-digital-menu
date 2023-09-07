package controllers

import (
	"log"

	"github.com/MhmoudGit/go-digital-menu/models"
	"gorm.io/gorm"
)

func GetCategories(db *gorm.DB, providerID uint) error {
	var categories []models.Category
	result := db.Where("provider_id = ?", providerID).Find(&categories)
	if result.Error != nil {
		return result.Error
	}
	log.Println(categories)
	return nil
}

func GetCategory(db *gorm.DB, id uint) error {
	var category models.Category
	result := db.First(&category, id)
	if result.Error != nil {
		return result.Error
	}
	log.Println(category)
	return nil
}

func CreateCategory(db *gorm.DB, category *models.PostCategory) error {
	categoryModel := &models.Category{
		Name:       category.Name,
		EnName:     category.EnName,
		Logo:       category.Logo,
		ProviderID: category.ProviderID,
	}

	// Create the category in the database
	result := db.Create(categoryModel)

	if result.Error != nil {
		return result.Error
	}
	log.Println("category was created successfully....")
	return nil
}

func DeleteCategory(db *gorm.DB, id uint) error {
	var category models.Category
	result := db.Delete(&category, id)
	if result.Error != nil {
		return result.Error
	}
	log.Println("category was deleted successfully....")
	return nil
}
