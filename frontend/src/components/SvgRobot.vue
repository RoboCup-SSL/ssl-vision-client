<script setup lang="ts">
import { computed } from 'vue'
import SvgText from '@/components/SvgText.vue'

const props = defineProps<{
  x: number
  y: number
  orientation: number
  id: number
  teamColor: 'YELLOW' | 'BLUE'
}>()

const center2Dribbler = 0.075
const radius = 0.09

const robotId = computed(() => {
  return String(props.id)
})

const botShape = computed(() => {
  const orient2CornerAngle = Math.acos(center2Dribbler / radius)
  const botRightX = props.x + Math.cos(-props.orientation + orient2CornerAngle) * radius
  const botRightY = -props.y + Math.sin(-props.orientation + orient2CornerAngle) * radius
  const botLeftX = props.x + Math.cos(-props.orientation - orient2CornerAngle) * radius
  const botLeftY = -props.y + Math.sin(-props.orientation - orient2CornerAngle) * radius

  return (
    `M ${botRightX} ${botRightY}` +
    ` A ${radius} ${radius} 0 1 1 ${botLeftX} ${botLeftY}` +
    ` L ${botRightX} ${botRightY}`
  )
})

const style = computed(() => {
  return {
    stroke: 'black',
    strokeWidth: 0.005,
    strokeOpacity: 1,
    fill: props.teamColor == 'YELLOW' ? 'yellow' : 'blue',
  }
})
</script>

<template>
  <path :d="botShape" :style="style" />
  <svg-text :x="x" :y="y" :text="robotId" :color="teamColor == 'YELLOW' ? 'black' : 'white'" />
</template>
