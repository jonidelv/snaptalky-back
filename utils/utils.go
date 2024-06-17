package utils

import (
	"log"
)

// LogError logs the error with a consistent format
func LogError(err error, message string) {
	if err != nil {
		log.Printf("%s: %v", message, err)
	}
}
