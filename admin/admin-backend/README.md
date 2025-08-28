# Panel de Administración de Poesía - Backend

Backend en Go con framework Gin para el panel de administración de poesía.

## Características

- **Autenticación JWT**: Sistema de login seguro con tokens JWT
- **CRUD completo**: Operaciones Create, Read, Update, Delete para poemas
- **Base de datos MySQL**: Usando GORM como ORM
- **Middleware de seguridad**: Autenticación y autorización
- **CORS configurado**: Para comunicación con frontend
- **Validación de datos**: Validación de entradas con Gin
- **Paginación**: Lista de poemas con paginación
- **Búsqueda**: Búsqueda en títulos y contenido

## Instalación

### Prerrequisitos

- Go 1.21+
- MySQL 8.0+
- Git

### Configuración

1. Clonar el repositorio:
\`\`\`bash
git clone <repository-url>
cd poetry-admin-backend
\`\`\`

2. Instalar dependencias:
\`\`\`bash
go mod download
\`\`\`

3. Configurar variables de entorno:
\`\`\`bash
cp .env.example .env
\`\`\`

4. Editar `.env` con tus configuraciones:
\`\`\`env
DB_HOST=localhost
DB_PORT=3306
DB_USER=tu_usuario
DB_PASSWORD=tu_password
DB_NAME=poetry_db
JWT_SECRET=tu-super-secreto-jwt-key
PORT=8000
\`\`\`

5. Crear la base de datos MySQL:
\`\`\`sql
CREATE DATABASE poetry_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
\`\`\`

## Uso

### Desarrollo

\`\`\`bash
go run main.go
\`\`\`

### Producción

\`\`\`bash
go build -o poetry-admin main.go
./poetry-admin
\`\`\`

## API Endpoints

### Autenticación

- **POST** `/api/auth/login` - Login de usuario
- **POST** `/api/auth/register` - Registro de usuario

### Poemas (requiere autenticación)

- **GET** `/api/poems` - Listar poemas (con paginación y búsqueda)
- **GET** `/api/poems/:id` - Obtener poema por ID
- **POST** `/api/poems` - Crear nuevo poema
- **PUT** `/api/poems/:id` - Actualizar poema
- **DELETE** `/api/poems/:id` - Eliminar poema

### Otros

- **GET** `/health` - Estado del servidor

## Estructura del Proyecto

\`\`\`
.
├── main.go              # Punto de entrada
├── go.mod               # Dependencias
├── .env.example         # Variables de entorno ejemplo
├── README.md            # Documentación
├── config/
│   └── database.go      # Configuración de BD
├── controllers/
│   ├── auth.go          # Controladores de autenticación
│   └── content.go       # Controladores de contenido
├── middleware/
│   └── auth.go          # Middleware de autenticación
├── models/
│   ├── user.go          # Modelo de usuario
│   └── content.go       # Modelo de contenido
└── routes/
    └── routes.go        # Configuración de rutas
\`\`\`

## Características Técnicas

### Autenticación JWT

- Tokens firmados con HS256
- Expiración de 24 horas
- Claims personalizados con información del usuario

### Base de Datos

- Soft deletes habilitados
- Timestamps automáticos
- Validaciones a nivel de modelo y base de datos

### Seguridad

- Passwords hasheados con bcrypt
- Validación de entrada en todos los endpoints
- CORS configurado para frontends específicos
- Headers de seguridad

### Performance

- Conexión persistente a base de datos
- Queries optimizados con GORM
- Paginación en listados
- Índices en campos de búsqueda

## Usuario Administrador por Defecto

El sistema crea automáticamente un usuario administrador:

- **Email**: admin@poesia.com
- **Username**: admin
- **Password**: admin123

**¡Cambia estas credenciales en producción!**

## Ejemplos de Uso

### Login

\`\`\`bash
curl -X POST http://localhost:8000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin123"}'
\`\`\`

### Crear Poema

\`\`\`bash
curl -X POST http://localhost:8000/api/poems \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "Mi Poema",
    "text": "Contenido del poema aquí..."
  }'
\`\`\`

### Listar Poemas con Paginación

\`\`\`bash
curl "http://localhost:8000/api/poems?page=1&limit=10&search=amor" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
\`\`\`

## Contribución

1. Fork el proyecto
2. Crea una rama feature: `git checkout -b feature/nueva-caracteristica`
3. Commit tus cambios: `git commit -am 'Agregar nueva característica'`
4. Push a la rama: `git push origin feature/nueva-caracteristica`
5. Crea un Pull Request

## Licencia

MIT License - ver el archivo LICENSE para detalles.

## Soporte

Para preguntas o problemas, crear un issue en el repositorio.
