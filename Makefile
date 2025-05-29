#TODO
# Define the binary name
BINARY_NAME=myapp

# Define the local bin directory for project-specific tools
LOCAL_BIN_DIR := $(CURDIR)/bin

# Path to the goose binary within the local bin directory
GOOSE_BIN := $(LOCAL_BIN_DIR)/goose

.PHONY: install-tools check-goose-installed

# Helper to check if goose is installed locally, and if not, install it
check-goose-installed:
	@if [ ! -f "$(GOOSE_BIN)" ]; then \
		echo "${YELLOW}Goose binary not found at $(GOOSE_BIN). Installing...${RESET}"; \
		$(MAKE) install-tools; \
	fi
	@if [ ! -f "$(GOOSE_BIN)" ]; then \
		echo "${RED}Error: Failed to install goose. Please check your setup.${RESET}"; \
		exit 1; \
	fi

install-tools: ## Installs Go tools required by the project into a local bin directory.
	@echo "${YELLOW}Installing project tools to $(LOCAL_BIN_DIR)...${RESET}"
	@mkdir -p $(LOCAL_BIN_DIR)
	GOBIN=$(LOCAL_BIN_DIR) go install github.com/pressly/goose/v3/cmd/goose@latest
	@echo "${GREEN}Tools installed.${RESET}"

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