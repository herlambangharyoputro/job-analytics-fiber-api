package mod26_data_quality

import (
"strconv"

"github.com/gofiber/fiber/v2"
"github.com/herlambangharyoputro/job-analytics-fiber-api/models/mod26_data_quality"
"github.com/herlambangharyoputro/job-analytics-fiber-api/services/mod26_data_quality"
"github.com/herlambangharyoputro/job-analytics-fiber-api/utils"
)

var qualityService = mod26_data_quality.NewQualityService()

// GetQualityMetrics retrieves quality metrics with pagination
func GetQualityMetrics(c *fiber.Ctx) error {
limit := c.QueryInt("limit", 10)
if limit > 100 {
limit = 100
}

metrics, err := qualityService.GetAllMetrics(limit)
if err != nil {
return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve metrics", err)
}

return utils.SuccessResponse(c, "Quality metrics retrieved successfully", fiber.Map{
"metrics": metrics,
"count":   len(metrics),
})
}

// GetLatestMetrics retrieves the latest quality metrics
func GetLatestMetrics(c *fiber.Ctx) error {
metric, err := qualityService.GetLatestMetric()
if err != nil {
return utils.ErrorResponse(c, fiber.StatusNotFound, "No metrics found", err)
}

return utils.SuccessResponse(c, "Latest metrics retrieved successfully", metric)
}

// GetQualitySummary retrieves quality summary with aggregated data
func GetQualitySummary(c *fiber.Ctx) error {
summary, err := qualityService.GetQualitySummary()
if err != nil {
return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to generate summary", err)
}

return utils.SuccessResponse(c, "Quality summary generated successfully", summary)
}

// GetQualityIssues retrieves all quality issues with filters
func GetQualityIssues(c *fiber.Ctx) error {
status := c.Query("status", "")      // open, resolved, ignored
severity := c.Query("severity", "")  // low, medium, high, critical
limit := c.QueryInt("limit", 50)

if limit > 200 {
limit = 200
}

issues, err := qualityService.GetQualityIssues(status, severity, limit)
if err != nil {
return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve issues", err)
}

return utils.SuccessResponse(c, "Quality issues retrieved successfully", fiber.Map{
"issues": issues,
"count":  len(issues),
"filters": fiber.Map{
"status":   status,
"severity": severity,
},
})
}

// GetQualityIssueByID retrieves a specific quality issue
func GetQualityIssueByID(c *fiber.Ctx) error {
id, err := strconv.ParseUint(c.Params("id"), 10, 32)
if err != nil {
return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid issue ID", err)
}

issue, err := qualityService.GetIssueByID(uint(id))
if err != nil {
return utils.ErrorResponse(c, fiber.StatusNotFound, "Issue not found", err)
}

return utils.SuccessResponse(c, "Quality issue retrieved successfully", issue)
}

// ResolveQualityIssue marks an issue as resolved
func ResolveQualityIssue(c *fiber.Ctx) error {
id, err := strconv.ParseUint(c.Params("id"), 10, 32)
if err != nil {
return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid issue ID", err)
}

if err := qualityService.ResolveIssue(uint(id)); err != nil {
return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to resolve issue", err)
}

return utils.SuccessResponse(c, "Quality issue resolved successfully", fiber.Map{
"id":     id,
"status": "resolved",
})
}

// GetFieldQualityChecks retrieves field-level quality checks
func GetFieldQualityChecks(c *fiber.Ctx) error {
metricIDStr := c.Query("metric_id", "")
var checks []interface{}

if metricIDStr != "" {
metricID, parseErr := strconv.ParseUint(metricIDStr, 10, 32)
if parseErr != nil {
return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid metric_id", parseErr)
}

result, getErr := qualityService.GetFieldQualityChecks(uint(metricID))
if getErr != nil {
return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve field checks", getErr)
}

// Convert to interface slice
for _, check := range result {
checks = append(checks, check)
}
} else {
limit := c.QueryInt("limit", 20)
result, getErr := qualityService.GetAllFieldQualityChecks(limit)
if getErr != nil {
return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve field checks", getErr)
}

// Convert to interface slice
for _, check := range result {
checks = append(checks, check)
}
}

return utils.SuccessResponse(c, "Field quality checks retrieved successfully", fiber.Map{
"fields": checks,
"count":  len(checks),
})
}

// GetValidationRules retrieves all validation rules
func GetValidationRules(c *fiber.Ctx) error {
activeStr := c.Query("active", "")
var isActive *bool

if activeStr != "" {
activeBool := activeStr == "true" || activeStr == "1"
isActive = &activeBool
}

rules, err := qualityService.GetValidationRules(isActive)
if err != nil {
return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve rules", err)
}

return utils.SuccessResponse(c, "Validation rules retrieved successfully", fiber.Map{
"rules": rules,
"count": len(rules),
})
}

// CreateValidationRule creates a new validation rule
func CreateValidationRule(c *fiber.Ctx) error {
var rule models.DataValidationRule

if err := c.BodyParser(&rule); err != nil {
return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err)
}

if err := qualityService.CreateValidationRule(&rule); err != nil {
return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create rule", err)
}

return utils.CreatedResponse(c, "Validation rule created successfully", rule)
}