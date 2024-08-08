package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/sirupsen/logrus"
    "github.com/3milly4ever/lambda-parser-landstar/internal/model"
    "github.com/3milly4ever/lambda-parser-landstar/internal/parser"
)

func SetupRoutes(app *fiber.App) {
    app.Post("/parse-email", parseEmail)
}

func parseEmail(c *fiber.Ctx) error {
    logrus.Info("Received request")
    var req model.Request
    if err := c.BodyParser(&req); err != nil {
        logrus.Error("Error parsing request body: ", err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    response, err := parser.ParseEmail(req)
    if err != nil {
        logrus.Error("Error parsing email: ", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error processing email",
        })
    }

    logrus.Info("Successfully parsed email")

    return c.JSON(response)
}
