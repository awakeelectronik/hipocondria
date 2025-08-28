-- Script de inicialización de la base de datos para el panel de poesía
-- Crear la base de datos
CREATE DATABASE IF NOT EXISTS poetry_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE poetry_db;

-- Tabla de usuarios
CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(191) UNIQUE NOT NULL,
    email VARCHAR(191) UNIQUE NOT NULL,
    password VARCHAR(191) NOT NULL,
    role VARCHAR(50) DEFAULT 'user',
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    deleted_at DATETIME(3) NULL,
    INDEX idx_users_deleted_at (deleted_at),
    INDEX idx_users_username (username),
    INDEX idx_users_email (email)
);

-- Tabla de contenido (poemas)
CREATE TABLE IF NOT EXISTS contents (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    text MEDIUMTEXT NOT NULL,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    deleted_at DATETIME(3) NULL,
    INDEX idx_contents_deleted_at (deleted_at),
    INDEX idx_contents_title (title),
    FULLTEXT(title, text)
);

-- Insertar usuario administrador por defecto (password: admin123)
-- Hash generado con bcrypt para "admin123"
INSERT IGNORE INTO users (username, email, password, role, created_at, updated_at) 
VALUES ('admin', 'admin@poesia.com', '$2a$14$Zl6z9/K.5vJ8yXV5gYHO5eD0qYBn2dVgXUfGkQHqSzVxWEyDfH8G.', 'admin', NOW(), NOW());

-- Insertar algunos poemas de ejemplo
INSERT IGNORE INTO contents (id, title, text, created_at, updated_at) VALUES 
(1, 'Caminante no hay camino', 'Caminante, son tus huellas
el camino y nada más;
Caminante, no hay camino,
se hace camino al andar.
Al andar se hace el camino,
y al volver la vista atrás
se ve la senda que nunca
se ha de volver a pisar.
Caminante no hay camino
sino estelas en la mar.', NOW(), NOW()),

(2, 'Volverán las oscuras golondrinas', 'Volverán las oscuras golondrinas
en tu balcón sus nidos a colgar,
y otra vez con el ala a sus cristales
jugando llamarán.
Pero aquellas que el vuelo refrenaban
tu hermosura y mi dicha a contemplar,
aquellas que aprendieron nuestros nombres...
¡esas... no volverán!', NOW(), NOW()),

(3, 'Mientras por competir con tu cabello', 'Mientras por competir con tu cabello,
oro bruñido al sol relumbra en vano;
mientras con menosprecio en medio el llano
mira tu blanca frente el lilio bello;
mientras a cada labio, por cogello,
siguen más ojos que al clavel temprano,
y mientras triunfa con desdén lozano
del luciente cristal tu gentil cuello:
goza cuello, cabello, labio y frente,
antes que lo que fue en tu edad dorada
oro, lilio, clavel, cristal luciente,
no sólo en plata o vïola troncada
se vuelva, mas tú y ello juntamente
en tierra, en humo, en polvo, en sombra, en nada.', NOW(), NOW());
