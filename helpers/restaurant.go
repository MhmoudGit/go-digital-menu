package helpers

import (
	"log"

	"github.com/MhmoudGit/go-digital-menu/models"
	"gorm.io/gorm"
)

// func GetRestaurants(db *gorm.DB, RestaurantID uint) ([]models.Restaurant, error) {
// 	var restaurants []models.Restaurant
// 	result := db.Where("restaurant_id = ?", RestaurantID).Find(&restaurants)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return restaurants, nil
// }

func GetRestaurant(db *gorm.DB, id uint) (models.Restaurant, error) {
	var restaurant models.Restaurant
	result := db.First(&restaurant, id)
	if result.Error != nil {
		return restaurant, result.Error
	}
	return restaurant, nil
}

func CreateRestaurant(db *gorm.DB, Restaurant *models.Restaurant) error {
	// Create the Restaurant in the database
	result := db.Create(Restaurant)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Restaurant was created successfully....")
	return nil
}

func UpdateRestaurant(db *gorm.DB, updateRestaurant *models.UpdateRestaurant, id, userId uint) error {
	var RestaurantToUpdate models.Restaurant
	result := db.First(&RestaurantToUpdate, id, userId).Save(updateRestaurant)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Restaurant was updated successfully....")
	return nil
}

func UpdateRestaurantImage(db *gorm.DB, RestaurantImage *models.UpdateRestaurantImage, id uint) error {
	var RestaurantToUpdate models.Restaurant
	result := db.First(&RestaurantToUpdate, id).Save(RestaurantImage)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Restaurant was updated successfully....")
	return nil
}

func UpdateRestaurantCover(db *gorm.DB, RestaurantCover *models.UpdateRestaurantCover, id uint) error {
	var RestaurantToUpdate models.Restaurant
	result := db.First(&RestaurantToUpdate, id).Save(RestaurantCover)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Restaurant was updated successfully....")
	return nil
}

func UpdateRestaurantTheme(db *gorm.DB, RestaurantTheme *models.UpdateRestaurantTheme, id uint) error {
	var RestaurantToUpdate models.Restaurant
	result := db.First(&RestaurantToUpdate, id).Save(RestaurantTheme)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Restaurant was updated successfully....")
	return nil
}

func DeleteRestaurant(db *gorm.DB, id uint) error {
	var Restaurant models.Restaurant
	result := db.Delete(&Restaurant, id)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Restaurant was deleted successfully....")
	return nil
}
