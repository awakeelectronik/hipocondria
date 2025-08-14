<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import LoadingGears from '../components/LoadingGears.vue'

const route = useRoute()
const loading = ref(true)
const error = ref('')
const article = ref(null)

async function fetchArticle(id) {
  loading.value = true
  error.value = ''
  try {
    const res = await fetch(`https://api.hipocondria.co/articles/${id}`)
    if (!res.ok) throw new Error('Error al cargar artículo')
    article.value = await res.json()
  } catch (e) {
    error.value = e.message || 'Error desconocido'
  } finally {
    loading.value = false
  }
}

onMounted(() => fetchArticle(route.params.id))
watch(() => route.params.id, (id) => fetchArticle(id))
</script>

<template>
  <article>
    <router-link to="/">← Volver</router-link>
    <div v-if="loading" class="loading-wrap"><LoadingGears /></div>
    <div v-else-if="error" class="error">{{ error }}</div>
    <div v-else>
      <h2 class="title">{{ article?.title }}</h2>
      <div v-if="article?.text" class="prose" v-html="article.text" />
      <pre v-else class="raw">{{ article }}</pre>
    </div>
  </article>
</template>

<style scoped>
.title { font-size: 4.8rem; margin: 1rem 0; font-weight: 800; font-family: var(--font-display); line-height: 1.1; }

@media (max-width: 768px) {
  .title { font-size: 3.2rem; margin: .75rem 0; }
}

@media (max-width: 480px) {
  .title { font-size: 2.4rem; margin: .5rem 0; line-height: 1.2; }
}

@media (max-width: 360px) {
  .title { font-size: 1.8rem; margin: .4rem 0; }
}
.raw { background: #f8fafc; padding: 1rem; overflow-x: auto; }
.error { color: #dc2626; }
</style>


