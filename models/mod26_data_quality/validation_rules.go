package mod26_data_quality

import "time"

// DataValidationRule represents validation rules
type DataValidationRule struct {
ID          uint      `gorm:"primaryKey"`
RuleName    string    `gorm:"size:100;not null;uniqueIndex"`
TableName   string    `gorm:"size:100;not null;index"`
FieldName   string    `gorm:"size:100;not null"`
RuleType    string    `gorm:"size:50;not null"` // required, format, range, custom
RuleValue   string    `gorm:"type:text"` // JSON or string representing rule
IsActive    bool      `gorm:"default:true;index"`
Description string    `gorm:"type:text"`
CreatedAt   time.Time
UpdatedAt   time.Time
}

// TableName overrides the table name for GORM
func (DataValidationRule) TableName() string {
return "mod26_validation_rules"
}