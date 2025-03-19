# Параметры
BINARY       = tmdb-cli
CMD_DIR      = ./cmd
MAIN_FILE    = $(CMD_DIR)/main.go
DOCKER_IMAGE = tm-api

.PHONY: all build run test tidy clean docker-build docker-up docker-down docker-clean

# Цель по умолчанию — сборка проекта
all: build

# Сборка бинарного файла
build:
	@echo "Сборка $(BINARY)..."
	go build -o $(BINARY) $(MAIN_FILE)

# Запуск приложения (сборка и запуск)
run: build
	@echo "Запуск $(BINARY)..."
	./$(BINARY)

# Запуск тестов
test:
	@echo "Запуск тестов..."
	go test ./...

# Приведение зависимостей в порядок
tidy:
	@echo "Приведение зависимостей в порядок..."
	go mod tidy

# Очистка: удаление бинарного файла
clean:
	@echo "Очистка проекта..."
	rm -f $(BINARY)

# Сборка Docker-образа
docker-build:
	@echo "Сборка Docker-образа $(DOCKER_IMAGE)..."
	docker build -t $(DOCKER_IMAGE) .

# Запуск docker-compose
docker-up:
	@echo "Запуск docker-compose..."
	docker-compose up --build

# Остановка docker-compose
docker-down:
	@echo "Остановка docker-compose..."
	docker-compose down

# Удаление Docker-образа
docker-clean:
	@echo "Удаление Docker-образа $(DOCKER_IMAGE)..."
	docker rmi $(DOCKER_IMAGE) || true
