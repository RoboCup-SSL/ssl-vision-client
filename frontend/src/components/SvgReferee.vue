<script setup lang="ts">
import SvgGoal from '@/components/SvgGoal.vue'
import type { SSL_GeometryFieldSize } from '@/proto/vision/ssl_vision_geometry_pb.ts'
import type { Referee } from '@/proto/gc/ssl_gc_referee_message_pb.ts'
import {computed, type CSSProperties} from "vue";

defineProps<{
  field: SSL_GeometryFieldSize
  referee: Referee
}>()

const placementPosStyle = computed((): CSSProperties => {
  return {
    stroke: 'black',
    strokeWidth: 0.01,
    fill: 'none'
  }
})
</script>

<template>
  <svg-goal :field="field" :team-color="'YELLOW'" :positive-half="!referee.blueTeamOnPositiveHalf" />
  <svg-goal :field="field" :team-color="'BLUE'" :positive-half="referee.blueTeamOnPositiveHalf" />

  <circle
    v-if="referee.designatedPosition"
    :cx="referee.designatedPosition.x / 1000"
    :cy="-referee.designatedPosition.y / 1000"
    :r="0.15"
    :style="placementPosStyle" />

</template>
