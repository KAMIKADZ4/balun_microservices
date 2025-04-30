package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func okHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"service": "users",
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5003"
	}

	r := gin.Default()
	r.GET("/health-check", okHandler)
	r.GET("/ready-check", okHandler)

	log.Println("Lets go!")
	r.Run("localhost:" + port)
}
