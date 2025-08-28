# Makefile para el proyecto de panel de administración de poesía

# Variables
BINARY_NAME=poetry-admin

# Comandos de desarrollo
.PHONY: dev build run clean test help

# Comando por defecto
help:
	@echo "Comandos disponibles:"
	@echo "  dev          - Ejecutar en modo desarrollo con recarga automática"
	@echo "  build        - Compilar la aplicación"
	@echo "  run          - Ejecutar la aplicación"
	@echo "  clean        - Limpiar archivos compilados"
	@echo "  test         - Ejecutar tests"
	@echo "  install-deps - Instalar dependencias de desarrollo"

# Desarrollo
dev:
	@echo "🔄 Iniciando servidor de desarrollo..."
	@if command -v air >/dev/null 2>&1; then \
		air; \
	else \
		echo "📦 Air no encontrado. Ejecutando go run..."; \
		go run main.go; \
	fi

# Compilación
build:
	@echo "🔨 Compilando aplicación..."
	@go build -o $(BINARY_NAME) main.go
	@echo "✅ Compilación completada: $(BINARY_NAME)"

# Ejecutar
run: build
	@echo "🚀 Ejecutando aplicación..."
	@./$(BINARY_NAME)

# Limpiar
clean:
	@echo "🧹 Limpiando archivos..."
	@rm -f $(BINARY_NAME)
	@rm -rf tmp/
	@go clean

# Tests
test:
	@echo "🧪 Ejecutando tests..."
	@go test -v ./...

# Instalar dependencias de desarrollo
install-deps:
	@echo "📦 Instalando dependencias de desarrollo..."
	@go install github.com/cosmtrek/air@latest
	@go mod download

# Producción
deploy: build
	@echo "🚀 Desplegando en producción..."
	@# Agregar comandos de despliegue aquí

# Base de datos
db-reset:
	@echo "🔄 Reiniciando base de datos..."
	@mysql -u root -p poetry_db < database/init.sql

# Linting y formato
fmt:
	@echo "📝 Formateando código..."
	@go fmt ./...

lint:
	@echo "🔍 Ejecutando linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "⚠️ golangci-lint no encontrado. Instalar con: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi
