<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { HIPOCONDRIA_LOGO_PATH } from './hipocondria-logo-path'


const canvasRef = ref < HTMLCanvasElement | null > (null)
const pos = ref({ x: 0, y: 0 })
const touching = ref(false)
const mobile = ref(false)

let raf = 0
let pixels: ImageData | null = null
let parts: Array<{
    x: number; y: number; bx: number; by: number; s: number;
    c: string; sc: string; life: number; hipo: boolean;
}> = []

/*--- constantes ---*/
const BG = '#0b0f19'
const BASE_PARTS = 15_000
const VERCEL_H = 19.7762
const HIPO_H = 60
const MAX_DIST = 240
const target = (cv: HTMLCanvasElement) =>
    Math.floor(BASE_PARTS * Math.sqrt((cv.width * cv.height) / (1920 * 1080)))

/*--- helpers ---*/
const size = (cv: HTMLCanvasElement) => {
    cv.width = innerWidth
    cv.height = innerHeight
    mobile.value = innerWidth < 768
}

function drawMask(ctx: CanvasRenderingContext2D, cv: HTMLCanvasElement) {
    const lh = mobile.value ? 60 : 120
    const vw = lh * (40 / VERCEL_H)
    const hw = lh * (400 / HIPO_H)
    const sp = mobile.value ? 40 : 80
    const tot = lh + sp + lh

    ctx.fillStyle = 'white'
    ctx.save()
    const MARGIN_X = 820
    const MARGIN_Y = 200
    ctx.translate(MARGIN_X, MARGIN_Y)

    /* Vercel */
    ctx.save()
    ctx.translate(-vw / 2, 0)
    const vScale = lh / VERCEL_H
    ctx.scale(vScale, vScale)
    ctx.beginPath()
    ctx.moveTo(23.3919, 0); ctx.lineTo(32.9188, 0)
    ctx.bezierCurveTo(36.7819, 0, 39.9136, 3.13165, 39.9136, 6.99475)
    ctx.lineTo(39.9136, 16.0805); ctx.lineTo(36.0006, 16.0805)
    ctx.lineTo(36.0006, 6.99475)
    ctx.bezierCurveTo(36.0006, 6.90167, 35.9969, 6.80925, 35.9898, 6.71766)
    ctx.lineTo(26.4628, 16.079)
    ctx.bezierCurveTo(26.4949, 16.08, 26.5272, 16.0805, 26.5595, 16.0805)
    ctx.lineTo(36.0006, 16.0805); ctx.lineTo(36.0006, 19.7762)
    ctx.lineTo(26.5595, 19.7762)
    ctx.bezierCurveTo(22.6964, 19.7762, 19.4788, 16.6139, 19.4788, 12.7508)
    ctx.lineTo(19.4788, 3.68923); ctx.lineTo(23.3919, 3.68923)
    ctx.lineTo(23.3919, 12.7508)
    ctx.bezierCurveTo(23.3919, 12.9253, 23.4054, 13.0977, 23.4316, 13.2668)
    ctx.lineTo(33.1682, 3.6995)
    ctx.bezierCurveTo(33.0861, 3.6927, 33.003, 3.68923, 32.9188, 3.68923)
    ctx.lineTo(23.3919, 3.68923); ctx.lineTo(23.3919, 0); ctx.closePath()
    ctx.moveTo(13.7688, 19.0956); ctx.lineTo(0, 3.68759); ctx.lineTo(5.53933, 3.68759)
    ctx.lineTo(13.6231, 12.7337); ctx.lineTo(13.6231, 3.68759); ctx.lineTo(17.7535, 3.68759)
    ctx.lineTo(17.7535, 17.5746)
    ctx.bezierCurveTo(17.7535, 19.6705, 15.1654, 20.6584, 13.7688, 19.0956)
    ctx.closePath()
    ctx.fill()
    ctx.restore()

    /* Hipocondria */
    ctx.save()
    ctx.translate(-hw / 2, lh + sp)
    const hScale = lh / HIPO_H
    ctx.scale(hScale, hScale)
    ctx.fill(new Path2D(HIPOCONDRIA_LOGO_PATH))
    ctx.restore()

    ctx.restore()
    pixels = ctx.getImageData(0, 0, cv.width, cv.height)
    ctx.clearRect(0, 0, cv.width, cv.height)
}

function makePart(cv: HTMLCanvasElement) {
    if (!pixels) return null
    const d = pixels.data
    for (let i = 0; i < 100; i++) {
        const x = (Math.random() * cv.width) | 0, y = (Math.random() * cv.height) | 0
        if (d[(y * cv.width + x) * 4 + 3] > 128) {
            const sp = mobile.value ? 40 : 80
            const hipo = y >= cv.height / 2 + sp / 2
            return {
                x, y, bx: x, by: y, s: Math.random() * 1.5 + 1,
                c: 'white', sc: hipo ? '#10B981' : '#00DCFF',
                life: Math.random() * 100 + 50, hipo
            }
        }
    }
    return null
}

function seed(cv: HTMLCanvasElement) {
    parts.length = 0
    for (let i = target(cv); i--;) {
        const p = makePart(cv); if (p) parts.push(p)
    }
}

function loop(ctx: CanvasRenderingContext2D, cv: HTMLCanvasElement) {
    ctx.clearRect(0, 0, cv.width, cv.height)
    ctx.fillStyle = BG; ctx.fillRect(0, 0, cv.width, cv.height)

    const { x: mx, y: my } = pos.value
    for (let i = 0; i < parts.length; i++) {
        const p = parts[i], dx = mx - p.x, dy = my - p.y, d = Math.hypot(dx, dy)
        if (d < MAX_DIST && (touching.value || !('ontouchstart' in window))) {
            const k = (MAX_DIST - d) / MAX_DIST, a = Math.atan2(dy, dx)
            p.x = p.bx - Math.cos(a) * k * 60
            p.y = p.by - Math.sin(a) * k * 60
            ctx.fillStyle = p.sc
        } else {
            p.x += (p.bx - p.x) * 0.1; p.y += (p.by - p.y) * 0.1
            ctx.fillStyle = 'white'
        }
        ctx.fillRect(p.x, p.y, p.s, p.s)
        if (--p.life <= 0) {
            const n = makePart(cv); n ? parts[i] = n : (parts.splice(i, 1), i--)
        }
    }
    while (parts.length < target(cv)) {
        const n = makePart(cv); if (n) parts.push(n)
    }
    raf = requestAnimationFrame(() => loop(ctx, cv))
}

/*--- ciclo de vida ---*/
let onResize: () => void, onMouse: (e: MouseEvent) => void, onTouch: (e: TouchEvent) => void,
    onTStart: () => void, onTEnd: () => void, onLeave: () => void

onMounted(() => {
    const cv = canvasRef.value!, ctx = cv.getContext('2d')!
    size(cv); drawMask(ctx, cv); seed(cv); loop(ctx, cv)

    const move = (x: number, y: number) => { pos.value = { x, y } }
    onResize = () => { cancelAnimationFrame(raf); size(cv); drawMask(ctx, cv); seed(cv); loop(ctx, cv) }
    onMouse = e => move(e.clientX, e.clientY)
    onTouch = e => { if (e.touches.length) { e.preventDefault(); move(e.touches[0].clientX, e.touches.clientY) } }
    onTStart = () => touching.value = true
    onTEnd = () => { touching.value = false; pos.value = { x: 0, y: 0 } }
    onLeave = () => { if (!('ontouchstart' in window)) pos.value = { x: 0, y: 0 } }

    window.addEventListener('resize', onResize)
    cv.addEventListener('mousemove', onMouse)
    cv.addEventListener('touchmove', onTouch, { passive: false })
    cv.addEventListener('mouseleave', onLeave)
    cv.addEventListener('touchstart', onTStart)
    cv.addEventListener('touchend', onTEnd)
})

onBeforeUnmount(() => {
    const cv = canvasRef.value
    if (cv) {
        cv.removeEventListener('mousemove', onMouse as any)
        cv.removeEventListener('touchmove', onTouch as any)
        cv.removeEventListener('mouseleave', onLeave as any)
        cv.removeEventListener('touchstart', onTStart as any)
        cv.removeEventListener('touchend', onTEnd as any)
    }
    window.removeEventListener('resize', onResize as any)
    cancelAnimationFrame(raf)
})

</script>

<template>
    <section class="relative w-full h-dvh flex items-center justify-center" style="background:#0b0f19">
        <h2 style="position: absolute" data-v-9852290e="" class="title">Contacto</h2>
        <p style="position: absolute; top: 300px" class="prose">Escríbenos a <a data-v-9852290e=""
                href="mailto:contacto@hipocondria.co">contacto@hipocondria.co</a></p>
        <canvas ref="canvasRef" class="w-full h-full absolute inset-0 touch-none" id="canvas"
            aria-label="Interactive particle effect with Vercel and hipocondria.co logos" />
    </section>
</template>

<style scoped>
@media screen and (max-width: 600px) {
  #canvas {
    display: none;
  }
}   
</style>