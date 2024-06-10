package routes

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func ProcessResponse(c *gin.Context) {
  text := c.PostForm("text")
  file, _, err := c.Request.FormFile("file")
  if err != nil && err != http.ErrMissingFile {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Placeholder for actual processing logic
  response := "Processed response: " + text
  if file != nil {
    response += " (with image)"
  }

  c.JSON(http.StatusOK, gin.H{"response": response})
}
