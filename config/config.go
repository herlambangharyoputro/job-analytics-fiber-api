package config

import (
    "log"
    "os"
    "strconv"
    "github.com/joho/godotenv"
)

type Config struct {
    AppName    string
    AppEnv     string
    AppPort    string
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    APIVersion string
    APIPrefix  string
}

var AppConfig *Config

func LoadConfig() {
    // Load .env file
    env := os.Getenv("APP_ENV")
    if env == "" {
        env = "development"
    }
    
    envFile := ".env." + env
    if err := godotenv.Load(envFile); err != nil {
        log.Printf("Warning: .env file not found, using environment variables")
    }

    AppConfig = &Config{
        AppName:    getEnv("APP_NAME", "IndoJobMarket Fiber API"),
        AppEnv:     getEnv("APP_ENV", "development"),
        AppPort:    getEnv("APP_PORT", "3000"),
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "3306"),
        DBUser:     getEnv("DB_USER", "root"),
        DBPassword: getEnv("DB_PASSWORD", ""),
        DBName:     getEnv("DB_NAME", "indojobmarket_dev"),
        APIVersion: getEnv("API_VERSION", "v1"),
        APIPrefix:  getEnv("API_PREFIX", "/api/v1"),
    }

    log.Printf("Configuration loaded: %s environment", AppConfig.AppEnv)
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
    valueStr := getEnv(key, "")
    if value, err := strconv.Atoi(valueStr); err == nil {
        return value
    }
    return defaultValue
}
