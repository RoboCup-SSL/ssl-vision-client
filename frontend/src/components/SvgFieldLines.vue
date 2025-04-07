<script setup lang="ts">
import type { SSL_GeometryFieldSize } from '@/proto/vision/ssl_vision_geometry_pb.ts'
import { computed } from 'vue'

const props = defineProps<{
  field: SSL_GeometryFieldSize
}>()

const fieldMaxX = computed(() => {
  return props.field.fieldLength / 2000
})
const fieldMaxY = computed(() => {
  return props.field.fieldWidth / 2000
})
const penaltyAreaMinX = computed(() => {
  return fieldMaxX.value - props.field.penaltyAreaDepth / 1000
})
const penaltyAreaMaxY = computed(() => {
  return props.field.penaltyAreaWidth / 2000
})

const lineStyle = computed(() => {
  return {
    stroke: 'white',
    strokeWidth: props.field.lineThickness / 1000,
    strokeOpacity: 1,
    fill: 'none',
  }
})
</script>

<template>
  <line :x1="fieldMaxX" :y1="fieldMaxY" :x2="fieldMaxX" :y2="-fieldMaxY" :style="lineStyle" />
  <line :x1="-fieldMaxX" :y1="fieldMaxY" :x2="-fieldMaxX" :y2="-fieldMaxY" :style="lineStyle" />
  <line :x1="fieldMaxX" :y1="fieldMaxY" :x2="-fieldMaxX" :y2="fieldMaxY" :style="lineStyle" />
  <line :x1="fieldMaxX" :y1="-fieldMaxY" :x2="-fieldMaxX" :y2="-fieldMaxY" :style="lineStyle" />

  <line :x1="0" :y1="-fieldMaxY" :x2="0" :y2="fieldMaxY" :style="lineStyle" />

  <circle :cx="0" :cy="0" :r="field.centerCircleRadius / 1000" :style="lineStyle" />

  <line
    :x1="penaltyAreaMinX"
    :y1="penaltyAreaMaxY"
    :x2="penaltyAreaMinX"
    :y2="-penaltyAreaMaxY"
    :style="lineStyle"
  />
  <line
    :x1="penaltyAreaMinX"
    :y1="penaltyAreaMaxY"
    :x2="fieldMaxX"
    :y2="penaltyAreaMaxY"
    :style="lineStyle"
  />
  <line
    :x1="penaltyAreaMinX"
    :y1="-penaltyAreaMaxY"
    :x2="fieldMaxX"
    :y2="-penaltyAreaMaxY"
    :style="lineStyle"
  />

  <line
    :x1="-penaltyAreaMinX"
    :y1="penaltyAreaMaxY"
    :x2="-penaltyAreaMinX"
    :y2="-penaltyAreaMaxY"
    :style="lineStyle"
  />
  <line
    :x1="-penaltyAreaMinX"
    :y1="penaltyAreaMaxY"
    :x2="-fieldMaxX"
    :y2="penaltyAreaMaxY"
    :style="lineStyle"
  />
  <line
    :x1="-penaltyAreaMinX"
    :y1="-penaltyAreaMaxY"
    :x2="-fieldMaxX"
    :y2="-penaltyAreaMaxY"
    :style="lineStyle"
  />

</template>
