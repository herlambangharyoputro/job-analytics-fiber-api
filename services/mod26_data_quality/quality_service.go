package mod26_data_quality

import (
"time"

"github.com/herlambangharyoputro/job-analytics-fiber-api/config"
"github.com/herlambangharyoputro/job-analytics-fiber-api/models/mod26_data_quality"
)

// QualityService handles business logic for data quality
type QualityService struct{}

// NewQualityService creates a new quality service instance
func NewQualityService() *QualityService {
return &QualityService{}
}

// GetAllMetrics retrieves all quality metrics
func (s *QualityService) GetAllMetrics(limit int) ([]mod26_data_quality.DataQualityMetric, error) {
var metrics []mod26_data_quality.DataQualityMetric
result := config.DB.Order("metric_date DESC").Limit(limit).Find(&metrics)
return metrics, result.Error
}

// GetLatestMetric retrieves the most recent quality metric
func (s *QualityService) GetLatestMetric() (*mod26_data_quality.DataQualityMetric, error) {
var metric mod26_data_quality.DataQualityMetric
result := config.DB.Order("metric_date DESC").First(&metric)
if result.Error != nil {
return nil, result.Error
}
return &metric, nil
}

// CreateMetric creates a new quality metric record
func (s *QualityService) CreateMetric(metric *mod26_data_quality.DataQualityMetric) error {
return config.DB.Create(metric).Error
}

// GetQualityIssues retrieves quality issues with filters
func (s *QualityService) GetQualityIssues(status string, severity string, limit int) ([]mod26_data_quality.DataQualityIssue, error) {
var issues []mod26_data_quality.DataQualityIssue
query := config.DB.Order("detected_at DESC")

if status != "" {
query = query.Where("status = ?", status)
}

if severity != "" {
query = query.Where("severity = ?", severity)
}

result := query.Limit(limit).Find(&issues)
return issues, result.Error
}

// GetIssueByID retrieves a specific quality issue
func (s *QualityService) GetIssueByID(id uint) (*mod26_data_quality.DataQualityIssue, error) {
var issue mod26_data_quality.DataQualityIssue
result := config.DB.First(&issue, id)
if result.Error != nil {
return nil, result.Error
}
return &issue, nil
}

// ResolveIssue marks an issue as resolved
func (s *QualityService) ResolveIssue(id uint) error {
now := time.Now()
return config.DB.Model(&mod26_data_quality.DataQualityIssue{}).
Where("id = ?", id).
Updates(map[string]interface{}{
"status":      "resolved",
"resolved_at": now,
}).Error
}

// CreateIssue creates a new quality issue
func (s *QualityService) CreateIssue(issue *mod26_data_quality.DataQualityIssue) error {
return config.DB.Create(issue).Error
}

// GetFieldQualityChecks retrieves field quality checks for a specific metric
func (s *QualityService) GetFieldQualityChecks(metricID uint) ([]mod26_data_quality.FieldQualityCheck, error) {
var checks []mod26_data_quality.FieldQualityCheck
result := config.DB.Where("metric_id = ?", metricID).Find(&checks)
return checks, result.Error
}

// GetAllFieldQualityChecks retrieves recent field quality checks
func (s *QualityService) GetAllFieldQualityChecks(limit int) ([]mod26_data_quality.FieldQualityCheck, error) {
var checks []mod26_data_quality.FieldQualityCheck
result := config.DB.Order("created_at DESC").Limit(limit).Find(&checks)
return checks, result.Error
}

// CreateFieldQualityCheck creates a new field quality check
func (s *QualityService) CreateFieldQualityCheck(check *mod26_data_quality.FieldQualityCheck) error {
return config.DB.Create(check).Error
}

// GetValidationRules retrieves all active validation rules
func (s *QualityService) GetValidationRules(isActive *bool) ([]mod26_data_quality.DataValidationRule, error) {
var rules []mod26_data_quality.DataValidationRule
query := config.DB.Order("created_at DESC")

if isActive != nil {
query = query.Where("is_active = ?", *isActive)
}

result := query.Find(&rules)
return rules, result.Error
}

// CreateValidationRule creates a new validation rule
func (s *QualityService) CreateValidationRule(rule *mod26_data_quality.DataValidationRule) error {
return config.DB.Create(rule).Error
}

// UpdateValidationRule updates an existing validation rule
func (s *QualityService) UpdateValidationRule(id uint, updates map[string]interface{}) error {
return config.DB.Model(&mod26_data_quality.DataValidationRule{}).
Where("id = ?", id).
Updates(updates).Error
}

// DeleteValidationRule soft deletes a validation rule (marks as inactive)
func (s *QualityService) DeleteValidationRule(id uint) error {
return config.DB.Model(&mod26_data_quality.DataValidationRule{}).
Where("id = ?", id).
Update("is_active", false).Error
}

// GetQualitySummary generates a summary of quality metrics
func (s *QualityService) GetQualitySummary() (map[string]interface{}, error) {
// Get latest metric
latestMetric, err := s.GetLatestMetric()
if err != nil {
return nil, err
}

// Count open issues by severity
var criticalCount, highCount, mediumCount, lowCount int64
config.DB.Model(&mod26_data_quality.DataQualityIssue{}).
Where("status = ? AND severity = ?", "open", "critical").Count(&criticalCount)
config.DB.Model(&mod26_data_quality.DataQualityIssue{}).
Where("status = ? AND severity = ?", "open", "high").Count(&highCount)
config.DB.Model(&mod26_data_quality.DataQualityIssue{}).
Where("status = ? AND severity = ?", "open", "medium").Count(&mediumCount)
config.DB.Model(&mod26_data_quality.DataQualityIssue{}).
Where("status = ? AND severity = ?", "open", "low").Count(&lowCount)

// Total active rules
var activeRulesCount int64
config.DB.Model(&mod26_data_quality.DataValidationRule{}).
Where("is_active = ?", true).Count(&activeRulesCount)

summary := map[string]interface{}{
"latest_metric": latestMetric,
"open_issues": map[string]interface{}{
"critical": criticalCount,
"high":     highCount,
"medium":   mediumCount,
"low":      lowCount,
"total":    criticalCount + highCount + mediumCount + lowCount,
},
"active_rules": activeRulesCount,
"timestamp":    time.Now(),
}

return summary, nil
}