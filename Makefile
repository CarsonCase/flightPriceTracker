#!Makefile
include .env


# Set your PostgreSQL database connection URL
DB_URL := ${DB_URL}

# Set the path to your migration files
MIGRATION_PATH := sql/migrations

# Default target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  migrate-up      : Apply migrations"
	@echo "  migrate-down    : Rollback migrations"
	@echo "  migrate-status  : View migration status"
	@echo "  help            : Show this help message"

# Start DB in docker-compose
dbup:
	docker-compose up -d

# Close DB in docker-compose
dbdown:
	docker-compose down

# Apply migrations
.PHONY: migrate-up
up:
	migrate -database $(DB_URL) -path cmd/server/$(MIGRATION_PATH) up

# Rollback migrations
.PHONY: migrate-down
down:
	migrate -database $(DB_URL) -path cmd/server/$(MIGRATION_PATH) down

# View migration status
.PHONY: migrate-status
status:
	migrate -database $(DB_URL) -path cmd/server/$(MIGRATION_PATH) status

server:
	go build ./cmd/server

goFlight:
	go build ./cmd/goFlight