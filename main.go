package main

import (
	"github.com/MhmoudGit/go-digital-menu/database"
	"github.com/MhmoudGit/go-digital-menu/models"
)

// Define a slice of model structs that you want to migrate.
var modelsToMigrate = []interface{}{
	models.Category{},
	models.Product{},
	// Add more model structs here if needed.
}

func main() {
	database.Connect()
	database.AutoMigrateDb(modelsToMigrate)
	defer database.Close()
}
