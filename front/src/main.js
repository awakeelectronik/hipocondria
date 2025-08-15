import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import './style.css'
import App from './App.vue'
import HomeView from './views/HomeView.vue'
import ArticleView from './views/ArticleView.vue'
import ContactView from './views/ContactView.vue'
import ContactDesktop from './views/ContactDesktop.vue'
import sanitizeHtml from 'sanitize-html';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', name: 'home', component: HomeView },
    { path: '/articles/:id', name: 'article', component: ArticleView, props: true },
    { path: '/contacto', name: 'contact', component: ContactView },
    { path: '/contactoD', name: 'contactD', component: ContactDesktop },
  ],
  scrollBehavior() {
    return { top: 0 }
  }
})
const app = createApp(App)

// Asignar sanitize como propiedad global
app.config.globalProperties.$sanitize = sanitizeHtml

// Registrar el router y montar la app
app.use(router)
app.mount('#app')
