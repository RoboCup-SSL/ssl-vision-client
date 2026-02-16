<script setup lang="ts">
import FieldVisualizer from '@/components/FieldVisualizer.vue'
import SvgVision from '@/components/SvgVision.vue'
import SvgReferee from '@/components/SvgReferee.vue'
import SvgTracked from '@/components/SvgTracked.vue'
import SourceSelector from '@/components/SourceSelector.vue'
import { computed, ref } from 'vue'
import {
  useTrackedFrame,
  useTrackedSources,
  useVisionDetection,
  useVisionGeometry,
} from '@/composables/vision.ts'
import { useReferee } from '@/composables/referee.ts'

const activeSource = ref('vision')
const { field } = useVisionGeometry()
const { detectionFrame } = useVisionDetection(activeSource)
const { referee } = useReferee()
const { trackedFrame } = useTrackedFrame(activeSource)
const { trackerSources } = useTrackedSources()

const sources = computed(() => {
  return {
    vision: 'vision',
    ...trackerSources.value,
  }
})
</script>

<template>
  <FieldVisualizer :field="field">
    <SvgVision v-if="detectionFrame" :detection-frame="detectionFrame" />
    <SvgReferee v-if="referee" :field="field" :referee="referee" />
    <SvgTracked v-if="trackedFrame" :tracked-frame="trackedFrame" />
  </FieldVisualizer>
  <source-selector :sources="sources" v-model="activeSource" />
</template>
