<script setup lang="ts">
import { computed, type CSSProperties } from 'vue'

const props = defineProps<{
  x: number
  y: number
  height: number
}>()

const radius = computed(() => {
  return 0.021 * (1 + 0.01 * props.height)
})

const style = computed((): CSSProperties => {
  return {
    fill: 'orange',
    fillOpacity: 1,
  }
})
const highlightStyle = computed((): CSSProperties => {
  return {
    stroke: 'orange',
    strokeWidth: 0.01,
    fill: 'none'
  }
})
</script>

<template>
  <circle :cx="x" :cy="y" :r="radius" :style="style" />

  <!-- ball highlighter -->
  <circle :cx="x" :cy="y" :r="0.5" :style="highlightStyle" />

</template>
