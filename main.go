package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"snaptalky/database"
	"snaptalky/routes"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set Gin mode based on GIN_MODE environment variable
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	// Connect to the database
	database.ConnectDatabase()

	// Initialize the Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)

	// Get the port from the environment, or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run the server
	r.Run(":" + port)
}
