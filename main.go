package main

import (
    "database/sql"
    "fmt"
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gin-contrib/cors"
)

// Modelo de datos para lista de artículos
type ApiArticleList struct {
    Id    int    `json:"id"`
    Title string `json:"title"`
}

// Modelo de datos para artículo completo
type ApiArticle struct {
    Id    int    `json:"id"`
    Title string `json:"title"`
    Text  string `json:"text"`
}

var db *sql.DB // Conexión a base de datos

func main() {
    // Configurar conexión a MySQL
    dsn := "tu_usuario:tu_contraseña@tcp(localhost:3306)/hipocondria"
    
    // Abrir conexión
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Printf("Error al conectar a la base de datos: %v\n", err)
        return // Terminar si hay error de conexión
    }

    // Verificar conexión
    if err = db.Ping(); err != nil {
        fmt.Printf("Error al verificar conexión: %v\n", err)
        return
    }

    // Cerrar conexión al final
    defer db.Close()

    // Configurar router Gin
    router := gin.Default()
    router.Use(cors.New(cors.Config{
        AllowOrigins:  []string{"*"},
        AllowMethods:  []string{"GET"},
        AllowHeaders:  []string{"Content-Type"},
    }))

    // Definir rutas
    router.GET("/articles", getArticles)
    router.GET("/articles/:id", getArticleById)

    // Iniciar servidor en puerto 8080
    router.Run(":8080")
}

// Controlador para listar artículos (solo ID y título)
func getArticles(c *gin.Context) {
    articles := []ApiArticleList{}
    
    rows, err := db.Query("SELECT id, title FROM content")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Error al obtener artículos",
            "detalle": err.Error(),
        })
        return
    }
    defer rows.Close()

    // Recargar todos los resultados
    for rows.Next() {
        var article ApiArticleList
        if err := rows.Scan(&article.Id, &article.Title); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error":   "Error al leer artículo",
                "detalle": err.Error(),
            })
            return
        }
        articles = append(articles, article)
    }

    c.JSON(http.StatusOK, articles)
}

// Controlador para obtener un artículo completo por ID
func getArticleById(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "ID inválido",
            "detalle": "El ID debe ser un número entero",
        })
        return
    }

    var article ApiArticle
    row := db.QueryRow("SELECT * FROM content WHERE id = ?", id)
    
    if err := row.Scan(&article.Id, &article.Title, &article.Text); err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{
                "error":   "Artículo no encontrado",
                "detalle": "No existe un artículo con este ID",
            })
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error":   "Error al obtener artículo",
                "detalle": err.Error(),
            })
        }
        return
    }

    c.JSON(http.StatusOK, article)
}
