<script setup lang="ts">
import {computed, onBeforeUnmount, onMounted, ref} from 'vue'

const svg = ref<SVGElement>()
const zoom = ref(1.0)
const translation = ref({x: 0, y: 0})
const activeTranslation = ref({x: 0, y: 0})
const mouseDownPoint = ref<{
  x: number
  y: number
} | null>(null)

function onScroll(event: WheelEvent) {
  const x = (event.offsetX - translation.value.x) / zoom.value
  const y = (event.offsetY - translation.value.y) / zoom.value

  const newZoom = zoom.value - event.deltaY / 150
  if (newZoom < 1) {
    zoom.value = 1
    translation.value = {x: 0, y: 0}
  } else {
    const dz = newZoom - zoom.value
    translation.value.x -= x * dz
    translation.value.y -= y * dz
    zoom.value = newZoom
  }

  event.preventDefault()
}

function onMouseMove(event: MouseEvent) {
  if (mouseDownPoint.value !== null) {
    activeTranslation.value = {
      x: event.clientX - mouseDownPoint.value.x,
      y: event.clientY - mouseDownPoint.value.y,
    }
  }
}

function onMouseDown(event: MouseEvent) {
  mouseDownPoint.value = {x: event.clientX, y: event.clientY}
}

function onMouseUp() {
  if (mouseDownPoint.value !== null) {
    translation.value = {
      x: translation.value.x + activeTranslation.value.x,
      y: translation.value.y + activeTranslation.value.y,
    }
    activeTranslation.value = {x: 0, y: 0}
    mouseDownPoint.value = null
  }
}

function onClick(event: KeyboardEvent) {
  if (event.key === ' ') {
    zoom.value = 1
    translation.value = {x: 0, y: 0}
  }
}

const transform = computed(() => {
  const translationX = translation.value.x + activeTranslation.value.x
  const translationY = translation.value.y + activeTranslation.value.y
  return `translate(${translationX}, ${translationY}) scale(${zoom.value})`
})

onMounted(() => {
  document.addEventListener('keydown', onClick)
  svg.value?.addEventListener('wheel', onScroll)
  svg.value?.addEventListener('mousemove', onMouseMove)
  svg.value?.addEventListener('mousedown', onMouseDown)
  svg.value?.addEventListener('mouseup', onMouseUp)
})
onBeforeUnmount(() => {
  document.removeEventListener('keydown', onClick)
  svg.value?.removeEventListener('wheel', onScroll)
  svg.value?.removeEventListener('mousemove', onMouseMove)
  svg.value?.removeEventListener('mousedown', onMouseDown)
  svg.value?.removeEventListener('mouseup', onMouseUp)
})
</script>

<template>
  <svg width="100%" height="100%" ref="svg">
    <g :transform="transform">
      <slot/>
    </g>
  </svg>
</template>

<style scoped>

</style>
