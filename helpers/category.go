package helpers

import (
	"log"

	"github.com/MhmoudGit/go-digital-menu/models"
	"gorm.io/gorm"
)

func GetCategories(db *gorm.DB, providerID uint) ([]models.Category, error) {
	var categories []models.Category
	result := db.Where("provider_id = ?", providerID).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func GetCategory(db *gorm.DB, id uint) (models.Category, error) {
	var category models.Category
	result := db.First(&category, id)
	if result.Error != nil {
		return category, result.Error
	}
	return category, nil
}

func CreateCategory(db *gorm.DB, category *models.PostCategory) error {
	categoryModel := &models.Category{
		Name:       category.Name,
		EnName:     category.EnName,
		Logo:       category.Logo,
		RestaurantID: category.RestaurantID,
	}

	// Create the category in the database
	result := db.Create(categoryModel)
	if result.Error != nil {
		return result.Error
	}
	log.Println("category was created successfully....")
	return nil
}

func UpdateCategory(db *gorm.DB, updateCategory *models.UpdateCategory, id, providerId uint) error {
	var categoryToUpdate models.Category
	result := db.First(&categoryToUpdate, id, providerId).Save(updateCategory)
	if result.Error != nil {
		return result.Error
	}
	log.Println("category was updated successfully....")
	return nil
}

func UpdateCategoryImage(db *gorm.DB, CategoryImage *models.UpdateCategoryImage, id, providerId uint) error {
	var categoryToUpdate models.Category
	result := db.First(&categoryToUpdate, id, providerId).Save(CategoryImage)
	if result.Error != nil {
		return result.Error
	}
	log.Println("category was updated successfully....")
	return nil
}

func DeleteCategory(db *gorm.DB, id, providerId uint) error {
	var category models.Category
	result := db.Delete(&category, id, providerId)
	if result.Error != nil {
		return result.Error
	}
	log.Println("category was deleted successfully....")
	return nil
}
