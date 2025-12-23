package database

import (
    "log"
    "github.com/herlambangharyoputro/job-analytics-fiber-api/config"
    "github.com/herlambangharyoputro/job-analytics-fiber-api/models/mod26_data_quality"
)

func RunMigrations() {
    log.Println("Running database migrations...")

    // Auto-migrate Module 26 tables
    err := config.DB.AutoMigrate(
        &mod26_data_quality.DataQualityMetric{},
        &mod26_data_quality.FieldQualityCheck{},
        &mod26_data_quality.DataQualityIssue{},
        &mod26_data_quality.DataValidationRule{},
    )

    if err != nil {
        log.Fatal("Migration failed:", err)
    }

    log.Println("Migrations completed successfully")
}
