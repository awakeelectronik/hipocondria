package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"poetry-admin/config"
	"poetry-admin/models"

	"github.com/gin-gonic/gin"
)

func GetContents(c *gin.Context) {
	var contents []models.Content

	// Parámetros de consulta
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")

	query := config.DB.Model(&models.Content{})

	// Búsqueda
	if search != "" {
		query = query.Where("title LIKE ? OR text LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Contar total
	var total int64
	query.Count(&total)

	// Paginación
	offset := (page - 1) * limit
	if err := query.Limit(limit).Offset(offset).Order("created_at DESC").Find(&contents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo contenidos"})
		return
	}

	// Formatear respuesta con excerpts
	var response []models.ContentResponse
	for _, content := range contents {
		excerpt := content.Text
		if len(excerpt) > 100 {
			excerpt = excerpt[:100] + "..."
		}

		response = append(response, models.ContentResponse{
			ID:        content.ID,
			Title:     content.Title,
			Text:      content.Text,
			Excerpt:   excerpt,
			CreatedAt: content.CreatedAt,
			UpdatedAt: content.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
		"pagination": gin.H{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": (total + int64(limit) - 1) / int64(limit),
		},
	})
}

func GetContent(c *gin.Context) {
	id := c.Param("id")
	var content models.Content

	println("Fetching content with ID:", id) // Debug log
	if err := config.DB.First(&content, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contenido no encontrado1"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": content})
}

func CreateContent(c *gin.Context) {
	var req models.ContentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos de entrada inválidos",
			"details": err.Error(),
		})
		return
	}

	content := models.Content{
		Title: strings.TrimSpace(req.Title),
		Text:  strings.TrimSpace(req.Text),
	}

	if err := config.DB.Create(&content).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creando contenido"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Contenido creado exitosamente",
		"data":    content,
	})
}

func UpdateContent(c *gin.Context) {
	id := c.Param("id")
	var content models.Content

	if err := config.DB.First(&content, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contenido no encontrado2"})
		return
	}

	var req models.ContentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos de entrada inválidos",
			"details": err.Error(),
		})
		return
	}

	content.Title = strings.TrimSpace(req.Title)
	content.Text = strings.TrimSpace(req.Text)

	if err := config.DB.Save(&content).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error actualizando contenido"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Contenido actualizado exitosamente",
		"data":    content,
	})
}

func DeleteContent(c *gin.Context) {
	id := c.Param("id")
	var content models.Content

	if err := config.DB.First(&content, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contenido no encontrado3"})
		return
	}

	if err := config.DB.Delete(&content).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error eliminando contenido"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contenido eliminado exitosamente"})
}
