run: swagger-gen
	go run cmd/main.go

# Database configuration from environment variables
PG_HOST ?= localhost
PG_PORT ?= 5437
PG_USERNAME ?= postgres
PG_PASS ?= nodirbek
PG_DB ?= postgres
DB_URL=postgres://$(PG_USERNAME):$(PG_PASS)@$(PG_HOST):$(PG_PORT)/$(PG_DB)?sslmode=disable
MIGRATIONS_DIR=./migrations

# Help command
help:
	@echo "Database Migration Commands:"
	@echo "  make migrate-create NAME=migration_name   - Create a new migration"
	@echo "  make migrate-up                          - Run all pending migrations"
	@echo "  make migrate-up-one                      - Run only the next pending migration"
	@echo "  make migrate-down                        - Roll back the most recent migration"
	@echo "  make migrate-down-all                    - Roll back all migrations"
	@echo "  make migrate-status                      - Show migration status"
	@echo "  make migrate-force V=version             - Force migration version"

# Create a new migration file
migrate-create:
	@if [ -z "$(NAME)" ]; then echo "Error: NAME is required. Use 'make migrate-create NAME=migration_name'"; exit 1; fi
	@mkdir -p $(MIGRATIONS_DIR)
	@migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME)
	@echo "Created migration files in $(MIGRATIONS_DIR)"

# Run all pending migrations
migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up
	@echo "Applied all pending migrations"

# Run only next pending migration
migrate-up-one:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up 1
	@echo "Applied one pending migration"

# Roll back the most recent migration
migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1
	@echo "Rolled back one migration"

# Roll back all migrations
migrate-down-all:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down
	@echo "Rolled back all migrations"

# Show migration status
migrate-status:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version
	@echo "Checking migration status"

# Force to specific version
migrate-force:
	@if [ -z "$(V)" ]; then echo "Error: V is required. Use 'make migrate-force V=version'"; exit 1; fi
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force $(V)
	@echo "Forced migration version to $(V)"

.PHONY: help migrate-create migrate-up migrate-up-one migrate-down migrate-down-all migrate-status migrate-force




# API Documentation Makefile
swagger-gen:
	@swag init --generalInfo cmd/main.go --output ./docs --parseDependency
	@echo "Swagger documentation generated successfully"

swagger-serve: swagger-gen
	@echo "Starting server with Swagger UI at http://localhost:6655/swagger/index.html"
	@go run main.go

# Clean generated documentation
swagger-clean:
	@rm -rf docs
	@echo "Documentation cleaned"

.PHONY: swagger-setup swagger-gen swagger-serve swagger-clean
