#!/bin/bash

# Script de instalación para el panel de administración de poesía

echo "🚀 Iniciando instalación del Panel de Administración de Poesía..."

# Verificar Go
if ! command -v go &> /dev/null; then
    echo "❌ Go no está instalado. Por favor instala Go 1.21+ primero."
    exit 1
fi

# Verificar versión de Go
GO_VERSION=$(go version | grep -o 'go[0-9.]*' | sed 's/go//')
if [ "$(printf '%s\n' "1.21" "$GO_VERSION" | sort -V | head -n1)" != "1.21" ]; then
    echo "❌ Se requiere Go 1.21 o superior. Versión actual: $GO_VERSION"
    exit 1
fi

echo "✅ Go versión $GO_VERSION detectado"

# Instalar dependencias
echo "📦 Instalando dependencias..."
go mod download

# Verificar MySQL
if ! command -v mysql &> /dev/null; then
    echo "⚠️ MySQL no está instalado. Asegúrate de tener MySQL 8.0+ instalado y ejecutándose."
fi

# Crear archivo .env si no existe
if [ ! -f .env ]; then
    echo "📝 Creando archivo .env..."
    cp .env.example .env
    echo "⚠️ Por favor edita el archivo .env con tus configuraciones de base de datos"
fi

# Compilar aplicación
echo "🔨 Compilando aplicación..."
go build -o poetry-admin main.go

echo "✅ Instalación completada!"
echo ""
echo "📋 Próximos pasos:"
echo "1. Edita el archivo .env con tus configuraciones"
echo "2. Crea la base de datos: mysql -u root -p < database/init.sql"
echo "3. Ejecuta la aplicación: ./poetry-admin"
echo ""
echo "🌐 El servidor estará disponible en: http://localhost:8000"
echo "👤 Usuario por defecto: admin / admin123"
