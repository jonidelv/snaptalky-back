package database

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jonidelv/snaptalky-back/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Check if DB_URL is set
	dbURL := os.Getenv("DB_URL")
	var dsn string

	if dbURL != "" {
		dsn = dbURL
	} else {
		// Load individual environment variables
		dbHost := os.Getenv("DB_HOST")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")
		dbPort := os.Getenv("DB_PORT")

		// Validate required environment variables
		missingVars := []string{}
		if dbHost == "" {
			missingVars = append(missingVars, "DB_HOST")
		}
		if dbUser == "" {
			missingVars = append(missingVars, "DB_USER")
		}
		if dbPassword == "" {
			missingVars = append(missingVars, "DB_PASSWORD")
		}
		if dbName == "" {
			missingVars = append(missingVars, "DB_NAME")
		}
		if dbPort == "" {
			missingVars = append(missingVars, "DB_PORT")
		}

		if len(missingVars) > 0 {
			utils.LogError(fmt.Errorf("missing environment variables: %v", missingVars), "DSN may failed due to missing variable")
		}

		// Create the DSN (Data Source Name)
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbHost, dbUser, dbPassword, dbName, dbPort)
	}

	// Determine the logger configuration based on the environment
	var config *gorm.Config
	if ginMode := os.Getenv("GIN_MODE"); ginMode == gin.DebugMode {
		config = &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold:             time.Second,
					LogLevel:                  logger.Info,
					IgnoreRecordNotFoundError: true,
					Colorful:                  false,
				},
			),
		}
	} else {
		config = &gorm.Config{}
	}

	// Connect to the database
	database, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		utils.LogError(err, "Failed to connect to database")
		return
	}

	DB = database
}
