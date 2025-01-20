GO = go
GOFMT = gofmt
GOCOVER = go test -coverprofile=coverage.out
GO_BIN = $(GO) build -o bin/app

PKG = github.com/sk-pathak/fissionNode-backend
BUILD_DIR = bin
SRC_DIR = cmd
APP_NAME = fissionNode-backend

.PHONY: all help build exec run test lint format create-migration apply-migration sqlc clean docker
all: help

.PHONY: help
help:
	@echo "Makefile for Go project"
	@echo ""
	@echo "Available targets:"
	@echo "  build            Build the Go application into a binary (bin/app/fissionNode-backend)"
	@echo "  exec             Build the application and run it immediately"
	@echo "  run              Run the Go application directly (cmd/server/main.go)"
	@echo "  test             Run all tests in the project"
	@echo "  format           Format the Go source code using gofmt"
	@echo "  create-migration Create a new database migration (requires 'name' argument)"
	@echo "  apply-migration  Apply all pending database migrations"
	@echo "  sqlc             Run SQL code generation using sqlc"
	@echo "  clean            Remove build artifacts from the bin directory"
	@echo "  docker           Start the application in a Docker environment using docker-compose"
	@echo "  help             Show this help message"
	@echo ""
	@echo "Usage:"
	@echo "  make <target>    Run the specified target command"

build:
	$(GO) build -o $(BUILD_DIR)/app/${APP_NAME} $(SRC_DIR)/server/main.go

exec:
	make build && $(BUILD_DIR)/app/${APP_NAME}

run:
	$(GO) run $(SRC_DIR)/server/main.go

test:
	$(GO) test ./...

format:
	@echo "Running Go formatter..."
	@find . -name '*.go' | xargs $(GOFMT) -s -w

create-migration:
	./scripts/create_migration.sh $(name)

apply-migration:
	./scripts/apply_migration.sh

sqlc:
	./scripts/sqlc.sh

clean:
	rm -rf $(BUILD_DIR)/app

docker:
	docker compose up
