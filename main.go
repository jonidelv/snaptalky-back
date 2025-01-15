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
	isProduction := os.Getenv("ENV") == "production"

	// Load .env file only if not in production
	if !isProduction {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, relying solely on environment variables")
		}
	}

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

	gin.SetMode(ginMode)

	database.ConnectDatabase()

	models.AutoMigrateModels()

	r := gin.Default()
	r.RemoveExtraSlash = true

	routes.SetupRoutes(r)

	hostPort := fmt.Sprintf("%s:%s", host, port)

	if err := r.Run(hostPort); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
