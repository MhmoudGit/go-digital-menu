package helpers

import (
	"log"

	"github.com/MhmoudGit/go-digital-menu/models"
	"gorm.io/gorm"
)

func GetProducts(db *gorm.DB, providerID uint) ([]models.Product, error) {
	var products []models.Product
	result := db.Where("provider_id = ?", providerID).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Println(products)
	return products, nil
}

func GetProduct(db *gorm.DB, id uint) (models.Product, error) {
	var product models.Product
	result := db.First(&product, id)
	if result.Error != nil {
		return product, result.Error
	}
	log.Println(product)
	return product, nil
}

func CreateProduct(db *gorm.DB, Product *models.PostProduct) error {
	ProductModel := &models.Product{
		Name:       Product.Name,
		EnName:     Product.EnName,
		Details:    Product.Details,
		EnDetails:  Product.EnDetails,
		Image:      Product.Image,
		Price:      Product.Price,
		Options1:   Product.Options1,
		IsActive:   Product.IsActive,
		CategoryID: Product.CategoryID,
		ProviderID: Product.ProviderID,
	}

	// Create the Product in the database
	result := db.Create(ProductModel)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Product was created successfully....")
	return nil
}

func UpdateProduct(db *gorm.DB, updateProduct *models.UpdateProduct, id uint) error {
	var ProductToUpdate models.Product
	result := db.First(&ProductToUpdate, id).Save(updateProduct)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Product was updated successfully....")
	return nil
}

func UpdateProductImage(db *gorm.DB, ProductImage *models.UpdateProductImage, id uint) error {
	var ProductToUpdate models.Product
	result := db.First(&ProductToUpdate, id).Save(ProductImage)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Product was updated successfully....")
	return nil
}

func DeleteProduct(db *gorm.DB, id uint) error {
	var Product models.Product
	result := db.Delete(&Product, id)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Product was deleted successfully....")
	return nil
}
