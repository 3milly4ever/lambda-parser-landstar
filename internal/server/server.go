package server

import (
    "github.com/gofiber/fiber/v2"
    "github.com/yourusername/project/internal/log"
    "github.com/yourusername/project/internal/routes"
)

func StartServer() {
    app := fiber.New()
    routes.SetupRoutes(app)

    log.InitLogger()
    logrus.Info("Starting server on port 3000")
    if err := app.Listen(":3000"); err != nil {
        logrus.Fatalf("Error starting server: %v", err)
    }
}
