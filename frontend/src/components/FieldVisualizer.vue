<script setup lang="ts">
import { computed, onBeforeUnmount, onMounted, provide, ref } from 'vue'
import SvgZoomPan from '@/components/SvgZoomPan.vue'
import type { SSL_GeometryFieldSize } from '@/proto/vision/ssl_vision_geometry_pb.ts'
import SvgFieldLines from '@/components/SvgFieldLines.vue'

const props = defineProps<{
  field: SSL_GeometryFieldSize
}>()

const wrapperDiv = ref<HTMLDivElement>()
const rotateField = ref(false)
provide('rotate-field', rotateField)

const fieldWithBoundary = computed(() => {
  return {
    minX: -(props.field.fieldLength / 2 + props.field.boundaryWidth) / 1000,
    minY: -(props.field.fieldWidth / 2 + props.field.boundaryWidth) / 1000,
    width: (props.field.fieldLength + props.field.boundaryWidth * 2) / 1000,
    height: (props.field.fieldWidth + props.field.boundaryWidth * 2) / 1000,
  }
})

const viewBox = computed(() => {
  if (rotateField.value) {
    return `${fieldWithBoundary.value.minY} ${fieldWithBoundary.value.minX} ${fieldWithBoundary.value.height} ${fieldWithBoundary.value.width}`
  }
  return `${fieldWithBoundary.value.minX} ${fieldWithBoundary.value.minY} ${fieldWithBoundary.value.width} ${fieldWithBoundary.value.height}`
})

const transform = computed(() => {
  if (rotateField.value) {
    return `rotate(90)`
  }
  return ''
})

const wrapperHeightObserver = new ResizeObserver(() => {
  if (!wrapperDiv.value) {
    return
  }
  const wl = props.field.fieldWidth / props.field.fieldLength
  const lw = props.field.fieldLength / props.field.fieldWidth
  const meanFieldRatio = (wl + lw) / 2
  const canvasRatio = wrapperDiv.value.clientHeight / wrapperDiv.value.clientWidth
  rotateField.value = meanFieldRatio < canvasRatio
})
onMounted(() => {
  wrapperHeightObserver.observe(wrapperDiv.value!)
})

onBeforeUnmount(() => {
  wrapperHeightObserver.unobserve(wrapperDiv.value!)
})
</script>

<template>
  <div class="wrapper" ref="wrapperDiv">
    <SvgZoomPan>
      <svg :viewBox="viewBox" width="100%" height="100%">
        <g :transform="transform">
          <!-- draw field background -->
          <rect
            :x="fieldWithBoundary.minX"
            :y="fieldWithBoundary.minY"
            :width="fieldWithBoundary.width"
            :height="fieldWithBoundary.height"
            fill="green"
          />

          <svg-field-lines :field="field" />
          <slot />
        </g>
      </svg>
    </SvgZoomPan>
  </div>
</template>

<style scoped>
.wrapper {
  width: 100%;
  height: 100%;
  overflow: hidden;
}
</style>
