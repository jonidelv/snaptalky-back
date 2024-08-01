package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"snaptalky/database"
	"snaptalky/models"
	"snaptalky/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)
	database.ConnectDatabase()
	models.AutoMigrateModels()

	// Initialize the Gin router
	r := gin.Default()
	r.RemoveExtraSlash = true
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	hostPort := host + ":" + port

	// Starting the server
	err := r.Run(hostPort)
	if err != nil {
		log.Printf("Error starting server: %v", err)
		return
	}
}
