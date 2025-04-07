<script setup lang="ts">
import SvgRobot from '@/components/SvgRobot.vue'
import SvgBall from '@/components/SvgBall.vue'
import type { TrackedFrame } from '@/proto/tracked/ssl_vision_detection_tracked_pb.ts'
import { Team } from '@/proto/gc/ssl_gc_common_pb.ts'

defineProps<{
  trackedFrame: TrackedFrame
}>()
</script>

<template>
  <SvgBall
    v-for="(s, i) in trackedFrame.balls"
    :key="'tracked-ball-' + i"
    :x="s.pos!.x"
    :y="s.pos!.y"
    :height="0"
  />
  <SvgRobot
    v-for="(s, i) in trackedFrame.robots"
    :key="'tracked-robot-' + i"
    :x="s.pos!.x"
    :y="s.pos!.y"
    :orientation="s.orientation"
    :id="s.robotId!.id"
    :team-color="s.robotId!.team === Team.YELLOW ? 'YELLOW' : 'BLUE'"
  />

</template>
