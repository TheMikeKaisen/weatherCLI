# Variables
BINARY_NAME=tempCalc
BUILD_DIR=bin

# Commands
build:
	@echo "Building the project..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go
	@echo "Build complete. Binary located at $(BUILD_DIR)/$(BINARY_NAME)"

clean:
	@echo "Cleaning up build files..."
	@rm -rf $(BUILD_DIR)
	@echo "Cleanup complete."

deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@go mod download
	@echo "Dependencies installed."

makeGlobal:
	@sudo mv bin/tempCalc /usr/local/bin/
