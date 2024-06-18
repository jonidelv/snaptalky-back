package models

import (
	"snaptalky/database"
	"snaptalky/utils"
)

func InitModels() {
	// AutoMigrate the User model
	if err := database.DB.AutoMigrate(&User{}); err != nil {
		utils.LogError(err, "Failed to migrate database")
		return
	}
}
