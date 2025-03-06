package models

import (
	"log"

	"github.com/jonidelv/snaptalky-back/database"
	"github.com/jonidelv/snaptalky-back/utils"
)

func AutoMigrateModels() {
	// AutoMigrate the User model
	if err := database.DB.AutoMigrate(&User{}, &Response{}); err != nil {
		utils.LogError(err, "Failed to migrate database", utils.Object{"path": "models/setup.go"})
		return
	}
	log.Println("Database migration completed successfully.")
}
