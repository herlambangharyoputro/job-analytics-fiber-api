package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/indojobmarket/job-analytics-fiber-api/config"
)

func SetupRoutes(app *fiber.App) {
    // API v1 group
    api := app.Group(config.AppConfig.APIPrefix)

    // Module 26: Data Quality routes
    SetupMod26Routes(api)

    // Welcome route
    api.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "message": "Welcome to IndoJobMarket Fiber API",
            "version": config.AppConfig.APIVersion,
            "endpoints": fiber.Map{
                "health":        "/health",
                "data_quality":  config.AppConfig.APIPrefix + "/quality",
            },
        })
    })
}
