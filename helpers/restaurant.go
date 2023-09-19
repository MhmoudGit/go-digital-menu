package helpers

import (
	// "log"

	"github.com/MhmoudGit/go-digital-menu/models"
	"gorm.io/gorm"
)

func GetRestaurants(db *gorm.DB, RestaurantID uint) ([]models.Restaurant, error) {
	var restaurants []models.Restaurant
	result := db.Where("restaurant_id = ?", RestaurantID).Find(&restaurants)
	if result.Error != nil {
		return nil, result.Error
	}
	return restaurants, nil
}

func GetRestaurant(db *gorm.DB, id uint) (models.Restaurant, error) {
	var restaurant models.Restaurant
	result := db.First(&restaurant, id)
	if result.Error != nil {
		return restaurant, result.Error
	}
	return restaurant, nil
}

// func CreateRestaurant(db *gorm.DB, Restaurant *models.PostRestaurant) error {
// 	RestaurantModel := &models.Restaurant{
// 		Email:       Restaurant.Email,
// 		Image:       Restaurant.Image,
// 		Name:        Restaurant.Name,
// 		EnName:      Restaurant.EnName,
// 		ServiceType: Restaurant.ServiceType,
// 		Whatsapp:    Restaurant.Whatsapp,
// 		Phone:       Restaurant.Phone,
// 		Address:     Restaurant.Address,
// 		EnAddress:   Restaurant.EnAddress,
// 		Facebook:    Restaurant.Facebook,
// 		Theme:       Restaurant.Theme,
// 		OpenedFrom:  Restaurant.OpenedFrom,
// 		OpenedTo:    Restaurant.OpenedTo,
// 		Url:         Restaurant.Url,
// 	}
// 	RestaurantModel.HashPassword(Restaurant.Password)

// 	// Create the Restaurant in the database
// 	result := db.Create(RestaurantModel)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	log.Println("Restaurant was created successfully....")
// 	return nil
// }

// func UpdateRestaurant(db *gorm.DB, updateRestaurant *models.UpdateRestaurant, id uint) error {
// 	var RestaurantToUpdate models.Restaurant
// 	result := db.First(&RestaurantToUpdate, id).Save(updateRestaurant)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	log.Println("Restaurant was updated successfully....")
// 	return nil
// }

// func UpdateRestaurantImage(db *gorm.DB, RestaurantImage string, id uint) error {
// 	var RestaurantToUpdate models.Restaurant
// 	result := db.First(&RestaurantToUpdate, id).Save(RestaurantImage)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	log.Println("Restaurant was updated successfully....")
// 	return nil
// }

// func DeleteRestaurant(db *gorm.DB, id uint) error {
// 	var Restaurant models.Restaurant
// 	result := db.Delete(&Restaurant, id)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	log.Println("Restaurant was deleted successfully....")
// 	return nil
// }
