<script setup lang="ts">
import {computed, type CSSProperties, onBeforeUnmount, onMounted, ref} from "vue"
import type {Field, PathSegment, Point, ShapeStyle} from "@/field";
import ZoomPanSvg from "@/components/ZoomPanSvg.vue";

const props = defineProps<{
  field: Field,
}>()

const wrapperDiv = ref<HTMLDivElement>()
const rotateField = ref(false)

const fieldWithBoundary = computed(() => {
  return {
    minX: -(props.field.fieldLength / 2 + props.field.boundaryWidth),
    minY: -(props.field.fieldWidth / 2 + props.field.boundaryWidth),
    width: props.field.fieldLength + props.field.boundaryWidth * 2,
    height: props.field.fieldWidth + props.field.boundaryWidth * 2,
  }
})

function pathFromD(pd: PathSegment[]) {
  let d = '';
  for (let s of pd) {
    d += s.type;
    for (let a of s.args) {
      d += ' ' + a
    }
  }
  return d;
}

function textTransform(p: Point) {
  if (rotateField.value) {
    return 'rotate(-90,' + p.x + ',' + p.y + ')'
  }
  return ''
}

function shapeStyle(style: ShapeStyle): CSSProperties {
  const defStyle = {
    stroke: 'white',
    strokeWidth: 10,
    fill: "",
    fillOpacity: 1,
    textAnchor: "middle",
    dominantBaseline: "central",
    font: "bold 7em sans-serif",
  }
  return {...defStyle, ...style} as CSSProperties
}

const viewBox = computed(() => {
  if (rotateField.value) {
    return `${fieldWithBoundary.value.minY} ${fieldWithBoundary.value.minX} ${fieldWithBoundary.value.height} ${fieldWithBoundary.value.width}`
  }
  return `${fieldWithBoundary.value.minX} ${fieldWithBoundary.value.minY} ${fieldWithBoundary.value.width} ${fieldWithBoundary.value.height}`
})

const onWindowResized = () => {
  if (!wrapperDiv.value) {
    return
  }
  const wl = props.field.fieldWidth / props.field.fieldLength
  const lw = props.field.fieldLength / props.field.fieldWidth
  const meanFieldRatio = (wl + lw) / 2;
  const canvasRatio = wrapperDiv.value.clientHeight / wrapperDiv.value.clientWidth
  rotateField.value = meanFieldRatio < canvasRatio
}

const transform = computed(() => {
  if (rotateField.value) {
    return `rotate(90)`
  }
  return ''
})

onMounted(() => {
  window.addEventListener('resize', onWindowResized)
  onWindowResized()
})
onBeforeUnmount(() => {
  window.removeEventListener('resize', onWindowResized)
})
</script>

<template>
  <div class="wrapper" ref="wrapperDiv">
    <ZoomPanSvg>
      <svg :viewBox="viewBox">
        <g :transform="transform">
          <!-- draw field background -->
          <rect
            :x="fieldWithBoundary.minX"
            :y="fieldWithBoundary.minY"
            :width="fieldWithBoundary.width"
            :height="fieldWithBoundary.height"
            fill="green"
          />

          <template v-for="(s,i) of props.field.shapes">
            <line
              v-if="s.line"
              :key="'shape-' + i"
              :x1="s.line.p1.x"
              :y1="s.line.p1.y"
              :x2="s.line.p2.x"
              :y2="s.line.p2.y"
              :style="shapeStyle(s.line)"
            />

            <circle
              v-if="s.circle"
              :key="'shape-' + i"
              :cx="s.circle.center.x"
              :cy="s.circle.center.y"
              :r="s.circle.radius"
              :style="shapeStyle(s.circle)"
            />

            <path
              v-if="s.path"
              :key="'shape-' + i"
              :d="pathFromD(s.path.d)"
              :style="shapeStyle(s.path)"
            />

            <text
              v-if="s.text"
              :key="'shape-' + i"
              :x="s.text.p.x"
              :y="s.text.p.y"
              :transform="textTransform(s.text.p)"
              :style="shapeStyle(s.text)"
            >
              {{ s.text.text }}
            </text>
          </template>
        </g>
      </svg>
    </ZoomPanSvg>
  </div>
</template>

<style scoped>
.wrapper {
  width: 100%;
  height: 100%;
}
</style>
