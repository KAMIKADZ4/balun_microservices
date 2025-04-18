package main 

import (
  "os"
  "log"
  "github.com/gofiber/fiber/v3"
  "github.com/gofiber/fiber/v3/middleware/logger"
)

func okHandler(context fiber.Ctx) error {
  return context.JSON(fiber.Map{
    "status": "OK",
    "service": "auth",
  })
}


func main() {
  port := os.Getenv("PORT")
  if port == "" {
    port = "5000"
  }

  app := fiber.New()

  app.Use(logger.New(logger.Config{
    Format: "${time} - [${method} ${status}] ${latency} ${url}\n",
  }))

  app.Get("/health-check", okHandler )
  app.Get("/ready-check", okHandler)

  log.Println("Lets go!")
  log.Fatal(app.Listen(":" + port))
}
