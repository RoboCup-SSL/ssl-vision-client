<script setup lang="ts">
import {computed, onBeforeUnmount, onMounted, ref} from "vue";
import type {Field, PathSegment, Point} from "@/field";

const props = defineProps<{
  field: Field,
}>()

const canvas = ref<SVGElement>()
const canvasWidth = ref(1)
const canvasHeight = ref(1)
const zoom = ref(1.0)
const translation = ref({x: 0, y: 0})
const activeTranslation = ref({x: 0, y: 0})
const mouseDownPoint = ref<Point | null>(null)

const viewBox = computed(() => {
  if (rotateField.value) {
    return (-(props.field.fieldWidth / 2 + props.field.boundaryWidth))
      + ' ' + (-(props.field.fieldLength / 2 + props.field.boundaryWidth))
      + ' ' + (props.field.fieldWidth + props.field.boundaryWidth * 2)
      + ' ' + (props.field.fieldLength + props.field.boundaryWidth * 2)
  }
  return (-(props.field.fieldLength / 2 + props.field.boundaryWidth))
    + ' ' + (-(props.field.fieldWidth / 2 + props.field.boundaryWidth))
    + ' ' + (props.field.fieldLength + props.field.boundaryWidth * 2)
    + ' ' + (props.field.fieldWidth + props.field.boundaryWidth * 2)
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

function updateCanvasWidth() {
  canvasWidth.value = canvas.value?.clientWidth || 0
}

function updateCanvasHeight() {
  canvasHeight.value = canvas.value?.clientHeight || 0
}

function textTransform(p: Point) {
  if (rotateField.value) {
    return 'rotate(-90,' + p.x + ',' + p.y + ')'
  }
  return ''
}

function onScroll(event: WheelEvent) {
  let newZoom = zoom.value - event.deltaY / 300
  if (newZoom < 1) {
    zoom.value = 1;
  } else {
    zoom.value = newZoom
  }
}

function onMouseMove(event: MouseEvent) {
  if (mouseDownPoint.value !== null) {
    activeTranslation.value = {
      x: event.clientX - mouseDownPoint.value.x,
      y: event.clientY - mouseDownPoint.value.y
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
    };
    activeTranslation.value = {x: 0, y: 0}
    mouseDownPoint.value = null
  }
}

function onClick(event: KeyboardEvent) {
  if (event.key === " ") {
    zoom.value = 1;
    translation.value = {x: 0, y: 0}
  }
}

const meanFieldRatio = computed(() => {
  let wl = props.field.fieldWidth / props.field.fieldLength
  let lw = props.field.fieldLength / props.field.fieldWidth
  return (wl + lw) / 2;
})
const rotateField = computed(() => {
  return meanFieldRatio.value < canvasRatio.value
})
const fieldToPixelRatioX = computed(() => {
  if (rotateField.value) {
    return (props.field.fieldWidth + props.field.boundaryWidth * 2) / canvasWidth.value
  }
  return (props.field.fieldLength + props.field.boundaryWidth * 2) / canvasWidth.value
})
const fieldToPixelRatioY = computed(() => {
  if (rotateField.value) {
    return (props.field.fieldLength + props.field.boundaryWidth * 2) / canvasHeight.value
  }
  return (props.field.fieldWidth + props.field.boundaryWidth * 2) / canvasHeight.value
})
const fieldTranslationX = computed(() => {
  return (translation.value.x + activeTranslation.value.x) * fieldToPixelRatioX.value
})
const fieldTranslationY = computed(() => {
  return (translation.value.y + activeTranslation.value.y) * fieldToPixelRatioY.value
})
const canvasRatio = computed(() => {
  return canvasHeight.value / canvasWidth.value
})
const getFieldTransformation = computed(() => {
  let scale = 'scale(' + zoom.value + ') '
  let transform = 'translate(' + fieldTranslationX.value + ',' + fieldTranslationY.value + ') '
  if (rotateField.value) {
    return transform + scale + 'rotate(90)'
  }
  return transform + scale
})

const defStyle = {
  strokeWidth: 10,
  stroke: 'white',
  fillOpacity: 1,
}
const defFontStyle = {
  strokeWidth: 0,
  textAnchor: "middle",
  dominantBaseline: "central",
  font: "bold 7em sans-serif",
}

onMounted(() => {
  // $nextTick(function () {
  window.addEventListener('resize', updateCanvasWidth)
  window.addEventListener('resize', updateCanvasHeight)
  canvas.value?.addEventListener("wheel", onScroll)
  document.addEventListener('keydown', onClick)

  updateCanvasWidth()
  updateCanvasHeight()
})
onBeforeUnmount(() => {
  window.removeEventListener('resize', updateCanvasWidth)
  window.removeEventListener('resize', updateCanvasHeight)
  canvas.value?.removeEventListener("wheel", onScroll)
  document.removeEventListener('keydown', onClick)
})

</script>

<template>
  <svg
    ref="canvas"
    @mousemove="onMouseMove"
    @mousedown="onMouseDown"
    @mouseup="onMouseUp"
    :viewBox="viewBox"
  >
    <!-- rotate field -->
    <g :transform="getFieldTransformation">

      <!-- draw field background -->
      <rect
        :x="-(props.field.fieldLength/2+props.field.boundaryWidth)"
        :y="-(props.field.fieldWidth/2+props.field.boundaryWidth)"
        :width="props.field.fieldLength+props.field.boundaryWidth*2"
        :height="props.field.fieldWidth+props.field.boundaryWidth*2"
        ref="background"
      />

      <template v-for="(s,i) of props.field.shapes">

        <line
          v-if="s.line"
          :key="'shape-' + i"
          :x1="s.line.p1.x"
          :y1="s.line.p1.y"
          :x2="s.line.p2.x"
          :y2="s.line.p2.y"
          :style="[defStyle, s.line]"
        />

        <circle
          v-if="s.circle"
          :key="'shape-' + i"
          :cx="s.circle.center.x"
          :cy="s.circle.center.y"
          :r="s.circle.radius"
          :style="[defStyle, s.circle]"
        />

        <path
          v-if="s.path"
          :key="'shape-' + i"
          :d="pathFromD(s.path.d)"
          :style="[defStyle, s.path]"
        />

        <text
          v-if="s.text"
          :key="'shape-' + i"
          :x="s.text.p.x"
          :y="s.text.p.y"
          :transform="textTransform(s.text.p)"
          :style="[defStyle, defFontStyle, s.text]"
        >
          {{ s.text.text }}
        </text>
      </template>
    </g>
  </svg>
</template>

<style scoped>
svg {
  width: 100%;
  height: 100%;
  display: table-row;
  fill: green;
}
</style>
