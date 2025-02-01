# Variable
APP_NAME=PAPER-WALLET-CORE-SERVICE
BUILD_DIR=bin
GOFILES=$(shell find . -name '*.go' -not -path "./vendor/*")
GOTEST_FLAGS=-race -coverprofile=coverage.out -covermode=atomic

# Default command
.PHONY: all
all: build

# Install dependencies
.PHONY: deps
deps:
	go mod tidy

# Code Format
.PHONY: fmt
fmt:
	go fmt ./...

# Build application
.PHONY: build
build:
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/main.go

# Run application
.PHONY: run
run: build
	./$(BUILD_DIR)/$(APP_NAME)

# Run tests
.PHONY: test
test:
	go test ./internal/... $(GOTEST_FLAGS)

# Delete binary
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

# Update swagger documentation
.PHONY: swag-init
swag-init:
	swag init -g cmd/main.go