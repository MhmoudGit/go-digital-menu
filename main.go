package main

import (
	"github.com/MhmoudGit/go-digital-menu/database"
	"github.com/MhmoudGit/go-digital-menu/models"
)

func main() {
	// Define a slice of model structs that you want to migrate.
	modelsToMigrate := []interface{}{
		&models.Provider{},
		&models.Category{},
		&models.Product{},
		// Add more model structs here if needed.
	}

	database.Connect()
	database.AutoMigrateDb(modelsToMigrate...)
	defer database.Close()
}

//create category
// category := &models.PostCategory{
// 	UpdateCategory: models.UpdateCategory{
// 		Name:   "Category Name",
// 		EnName: "Category EnName",
// 		Logo:   "Category Logo",
// 	},
// 	ProviderID: 1, // Example value for ProviderID
// }
