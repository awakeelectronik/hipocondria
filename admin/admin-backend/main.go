package main

import (
	"log"
	"os"
	"time"

	"poetry-admin/config"
	"poetry-admin/controllers"
	_ "poetry-admin/middleware"
	"poetry-admin/models"
	"poetry-admin/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Inicializar base de datos
	config.ConnectDatabase()

	// Ejecutar migraciones
	config.DB.AutoMigrate(&models.User{}, &models.Content{})

	// Crear usuario administrador por defecto si no existe
	createDefaultAdmin()

	// Configurar Gin
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Configurar CORS
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080", "*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(corsConfig))

	// Middleware de logging
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Configurar rutas
	routes.SetupRoutes(r)

	// Obtener puerto del entorno o usar 8000 por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Servidor iniciando en puerto %s", port)
	log.Fatal(r.Run(":" + port))
}

func createDefaultAdmin() {
	var user models.User
	if err := config.DB.Where("email = ?", "admin@poesia.com").First(&user).Error; err != nil {
		// Usuario admin no existe, crear uno
		hashedPassword, _ := controllers.HashPassword("admin123")
		admin := models.User{
			Username: "admin",
			Email:    "admin@poesia.com",
			Password: hashedPassword,
			Role:     "admin",
		}
		config.DB.Create(&admin)
		log.Println("Usuario administrador creado: admin@poesia.com / admin123")
	}
}
