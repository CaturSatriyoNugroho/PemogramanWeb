package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/ping", func(c *gin.Context) {
		dbStatus := "Disconnected"
		if os.Getenv("DB_HOST") != "" {
			dbStatus = "Connected to " + os.Getenv("DB_HOST")
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"status":  "Backend is running",
			"db":      dbStatus,
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting server on port %s...\n", port)
	r.Run(":" + port)
}
