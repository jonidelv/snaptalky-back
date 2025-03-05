package utils

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/rollbar/rollbar-go"
)

// InitRollbar initializes the Rollbar client with the provided configuration
func InitRollbar() {
	token := os.Getenv("ROLLBAR_TOKEN")
	isProduction := os.Getenv("ENV") == "production"

	if isProduction {
		rollbar.SetToken(token)
		rollbar.SetEnvironment("production")
		rollbar.SetServerRoot("github.com/jonidelv/snaptalky-back")
	}
}

// LogError logs an error to Rollbar in production, console in development
func LogError(err error, message string, extras ...map[string]interface{}) {
	if err == nil {
		return
	}

	// Get the caller's file and line number
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}

	// Create a custom error message with file and line information
	customMsg := fmt.Sprintf("%s - %s:%d", message, file, line)

	isProduction := os.Getenv("ENV") == "production"
	if isProduction {
		// If there are additional context/extras provided
		if len(extras) > 0 {
			// Merge the custom message into the extras map
			extrasWithMsg := extras[0]
			extrasWithMsg["message"] = customMsg
			rollbar.ErrorWithExtras(rollbar.ERR, err, extrasWithMsg)
		} else {
			// Pass both error and custom message to Rollbar
			rollbar.ErrorWithExtras(rollbar.ERR, err, Object{
				"message": customMsg,
			})
		}
	} else {
		// In development, just log to console
		log.Printf("%s: %v", customMsg, err)
	}
}
