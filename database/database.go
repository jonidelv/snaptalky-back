package database

import (
  "fmt"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "os"
  "snaptalky/models"
  "snaptalky/utils"
)

var DB *gorm.DB

func ConnectDatabase() {
  // Load environment variables from .env file
  dbHost := os.Getenv("DB_HOST")
  dbUser := os.Getenv("DB_USER")
  dbPassword := os.Getenv("DB_PASSWORD")
  dbName := os.Getenv("DB_NAME")
  dbPort := os.Getenv("DB_PORT")

  // Create the DSN (Data Source Name)
  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
    dbHost, dbUser, dbPassword, dbName, dbPort)

  // Connect to the database
  database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    utils.LogError(err, "Failed to connect to database")
    return
  }

  // AutoMigrate the User model
  if err := database.AutoMigrate(&models.User{}); err != nil {
    utils.LogError(err, "Failed to migrate database")
    return
  }

  DB = database
}
