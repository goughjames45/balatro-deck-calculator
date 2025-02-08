# Define variables
PROJECT_NAME := main
GO_FILES := $(shell find . -name '*.go' -not -path "./vendor/*")
BINARY_PATH := ./bin
BINARY_NAME := $(PROJECT_NAME)
LDFLAGS := -ldflags="-s -w"

# Default target
all: build

# Build the application
build:
	go build $(LDFLAGS) -o $(BINARY_PATH)/$(BINARY_NAME) ./cmd/$(PROJECT_NAME)

# Run the application
run: build
	$(BINARY_PATH)/$(BINARY_NAME)

# Format code
fmt:
	go fmt $(GO_FILES)

# Clean build artifacts
clean:
	rm -rf $(BINARY_PATH)

# Install dependencies
deps:
	go mod tidy

# Test the application
test:
	go test ./...

# Show help message
help:
	@echo "Usage: make <target>"
	@echo "Targets:"
	@echo "  build: Build the application"
	@echo "  run: Run the application"
	@echo "  fmt: Format code"
	@echo "  clean: Clean build artifacts"
	@echo "  deps: Install dependencies"
	@echo "  test: Run tests"
