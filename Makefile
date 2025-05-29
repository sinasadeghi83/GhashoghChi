#TODO
# Define the binary name
BINARY_NAME=myapp

# Default target
# all: build

# # Build the project
# build:
# 	go build -o ${BINARY_NAME} main.go

# # Run the application
# run:
# 	go run main.go

run-%:
	go run $*/$(MODULE)/main.go

# Clean build artifacts
# clean:
# 	go clean
# 	rm -f ${BINARY_NAME}

# # Run tests
# test:
# 	go test ./...

# # Run tests with coverage
# test-cover:
# 	go test -cover ./...

# # Install dependencies
# deps:
# 	go mod download

# # Format code
# fmt:
# 	go fmt ./...

# # Vet code (static analysis)
# vet:
# 	go vet ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

.PHONY: all build run clean test test-cover deps fmt vet lint