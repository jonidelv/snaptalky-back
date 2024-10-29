package models

import (
  "github.com/jonidelv/snaptalky-back/database"
  "github.com/jonidelv/snaptalky-back/utils"
  "log"
)

func AutoMigrateModels() {
  // AutoMigrate the User model
  if err := database.DB.AutoMigrate(&User{}, &Response{}); err != nil {
    utils.LogError(err, "Failed to migrate database")
    return
  }
  log.Println("Database migration completed successfully.")
}
