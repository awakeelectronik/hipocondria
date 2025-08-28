package routes

import (
	"net/http"

	"poetry-admin/controllers"
	"poetry-admin/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Ruta de salud
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Servidor funcionando correctamente"})
	})

	// Grupo de rutas de API
	api := r.Group("/api")
	{
		// Rutas de autenticación (públicas)
		auth := api.Group("/auth")
		{
			auth.POST("/login", controllers.Login)
			auth.POST("/register", controllers.Register)
		}

		// Rutas protegidas
		protected := api.Group("/")
		protected.Use(middleware.AuthRequired())
		{
			// Rutas de contenido (poemas)
			content := protected.Group("/poems")
			{
				content.GET("", controllers.GetContents)
				content.GET("/:id", controllers.GetContent)
				content.POST("/", controllers.CreateContent)
				content.PUT("/:id", controllers.UpdateContent)
				content.DELETE("/:id", controllers.DeleteContent)
			}
		}
	}
}
