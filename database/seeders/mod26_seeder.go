package seeders

import (
"log"
"time"

"github.com/herlambangharyoputro/job-analytics-fiber-api/config"
"github.com/herlambangharyoputro/job-analytics-fiber-api/models/mod26_data_quality"
)

// SeedMod26DataQuality seeds realistic data quality metrics based on actual data issues
func SeedMod26DataQuality() {
log.Println(" Seeding Module 26: Data Quality Monitoring...")

// Clear existing data (optional - comment out if you want to keep)
config.DB.Exec("DELETE FROM mod26_field_quality_checks")
config.DB.Exec("DELETE FROM mod26_quality_issues")
config.DB.Exec("DELETE FROM mod26_quality_metrics")
config.DB.Exec("DELETE FROM mod26_validation_rules")

// 1. Seed Validation Rules
seedValidationRules()

// 2. Seed Quality Metrics (last 7 days)
seedQualityMetrics()

// 3. Seed Quality Issues
seedQualityIssues()

// 4. Seed Field Quality Checks
seedFieldQualityChecks()

log.Println("✅ Module 26 seeding completed!")
}

func seedValidationRules() {
rules := []mod26_data_quality.DataValidationRule{
{
RuleName:    "job_title_required",
TargetTable: "jobs",
FieldName:   "judul",
RuleType:    "required",
RuleValue:   "not_null",
IsActive:    true,
Description: "Job title must not be empty",
},
{
RuleName:    "company_name_required",
TargetTable: "jobs",
FieldName:   "perusahaan",
RuleType:    "required",
RuleValue:   "not_null",
IsActive:    true,
Description: "Company name must not be empty",
},
{
RuleName:    "location_required",
TargetTable: "jobs",
FieldName:   "lokasi",
RuleType:    "required",
RuleValue:   "not_null",
IsActive:    true,
Description: "Location must not be empty",
},
{
RuleName:    "salary_format_validation",
TargetTable: "jobs",
FieldName:   "gaji",
RuleType:    "format",
RuleValue:   `{"pattern": "^Rp\\s*\\d+(\\.\\d+)?\\s*(Juta|Ribu)?"}`,
IsActive:    true,
Description: "Salary must follow Indonesian format (Rp X Juta)",
},
{
RuleName:    "date_format_validation",
TargetTable: "jobs",
FieldName:   "tanggal_posting",
RuleType:    "format",
RuleValue:   `{"pattern": "^\\d{2}-\\d{2}-\\d{2}\\s\\d{2}:\\d{2}$"}`,
IsActive:    true,
Description: "Date must follow format: DD-MM-YY HH:MM",
},
{
RuleName:    "benefit_no_error",
TargetTable: "jobs",
FieldName:   "benefit",
RuleType:    "custom",
RuleValue:   `{"exclude": "#NAME?"}`,
IsActive:    true,
Description: "Benefit field must not contain #NAME? error",
},
{
RuleName:    "employee_count_validation",
TargetTable: "companies",
FieldName:   "jumlah_karyawan",
RuleType:    "custom",
RuleValue:   `{"exclude": "undefined"}`,
IsActive:    true,
Description: "Employee count must not be 'undefined'",
},
{
RuleName:    "skills_required",
TargetTable: "jobs",
FieldName:   "keahlian",
RuleType:    "required",
RuleValue:   "not_empty",
IsActive:    true,
Description: "Skills field should not be empty",
},
}

for _, rule := range rules {
if err := config.DB.Create(&rule).Error; err != nil {
log.Printf("  Failed to seed rule %s: %v", rule.RuleName, err)
}
}

log.Printf(" Seeded %d validation rules", len(rules))
}

func seedQualityMetrics() {
// Generate metrics for last 7 days with realistic scores
metricsData := []struct {
daysAgo           int
completeness      float64
accuracy          float64
consistency       float64
timeliness        float64
totalRecords      int
validRecords      int
invalidRecords    int
}{
{0, 68.5, 83.2, 79.8, 95.1, 1250, 897, 353},  // Today - worst
{1, 69.3, 84.1, 80.5, 95.3, 1230, 905, 325},
{2, 70.8, 85.7, 81.2, 95.8, 1200, 920, 280},
{3, 72.1, 86.4, 82.1, 96.1, 1180, 935, 245},
{4, 73.5, 87.2, 83.0, 96.5, 1150, 948, 202},
{5, 74.8, 88.1, 84.2, 96.8, 1120, 962, 158},
{6, 76.2, 89.0, 85.5, 97.2, 1100, 978, 122},  // 6 days ago - best
}

for _, data := range metricsData {
metricDate := time.Now().AddDate(0, 0, -data.daysAgo)
overall := (data.completeness + data.accuracy + data.consistency + data.timeliness) / 4.0

metric := mod26_data_quality.DataQualityMetric{
MetricDate:        metricDate,
TotalRecords:      data.totalRecords,
ValidRecords:      data.validRecords,
InvalidRecords:    data.invalidRecords,
CompletenessScore: data.completeness,
AccuracyScore:     data.accuracy,
ConsistencyScore:  data.consistency,
TimelinessScore:   data.timeliness,
OverallScore:      overall,
}

if err := config.DB.Create(&metric).Error; err != nil {
log.Printf("⚠️  Failed to seed metric for %s: %v", metricDate.Format("2006-01-02"), err)
}
}

log.Printf(" Seeded %d quality metrics (last 7 days)", len(metricsData))
}

func seedQualityIssues() {
now := time.Now()

issues := []mod26_data_quality.DataQualityIssue{
{
IssueType:   "missing",
Severity:    "medium",
TargetTable: "jobs",
FieldName:   "gaji",
RecordID:    "job_12345",
Description: "Salary field is empty - affects 35% of records",
DetectedAt:  now.Add(-2 * time.Hour),
Status:      "open",
},
{
IssueType:   "invalid",
Severity:    "high",
TargetTable: "jobs",
FieldName:   "benefit",
RecordID:    "job_23456",
Description: "Benefit field contains #NAME? error from Excel",
DetectedAt:  now.Add(-3 * time.Hour),
Status:      "open",
},
{
IssueType:   "invalid",
Severity:    "low",
TargetTable: "companies",
FieldName:   "jumlah_karyawan",
RecordID:    "company_5678",
Description: "Employee count shows 'undefined Karyawan'",
DetectedAt:  now.Add(-5 * time.Hour),
Status:      "open",
},
{
IssueType:   "missing",
Severity:    "critical",
TargetTable: "jobs",
FieldName:   "all_fields",
RecordID:    "job_34567",
Description: "Empty row detected at end of CSV file",
DetectedAt:  now.Add(-1 * time.Hour),
Status:      "open",
},
{
IssueType:   "missing",
Severity:    "medium",
TargetTable: "jobs",
FieldName:   "tanggung_jawab",
RecordID:    "job_45678",
Description: "Job responsibilities field is empty",
DetectedAt:  now.Add(-4 * time.Hour),
Status:      "open",
},
{
IssueType:   "missing",
Severity:    "low",
TargetTable: "jobs",
FieldName:   "benefit",
RecordID:    "job_56789",
Description: "Benefits information not provided - common issue (70% missing)",
DetectedAt:  now.Add(-6 * time.Hour),
Status:      "ignored",
},
{
IssueType:   "inconsistent",
Severity:    "medium",
TargetTable: "jobs",
FieldName:   "gaji",
RecordID:    "job_67890",
Description: "Inconsistent salary format: some use 'Juta', others use numeric only",
DetectedAt:  now.Add(-24 * time.Hour),
Status:      "resolved",
},
{
IssueType:   "duplicate",
Severity:    "high",
TargetTable: "jobs",
FieldName:   "url",
RecordID:    "job_78901",
Description: "Duplicate job posting detected (same URL scraped twice)",
DetectedAt:  now.Add(-12 * time.Hour),
Status:      "open",
},
}

// Set resolved_at for resolved issues
resolvedTime := now.Add(-12 * time.Hour)
issues[6].ResolvedAt = &resolvedTime

for _, issue := range issues {
if err := config.DB.Create(&issue).Error; err != nil {
log.Printf("  Failed to seed issue: %v", err)
}
}

log.Printf(" Seeded %d quality issues", len(issues))
}

func seedFieldQualityChecks() {
// Get the latest metric ID
var latestMetric mod26_data_quality.DataQualityMetric
if err := config.DB.Order("metric_date DESC").First(&latestMetric).Error; err != nil {
log.Printf("  No metrics found, skipping field quality checks")
return
}

// Field quality data based on DATA_SCHEMA.md
fieldChecks := []struct {
fieldName        string
totalValues      int
nullCount        int
emptyCount       int
invalidCount     int
validCount       int
completenessRate float64
}{
{"judul", 1250, 0, 0, 0, 1250, 100.0},           // 100% complete
{"perusahaan", 1250, 0, 0, 0, 1250, 100.0},      // 100% complete
{"lokasi", 1250, 0, 0, 0, 1250, 100.0},          // 100% complete
{"gaji", 1250, 438, 0, 0, 812, 65.0},            // 65% complete (35% missing)
{"tanggung_jawab", 1250, 200, 50, 0, 1000, 80.0}, // 80% complete
{"kualifikasi", 1250, 150, 38, 0, 1062, 85.0},   // 85% complete
{"keahlian", 1250, 250, 63, 0, 937, 75.0},       // 75% complete
{"benefit", 1250, 750, 125, 50, 325, 26.0},      // 26% complete (70% missing + errors)
{"industri", 1250, 0, 0, 0, 1250, 100.0},        // 100% complete
{"pendidikan", 1250, 0, 0, 0, 1250, 100.0},      // 100% complete
}

for _, data := range fieldChecks {
check := mod26_data_quality.FieldQualityCheck{
MetricID:         latestMetric.ID,
FieldName:        data.fieldName,
TotalValues:      data.totalValues,
NullCount:        data.nullCount,
EmptyCount:       data.emptyCount,
InvalidCount:     data.invalidCount,
ValidCount:       data.validCount,
CompletenessRate: data.completenessRate,
}

if err := config.DB.Create(&check).Error; err != nil {
log.Printf("  Failed to seed field check for %s: %v", data.fieldName, err)
}
}

log.Printf("✓ Seeded %d field quality checks", len(fieldChecks))
}