.PHONY: build test clean install run fmt vet help

# Build variables
BINARY_NAME=pokedexcli
CMD_PATH=./cmd/pokedexcli
BUILD_DIR=.
GO_FILES=$(shell find . -name "*.go" -type f)

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the binary
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_PATH)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

test: ## Run tests
	@echo "Running tests..."
	@go test ./...

clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	@rm -f $(BUILD_DIR)/$(BINARY_NAME)
	@rm -f coverage.out coverage.html
	@go clean

install: ## Install the binary to $GOPATH/bin
	@echo "Installing $(BINARY_NAME) to $$GOPATH/bin..."
	@go install $(CMD_PATH)

run: build ## Build and run the application
	@echo "Starting $(BINARY_NAME)..."
	@./$(BINARY_NAME)

fmt: ## Format Go code
	@echo "Formatting code..."
	@gofmt -s -w .

vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...

lint: vet ## Run linting (currently just vet)
	@echo "Linting complete"

coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out

coverage-html: coverage ## Generate HTML coverage report
	@echo "Generating HTML coverage report..."
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

dev-setup: ## Set up development environment
	@echo "Setting up development environment..."
	@go mod tidy
	@make fmt
	@make test
	@echo "Development setup complete"

# Build info
info: ## Show build information
	@echo "Binary name: $(BINARY_NAME)"
	@echo "Command path: $(CMD_PATH)"
	@echo "Go version: $$(go version)"
	@echo "Go files found: $$(echo $(GO_FILES) | wc -w)"