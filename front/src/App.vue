<script setup>
</script>

<template>
  <div class="layout">
    <header class="site-header">
      <div class="container">
        <h1 class="site-brand"><router-link to="/">Hipocondria.co</router-link></h1>
        <nav class="site-nav" v-if="$route.name !== 'article'">
          <router-link to="/">Artículos</router-link>
          <router-link to="/contacto" id="contactM">Contacto</router-link>
          <router-link to="/contactoD" id="contactD">Contacto</router-link>
        </nav>
      </div>
    </header>
    <main class="container">
      <router-view />
    </main>
    <footer class="site-footer">
      <div class="container">© {{ new Date().getFullYear() }} hipocondria.co</div>
    </footer>
  </div>
  
</template>

<style scoped>
.layout { min-height: 100vh; display: flex; flex-direction: column; }
.container { width: 100%; max-width: 1280px; margin: 0 auto; padding: 1.5rem; display: flex; align-items: center; justify-content: space-between; }
.site-header .container { flex-wrap: wrap; gap: .5rem 1rem; }
.site-brand { margin: 0; font-size: 6rem; line-height: 1.05; }
.site-brand a { color: inherit; box-shadow: none; }
.site-nav { display: flex; flex-wrap: wrap; align-items: center; gap: .75rem; }

@media (max-width: 640px) { .site-brand { font-size: 2.2rem; } #contactD { display: none; } }
@media (min-width: 641px) { #contactM { display: none; } }
@media (min-width: 1536px) { .container { max-width: 90vw; } .site-brand { font-size: 7.5rem; } }
main.container { flex: 1; display: block; }
</style>
