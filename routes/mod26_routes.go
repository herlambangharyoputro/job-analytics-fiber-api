package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/herlambangharyoputro/job-analytics-fiber-api/controllers/mod26_data_quality"
)

func SetupMod26Routes(api fiber.Router) {
    quality := api.Group("/quality")

    // Quality metrics endpoints
    quality.Get("/metrics", mod26_data_quality.GetQualityMetrics)
    quality.Get("/metrics/latest", mod26_data_quality.GetLatestMetrics)
    quality.Get("/metrics/summary", mod26_data_quality.GetQualitySummary)

    // Quality issues endpoints
    quality.Get("/issues", mod26_data_quality.GetQualityIssues)
    quality.Get("/issues/:id", mod26_data_quality.GetQualityIssueByID)
    quality.Post("/issues/:id/resolve", mod26_data_quality.ResolveQualityIssue)

    // Field quality endpoints
    quality.Get("/fields", mod26_data_quality.GetFieldQualityChecks)

    // Validation rules endpoints
    quality.Get("/rules", mod26_data_quality.GetValidationRules)
    quality.Post("/rules", mod26_data_quality.CreateValidationRule)
}
