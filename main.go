package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jonidelv/snaptalky-back/database"
	"github.com/jonidelv/snaptalky-back/models"
	"github.com/jonidelv/snaptalky-back/routes"
)

func main() {
	// Determine if the application is running in production
	isProduction := os.Getenv("ENV") == "production"

	// Load .env file only if not in production
	if !isProduction {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, relying solely on environment variables")
		}
	}

	// Retrieve required environment variables
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		log.Fatal("Required environment variable GIN_MODE is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Required environment variable PORT is not set")
	}

	host := os.Getenv("HOST")
	if host == "" {
		log.Fatal("Required environment variable HOST is not set")
	}

	// (Optional) Retrieve other environment variables as needed
	// For example, if your database connection relies on DATABASE_URL internally,
	// ensure that it's set within the database package or handle it here.

	// Set Gin mode
	gin.SetMode(ginMode)

	// Initialize database
	database.ConnectDatabase()

	// Auto-migrate database models
	models.AutoMigrateModels()

	// Initialize the Gin router
	r := gin.Default()
	r.RemoveExtraSlash = true

	// Setup routes
	routes.SetupRoutes(r)

	// Construct host and port
	hostPort := fmt.Sprintf("%s:%s", host, port)

	// Start the server
	if err := r.Run(hostPort); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
