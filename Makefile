BINARY_NAME := tiktoklive
GO_Mod:=Sunny
# Go commands
GO_BUILD := go build
GO_CLEAN := go clean
GO_TIDY := go mod tidy
LDFLAGS := -ldflags="-s -w"
# 生成 Go 代码的目标目录
PROTO_OUT_DIR := tiktok_hack/generated
# Build binary for Windows with CGO enabled
build-windows:
	@echo "Building application for Windows with CGO enabled..."
	set CGO_ENABLED=1 && $(GO_BUILD) $(LDFLAGS) -o $(BINARY_NAME)-windows.exe $(GO_Mod)
# Install dependencies
install:
	@echo "Installing dependencies..."
	$(GO_TIDY)
# Generate Go code from .proto files
proto:
	@echo "Generating Go code from .proto files..."
	protoc -I./tiktok_hack/webcast/proto --go_out=$(PROTO_OUT_DIR) --go_opt=paths=source_relative ./tiktok_hack/webcast/proto/*.proto
# Clean generated files
clean:
	@echo "Cleaning build artifacts..."
	$(GO_CLEAN)
	del /F /Q $(BINARY_NAME)-windows.exe
# Default target
all: build-windows
# Help information
help:
	@echo "Usage:"
	@echo "  make build-windows - Build application for Windows with CGO enabled"
	@echo "  make install       - Install dependencies"
	@echo "  make clean         - Clean build artifacts"
	@echo "  make proto         - Generate Go code from .proto files"
	@echo "  make help          - Display this help message"

.PHONY: build-windows install clean proto help all