<script setup>
</script>

<template>
  <div class="gears" role="status" aria-live="polite" aria-label="Cargando">
    <svg class="svg" viewBox="0 0 120 60" xmlns="http://www.w3.org/2000/svg" focusable="false" aria-hidden="true">
      <g class="gear g1">
        <circle cx="30" cy="30" r="14" class="ring" />
        <circle cx="30" cy="30" r="6" class="hub" />
        <rect x="27" y="6" width="6" height="10" rx="1" ry="1" transform="rotate(0 30 30)"/>
        <rect x="27" y="6" width="6" height="10" rx="1" ry="1" transform="rotate(45 30 30)"/>
        <rect x="27" y="6" width="6" height="10" rx="1" ry="1" transform="rotate(90 30 30)"/>
        <rect x="27" y="6" width="6" height="10" rx="1" ry="1" transform="rotate(135 30 30)"/>
        <rect x="27" y="6" width="6" height="10" rx="1" ry="1" transform="rotate(180 30 30)"/>
        <rect x="27" y="6" width="6" height="10" rx="1" ry="1" transform="rotate(225 30 30)"/>
        <rect x="27" y="6" width="6" height="10" rx="1" ry="1" transform="rotate(270 30 30)"/>
        <rect x="27" y="6" width="6" height="10" rx="1" ry="1" transform="rotate(315 30 30)"/>
      </g>
      <g class="gear g2">
        <circle cx="78" cy="30" r="11" class="ring" />
        <circle cx="78" cy="30" r="5" class="hub" />
        <rect x="75" y="8" width="6" height="8" rx="1" ry="1" transform="rotate(0 78 30)"/>
        <rect x="75" y="8" width="6" height="8" rx="1" ry="1" transform="rotate(40 78 30)"/>
        <rect x="75" y="8" width="6" height="8" rx="1" ry="1" transform="rotate(80 78 30)"/>
        <rect x="75" y="8" width="6" height="8" rx="1" ry="1" transform="rotate(120 78 30)"/>
        <rect x="75" y="8" width="6" height="8" rx="1" ry="1" transform="rotate(160 78 30)"/>
        <rect x="75" y="8" width="6" height="8" rx="1" ry="1" transform="rotate(200 78 30)"/>
        <rect x="75" y="8" width="6" height="8" rx="1" ry="1" transform="rotate(240 78 30)"/>
        <rect x="75" y="8" width="6" height="8" rx="1" ry="1" transform="rotate(280 78 30)"/>
        <rect x="75" y="8" width="6" height="8" rx="1" ry="1" transform="rotate(320 78 30)"/>
      </g>
    </svg>
    <span class="sr-only">Cargando…</span>
  </div>
</template>

<style scoped>
.gears { display: inline-flex; align-items: center; justify-content: center; padding: .5rem; color: var(--color-accent); }
.svg { width: 120px; height: 60px; display: block; }
.gear { stroke: currentColor; fill: none; stroke-width: 2.2; transform-box: fill-box; transform-origin: center; }
.g1 { animation: spin 2.4s linear infinite; }
.g2 { animation: spin 3s linear infinite reverse; }
.ring { opacity: .9; }
.hub { fill: none; stroke-width: 0; }

@keyframes spin { to { transform: rotate(360deg); } }

@media (prefers-reduced-motion: reduce) { .g1, .g2 { animation: none; } }

.sr-only { position: absolute; width: 1px; height: 1px; padding: 0; margin: -1px; overflow: hidden; clip: rect(0, 0, 0, 0); white-space: nowrap; border: 0; }
</style>


