package main

import (
	"github.com/MhmoudGit/go-digital-menu/database"
)

func main() {
	database.Connect()
	database.AutoMigrateDb()
	defer database.Close()
}
