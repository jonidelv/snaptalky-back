package models

import (
	"log"
	"snaptalky/database"
	"snaptalky/utils"
)

func AutoMigrateModels() {
	// AutoMigrate the User model
	if err := database.DB.AutoMigrate(&User{}); err != nil {
		utils.LogError(err, "Failed to migrate database")
		return
	}
	log.Println("Database migration completed successfully.")
}
