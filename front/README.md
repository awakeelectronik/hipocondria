# Hipocondria Frontend

[![Vue.js](https://img.shields.io/badge/Vue.js-3.5.18+-4FC08D.svg?style=flat&logo=vue.js&logoColor=white)](https://vuejs.org/)
[![Vite](https://img.shields.io/badge/Vite-7.1.2+-646CFF.svg?style=flat&logo=vite&logoColor=white)](https://vitejs.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

> **Muchos autores, muchos estilos, un blog** - Frontend moderno para el blog colaborativo de Hipocondria

## 📖 Descripción

Hipocondria Frontend es una aplicación web moderna construida con Vue.js 3 que sirve como interfaz de usuario para el blog colaborativo [hipocondria.co](https://hipocondria.co). La aplicación presenta un diseño limpio y minimalista que permite a los usuarios explorar artículos de múltiples autores con diferentes estilos de escritura.

## ✨ Características

- **🚀 Rendimiento optimizado** - Construido con Vite para desarrollo rápido y builds optimizados
- **📱 Diseño responsive** - Interfaz adaptativa que funciona perfectamente en todos los dispositivos
- **🎨 UI moderna** - Interfaz limpia y minimalista con animaciones suaves
- **📝 Gestión de artículos** - Visualización y navegación por artículos del blog
- **🔄 Carga dinámica** - Sistema de carga asíncrona con indicadores visuales
- **🎯 Navegación intuitiva** - Enrutamiento eficiente entre vistas principales
- **⚡ Componentes reutilizables** - Arquitectura modular con componentes Vue.js

## 🛠️ Tecnologías

- **Frontend Framework**: [Vue.js 3](https://vuejs.org/) (Composition API)
- **Build Tool**: [Vite](https://vitejs.dev/) - Build tool moderno y rápido
- **Routing**: [Vue Router 4](https://router.vuejs.org/) - Enrutamiento del lado del cliente
- **Styling**: CSS moderno con variables CSS y diseño responsive
- **Development**: Hot Module Replacement (HMR) y proxy de desarrollo

## 📁 Estructura del Proyecto

```
front/
├── public/                 # Archivos estáticos públicos
├── src/
│   ├── assets/            # Recursos estáticos (imágenes, iconos)
│   ├── components/        # Componentes Vue reutilizables
│   │   ├── HelloWorld.vue
│   │   └── LoadingGears.vue
│   ├── views/             # Vistas principales de la aplicación
│   │   ├── HomeView.vue   # Vista principal con lista de artículos
│   │   ├── ArticleView.vue # Vista individual de artículo
│   │   └── ContactView.vue # Página de contacto
│   ├── App.vue            # Componente raíz de la aplicación
│   ├── main.js            # Punto de entrada de la aplicación
│   └── style.css          # Estilos globales
├── index.html             # Plantilla HTML principal
├── package.json           # Dependencias y scripts del proyecto
├── vite.config.js         # Configuración de Vite
└── README.md              # Este archivo
```

## 🚀 Instalación y Uso

### Prerrequisitos

- [Node.js](https://nodejs.org/) (versión 16 o superior)
- [npm](https://www.npmjs.com/) o [yarn](https://yarnpkg.com/)

### Instalación

1. **Clonar el repositorio**
   ```bash
   git clone https://github.com/awakeelectronik/hipocondria.git
   cd hipocondria/front
   ```

2. **Instalar dependencias**
   ```bash
   npm install
   # o
   yarn install
   ```

3. **Ejecutar en modo desarrollo**
   ```bash
   npm run dev
   # o
   yarn dev
   ```

4. **Abrir en el navegador**
   ```
   http://localhost:5173
   ```

### Scripts Disponibles

- `npm run dev` - Inicia el servidor de desarrollo con HMR
- `npm run build` - Construye la aplicación para producción
- `npm run preview` - Previsualiza la build de producción localmente

## 🌐 API y Backend

La aplicación se conecta con la API de Hipocondria en `https://api.hipocondria.co` para:

- **Obtener artículos**: `GET /articles` - Lista todos los artículos disponibles
- **Obtener artículo específico**: `GET /articles/{id}` - Recupera un artículo por ID

### Configuración de Proxy

Durante el desarrollo, Vite está configurado para hacer proxy de las llamadas a la API:

```javascript
// vite.config.js
server: {
  proxy: {
    '/api': {
      target: 'https://api.hipocondria.co',
      changeOrigin: true,
      rewrite: (path) => path.replace(/^\/api/, ''),
      secure: true,
    }
  }
}
```

## 🎨 Componentes Principales

### LoadingGears.vue
Componente de carga personalizado con animación de engranajes giratorios, diseñado para proporcionar feedback visual durante las operaciones asíncronas.

### HomeView.vue
Vista principal que muestra la lista de artículos disponibles con:
- Títulos y fechas de publicación
- Iconos dinámicos basados en el ID del artículo
- Navegación a artículos individuales
- Manejo de estados de carga y error

### ArticleView.vue
Vista para mostrar artículos individuales con:
- Título del artículo
- Contenido renderizado en HTML
- Navegación de regreso a la lista principal
- Manejo de errores de carga

### ContactView.vue
Página de contacto simple con información de contacto del blog.

## 📱 Diseño Responsive

La aplicación está diseñada para funcionar perfectamente en todos los tamaños de pantalla:

- **Desktop**: Diseño completo con tipografía grande y espaciado generoso
- **Tablet**: Adaptación automática para pantallas medianas
- **Mobile**: Optimización para dispositivos móviles con navegación táctil

## 🔧 Configuración de Desarrollo

### Variables de Entorno

Para desarrollo local, puedes crear un archivo `.env.local`:

```env
VITE_API_BASE_URL=https://api.hipocondria.co
```

### Proxy de Desarrollo

El proxy de desarrollo está configurado para redirigir las llamadas `/api/*` al backend de producción, facilitando el desarrollo local.

## 🚀 Despliegue

### Build de Producción

```bash
npm run build
```

Esto generará una carpeta `dist/` con los archivos optimizados para producción.

### Servidor de Previsualización

```bash
npm run preview
```

Permite previsualizar la build de producción antes del despliegue.

## 🤝 Contribución

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo `LICENSE` para más detalles.

## 📞 Contacto

- **Email**: contacto@hipocondria.co
- **Sitio Web**: [hipocondria.co](https://hipocondria.co)
- **GitHub**: [@awakeelectronik](https://github.com/awakeelectronik)

## 🙏 Agradecimientos

- [Vue.js](https://vuejs.org/) por el framework increíble
- [Vite](https://vitejs.dev/) por la herramienta de build moderna
- [Vue Router](https://router.vuejs.org/) por el sistema de enrutamiento
- La comunidad de desarrolladores que contribuye al ecosistema Vue

---

**Desarrollado con ❤️ por el equipo de Hipocondria**
