package main

import (
"log"
"fmt"

"github.com/gofiber/fiber/v2"
"github.com/gofiber/fiber/v2/middleware/cors"
"github.com/gofiber/fiber/v2/middleware/logger"
"github.com/gofiber/fiber/v2/middleware/recover"

"github.com/herlambangharyoputro/job-analytics-fiber-api/config"
"github.com/herlambangharyoputro/job-analytics-fiber-api/database"
"github.com/herlambangharyoputro/job-analytics-fiber-api/routes"
)

func main() {
// Load configuration
config.LoadConfig()

// Initialize Fiber app
app := fiber.New(fiber.Config{
AppName: config.AppConfig.AppName,
ErrorHandler: func(c *fiber.Ctx, err error) error {
code := fiber.StatusInternalServerError
if e, ok := err.(*fiber.Error); ok {
code = e.Code
}
return c.Status(code).JSON(fiber.Map{
"success": false,
"message": "An error occurred",
"error":   err.Error(),
})
},
})

// Middleware
app.Use(recover.New())
app.Use(logger.New())
app.Use(cors.New(cors.Config{
AllowOrigins: "*",
AllowHeaders: "Origin, Content-Type, Accept, Authorization",
AllowMethods: "GET, POST, PUT, DELETE, PATCH, OPTIONS",
}))

// Connect to database
config.ConnectDatabase()

// Run migrations
database.RunMigrations()

// Run seeders (comment out after first run if you want to keep data)
database.RunSeeders()

// Setup routes
routes.SetupRoutes(app)

// Health check
app.Get("/health", func(c *fiber.Ctx) error {
return c.JSON(fiber.Map{
"status":  "ok",
"message": "IndoJobMarket Fiber API is running",
"version": config.AppConfig.APIVersion,
})
})

// Start server
port := config.AppConfig.AppPort
log.Printf(" Server starting on port %s", port)
log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}