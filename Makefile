.PHONY: build test clean install

BUILD_DIR=bin
BINARY_NAME=mob

build:
	@echo "Building mob..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/mob

build-all:
	@echo "Building mob for all platforms..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 ./cmd/mob
	GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 ./cmd/mob

test:
	@echo "Running tests..."
	go test -v ./...

clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_DIR)

install: build
	@echo "Installing mob..."
	install -m 755 $(BUILD_DIR)/$(BINARY_NAME) ${HOME}/.local/bin/$(BINARY_NAME)
	@echo "✅ Installed to ${HOME}/.local/bin/$(BINARY_NAME)"

lint:
	@echo "Running linter..."
	golangci-lint run

fmt:
	@echo "Formatting code..."
	go fmt ./...
