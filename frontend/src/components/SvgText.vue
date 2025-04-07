<script setup lang="ts">
import { computed, type CSSProperties, inject, type Ref } from 'vue'

const props = defineProps<{
  x: number
  y: number
  text: string
  color?: string
}>()

const rotateField = inject<Ref<boolean>>('rotate-field')!

function textTransform(x: number, y: number) {
  if (rotateField.value) {
    return 'rotate(-90,' + x + ',' + y + ')'
  }
  return ''
}

const style = computed((): CSSProperties => {
  return {
    strokeWidth: 0,
    fill: props.color || 'white',
    fillOpacity: 1,
    textAnchor: 'middle',
    dominantBaseline: 'central',
    font: 'bold 0.007em sans-serif',
  }
})
</script>

<template>
  <text :x="x" :y="y" :transform="textTransform(x, y)" :style="style">
    {{ text }}
  </text>
</template>
