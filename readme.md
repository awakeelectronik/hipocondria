# Hipocondria API

Minimal HTTP server for personal blog API  
Implementado en Go con Gin para servir artículos almacenados en una base de datos MySQL.

## 🔧 Requisitos

- **Go 1.24.4**
- **MySQL** o **MariaDB**
- Credenciales de base de datos con acceso al esquema `hipocondria`

## 🚀 Instalación y ejecución

Clona el repositorio, instala dependencias y ejecuta el servidor:
```bash

git clone https://github.com/awakeelectronik/hipocondria.git
cd hipocondria/api
go mod download
go mod tidy
go build
./hipocondria # o alternativamente:
go run main.go
```


Por defecto, el servidor escucha en **localhost:5001**.

## 📡 Endpoints disponibles

### Obtener todos los artículos

`curl http://localhost:5001/articles`



### Obtener artículo con ID = 1

`curl http://localhost:5001/articles/1`



Las respuestas se devuelven en formato JSON con una estructura estándar:
`
{
"id": 1,
"title": "Título del artículo"
}
`


En caso de error, se retorna:
`
{
"error": "Mensaje genérico",
"detalle": "Detalle técnico del error"
}`



## 🧱 Diseño del proyecto

- **Responsabilidades bien diferenciadas**: lógica de servidor, rutas, manejo de base de datos, errores y respuestas separadas y organizadas.
- **Manejo de errores detallado**: cada error devuelve una clave `detalle` con información útil para depuración o usuario.
- **Respuestas JSON consistentes**: siempre con formato uniforme para facilitar integración con front-end o consumo por terceros.
- **Conexión segura a la base de datos**: uso de `database/sql`, validación de conexión (`Ping()`) y cierre adecuado con `defer`.
- **Arquitectura RESTful** minimalista, clara y fácil de extender.

## ⚙️ Variables de entorno (si aplica)

Define tus credenciales de conexión (como mínimo):
```bash
DB_USER=usuario
DB_PASS=contraseña
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=hipocondria
```

