<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import LoadingGears from '../components/LoadingGears.vue'

const router = useRouter()
const loading = ref(true)
const error = ref('')
const articles = ref([])

async function fetchArticles() {
  loading.value = true
  error.value = ''
  try {
    const res = await fetch('https://api.hipocondria.co/articles')
    if (!res.ok) throw new Error('Error al cargar artículos')
    articles.value = await res.json()
  } catch (e) {
    error.value = e.message || 'Error desconocido'
  } finally {
    loading.value = false
  }
}

function openArticle(id) {
  router.push({ name: 'article', params: { id } })
}

onMounted(fetchArticles)

const iconNames = ['play','book','grid','code','docs']
const iconSvgs = {
  play: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polygon points="8 5 8 19 19 12"/></svg>`,
  book: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 5h11a3 3 0 0 1 3 3v11H7a3 3 0 0 0-3 3V5z"/><path d="M18 19V8a3 3 0 0 0-3-3"/></svg>`,
  grid: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="3" width="8" height="8"/><rect x="13" y="3" width="8" height="8"/><rect x="3" y="13" width="8" height="8"/><rect x="13" y="13" width="8" height="8"/></svg>`,
  code: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="8 4 3 12 8 20"></polyline><polyline points="16 4 21 12 16 20"></polyline></svg>`,
  docs: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"></polyline><line x1="16" y1="13" x2="8" y2="13"></line><line x1="16" y1="17" x2="8" y2="17"></line></svg>`
}

function iconNameFor(seed) {
  const s = String(seed ?? '')
  let h = 0
  for (let i = 0; i < s.length; i++) { h = (h << 5) - h + s.charCodeAt(i); h |= 0 }
  const idx = Math.abs(h) % iconNames.length
  return iconNames[idx]
}

function getIconSvg(name) {
  return iconSvgs[name] || iconSvgs.play
}
</script>

<template>
  <section>
    <h2 class="title">Muchos autores, muchos estilos, un blog</h2>
    <div v-if="loading" class="loading-wrap"><LoadingGears /></div>
    <p v-else-if="error" class="error">{{ error }}</p>
    <ul v-else class="post-list">
      <li v-for="a in articles" :key="a.id" class="post-item">
        <router-link class="post-link" :to="{ name: 'article', params: { id: a.id } }">
          <span class="card-icon" aria-hidden="true" v-html="getIconSvg(iconNameFor(a.id))"></span>
          <div class="card-body">
            <h3 class="post-title">{{ a.title }}</h3>
            <div class="post-meta">{{ a?.date || '' }}</div>
          </div>
          <svg class="card-arrow" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M7 17L17 7"/><path d="M7 7h10v10"/></svg>
        </router-link>
      </li>
    </ul>
  </section>
</template>

<style scoped>
.title { font-size: 1.75rem; font-weight: 700; margin-bottom: 1rem; }
.error { color: #dc2626; }
</style>


