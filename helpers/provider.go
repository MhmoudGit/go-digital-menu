package helpers

import (
	"log"

	"github.com/MhmoudGit/go-digital-menu/models"
	"gorm.io/gorm"
)

func GetProviders(db *gorm.DB, providerID uint) ([]models.Provider, error) {
	var providers []models.Provider
	result := db.Where("provider_id = ?", providerID).Find(&providers)
	if result.Error != nil {
		return nil, result.Error
	}
	return providers, nil
}

func GetProvider(db *gorm.DB, id uint) (models.Provider, error) {
	var Provider models.Provider
	result := db.First(&Provider, id)
	if result.Error != nil {
		return Provider, result.Error
	}
	return Provider, nil
}

func GetProviderByEmail(db *gorm.DB, email string) (models.Provider, error) {
	var Provider models.Provider
	result := db.Where("email = ?", email).First(&Provider)
	if result.Error != nil {
		return Provider, result.Error
	}
	return Provider, nil
}

func CreateProvider(db *gorm.DB, Provider *models.PostProvider) error {
	ProviderModel := &models.Provider{
		Email:       Provider.Email,
		Image:       Provider.Image,
		Name:        Provider.Name,
		EnName:      Provider.EnName,
		ServiceType: Provider.ServiceType,
		Whatsapp:    Provider.Whatsapp,
		Phone:       Provider.Phone,
		Address:     Provider.Address,
		EnAddress:   Provider.EnAddress,
		Facebook:    Provider.Facebook,
		Theme:       Provider.Theme,
		OpenedFrom:  Provider.OpenedFrom,
		OpenedTo:    Provider.OpenedTo,
		Url:         Provider.Url,
	}
	ProviderModel.HashPassword(Provider.Password)

	// Create the Provider in the database
	result := db.Create(ProviderModel)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Provider was created successfully....")
	return nil
}

func UpdateProvider(db *gorm.DB, updateProvider *models.UpdateProvider, id uint) error {
	var ProviderToUpdate models.Provider
	result := db.First(&ProviderToUpdate, id).Save(updateProvider)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Provider was updated successfully....")
	return nil
}

func UpdateProviderImage(db *gorm.DB, ProviderImage *models.UpdateProviderImage, id uint) error {
	var ProviderToUpdate models.Provider
	result := db.First(&ProviderToUpdate, id).Save(ProviderImage)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Provider was updated successfully....")
	return nil
}

func DeleteProvider(db *gorm.DB, id uint) error {
	var Provider models.Provider
	result := db.Delete(&Provider, id)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Provider was deleted successfully....")
	return nil
}
