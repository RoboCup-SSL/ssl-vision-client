<script setup lang="ts">
import SvgRobot from '@/components/SvgRobot.vue'
import SvgBall from '@/components/SvgBall.vue'
import type { SSL_DetectionFrame } from '@/proto/vision/ssl_vision_detection_pb.ts'

defineProps<{
  detectionFrame: SSL_DetectionFrame
}>()
</script>

<template>
  <SvgBall
    v-for="(s, i) in detectionFrame.balls"
    :key="'ball-' + i"
    :x="s.x / 1000"
    :y="s.y / 1000"
    :height="0"
  />
  <SvgRobot
    v-for="(s, i) in detectionFrame.robotsYellow"
    :key="'robot-yellow-' + i"
    :x="s.x / 1000"
    :y="s.y / 1000"
    :orientation="s.orientation"
    :id="s.robotId"
    :team-color="'YELLOW'"
  />
  <SvgRobot
    v-for="(s, i) in detectionFrame.robotsBlue"
    :key="'robot-yellow-' + i"
    :x="s.x / 1000"
    :y="s.y / 1000"
    :orientation="s.orientation"
    :id="s.robotId"
    :team-color="'BLUE'"
  />
</template>
