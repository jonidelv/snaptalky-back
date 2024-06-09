package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"snaptalky/database"
	"snaptalky/routes"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	database.ConnectDatabase()

	// Initialize the Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)

	// Run the server
	r.Run()
}
