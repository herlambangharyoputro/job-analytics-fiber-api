# IndoJobMarket Fiber API

**Module 26: Data Quality Monitoring Dashboard - Backend**

## Tech Stack
- **Framework**: Fiber (Go)
- **Database**: MySQL (local) / PostgreSQL (production)
- **ORM**: GORM
- **Deployment**: Railway

## Prerequisites
- Go 1.21+
- MySQL 8.0+

## Installation

\\\powershell
# Install dependencies
go mod download

# Copy environment file
Copy-Item .env.example .env.development

# Update .env.development with your MySQL credentials
\\\

## Database Setup

\\\sql
-- Create database
CREATE DATABASE indojobmarket_dev CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
\\\

## Running the Application

\\\powershell
# Development mode
go run main.go
\\\

Server will start on http://localhost:8003

## API Endpoints

### Health Check
- GET /health

### Module 26: Data Quality
- GET /api/v1/quality/metrics
- GET /api/v1/quality/metrics/latest
- GET /api/v1/quality/issues
- GET /api/v1/quality/fields
- GET /api/v1/quality/rules

## Project Structure

\\\
job-analytics-fiber-api/
 config/              # Configuration files
 controllers/         # Request handlers
    mod26_data_quality/
 database/            # Database migrations
 models/              # Database models
    mod26_data_quality/
 routes/              # API routes
 services/            # Business logic
    mod26_data_quality/
 utils/               # Utility functions
 main.go              # Application entry point
 .env.development     # Environment variables
\\\

## Development Roadmap

- [x] Phase 1: Backend Setup (Week 1) 
- [ ] Phase 2: Core Backend Development (Week 1-2)
- [ ] Phase 3: Frontend Setup (Week 2-3)
- [ ] Phase 4: Frontend Development (Week 3-4)
- [ ] Phase 5: Integration & Testing (Week 4)
- [ ] Phase 6: Deployment (Week 5)
