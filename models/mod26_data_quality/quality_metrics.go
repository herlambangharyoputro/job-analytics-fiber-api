package mod26_data_quality

import "time"

// DataQualityMetric represents overall quality metrics
type DataQualityMetric struct {
ID               uint      `gorm:"primaryKey"`
MetricDate       time.Time `gorm:"index;not null"`
TotalRecords     int       `gorm:"not null"`
ValidRecords     int       `gorm:"not null"`
InvalidRecords   int       `gorm:"not null"`
CompletenessScore float64  `gorm:"type:decimal(5,2)"`
AccuracyScore    float64   `gorm:"type:decimal(5,2)"`
ConsistencyScore float64   `gorm:"type:decimal(5,2)"`
TimelinessScore  float64   `gorm:"type:decimal(5,2)"`
OverallScore     float64   `gorm:"type:decimal(5,2);index"`
CreatedAt        time.Time
UpdatedAt        time.Time
}

// TableName overrides the table name
func (DataQualityMetric) TableName() string {
return "mod26_quality_metrics"
}

// FieldQualityCheck represents quality check for specific fields
type FieldQualityCheck struct {
ID              uint      `gorm:"primaryKey"`
MetricID        uint      `gorm:"index;not null"`
FieldName       string    `gorm:"size:100;not null;index"`
TotalValues     int       `gorm:"not null"`
NullCount       int       `gorm:"not null"`
EmptyCount      int       `gorm:"not null"`
InvalidCount    int       `gorm:"not null"`
ValidCount      int       `gorm:"not null"`
CompletenessRate float64  `gorm:"type:decimal(5,2)"`
CreatedAt       time.Time
}

// TableName overrides the table name
func (FieldQualityCheck) TableName() string {
return "mod26_field_quality_checks"
}

// DataQualityIssue represents specific quality issues found
type DataQualityIssue struct {
ID          uint       `gorm:"primaryKey"`
IssueType   string     `gorm:"size:50;not null;index"` // missing, invalid, duplicate, outlier
Severity    string     `gorm:"size:20;not null;index"` // low, medium, high, critical
TargetTable string     `gorm:"column:table_name;size:100;not null"` // Renamed from TableName to avoid conflict
FieldName   string     `gorm:"size:100"`
RecordID    string     `gorm:"size:100"`
Description string     `gorm:"type:text"`
DetectedAt  time.Time  `gorm:"index;not null"`
ResolvedAt  *time.Time
Status      string     `gorm:"size:20;default:'open';index"` // open, resolved, ignored
CreatedAt   time.Time
UpdatedAt   time.Time
}

// TableName overrides the table name for GORM
func (DataQualityIssue) TableName() string {
return "mod26_quality_issues"
}