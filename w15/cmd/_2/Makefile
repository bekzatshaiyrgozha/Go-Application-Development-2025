# Переменные
APP_NAME = myapp
BUILD_DIR = build
MAIN_PKG = .
DOCKER_IMAGE = myapp:latest
VERSION ?= 1.0.0

# Переменные для Go
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
GOMOD = $(GOCMD) mod
GOFMT = $(GOCMD) fmt

# Переменные времени сборки для внедрения в бинарный файл
BUILD_TIME = $(shell date -u '+%Y-%m-%d %H:%M:%S')
GIT_COMMIT = $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
LD_FLAGS = -ldflags "-X 'main.Version=$(VERSION)' -X 'main.BuildTime=$(BUILD_TIME)' -X 'main.GitCommit=$(GIT_COMMIT)'"

.PHONY: all build clean test run help init-module

# Цель по умолчанию
all: build

# Инициализация модуля (если go.mod отсутствует)
init-module:
	@if [ ! -f go.mod ]; then \
		echo "Initializing Go module..."; \
		$(GOMOD) init example.com/myapp; \
	fi

# Сборка приложения
build: init-module
	@echo "Building application..."
	@mkdir -p $(BUILD_DIR)
	@$(GOBUILD) $(LD_FLAGS) -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PKG)
	@echo "Build completed: $(BUILD_DIR)/$(APP_NAME)"

# Очистка сборки
clean:
	@echo "Cleaning..."
	@$(GOCLEAN)
	@rm -rf $(BUILD_DIR)
	@echo "Clean completed"

# Запуск тестов
test: init-module
	@echo "Running tests..."
	@$(GOTEST) -v ./...
	@echo "Tests completed"

# Запуск приложения
run: build
	@echo "Running application..."
	@./$(BUILD_DIR)/$(APP_NAME)

# Подсказка
help:
	@echo "Available commands:"
	@echo "  make all         - Build the application"
	@echo "  make build       - Build the application"
	@echo "  make clean       - Clean build artifacts"
	@echo "  make test        - Run tests"
	@echo "  make run         - Build and run the application"
	@echo "  make init-module - Initialize Go module (if needed)"
	@echo "  make help        - Show this help message"