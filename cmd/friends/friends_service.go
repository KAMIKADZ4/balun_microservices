package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5002"
	}
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}\n",
		},
	))
	e.Use(middleware.Recover())

	// Routes
	e.GET("/health-check", okHandler)
	e.GET("/ready-check", okHandler)

	// Start server
	log.Println("Lets go!")
	if err := e.Start(":" + port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("failed to start server", "error", err)
	}
}

// Handler
func okHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "OK",
		"service": "friends",
	})
}
