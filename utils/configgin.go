package utils

import (
	"github.com/gin-gonic/gin"
)

// InitGin initializes and configures a Gin engine based on the environment
func InitGin(isProduction bool) *gin.Engine {
	var r *gin.Engine
	if isProduction {
		// In production, use a minimal logger that doesn't log request bodies or sensitive data
		r = gin.New()
		r.Use(gin.Recovery())
		// Custom minimal logger that doesn't log request bodies
		r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
			SkipPaths: []string{"/health"}, // Skip logging health check endpoints
		}))
	} else {
		// In development, use the default logger
		r = gin.Default()
	}

	r.RemoveExtraSlash = true

	return r
}
