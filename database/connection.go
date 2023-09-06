package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	var err error
	// Load the environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Create a connection to the PostgreSQL database.
	dsn := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", dbUserName, dbPassword, dbname)

	// Open db connection
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	} else {
		log.Println("database connected successfully...")
	}
}

func AutoMigrateDb(models ...interface{}) {
	// // AutoMigrate will create tables if they don't exist based on the model structs.
	err := Db.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("Error migrating database tables: %v", err)
	}
	log.Println("Tables created/updated successfully...")
}

func Close() {
	// Close db
	dbInstance, _ := Db.DB()
	_ = dbInstance.Close()
	log.Println("database is closed successfully...")
}
