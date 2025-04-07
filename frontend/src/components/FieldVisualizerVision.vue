<script setup lang="ts">
import FieldVisualizer from '@/components/FieldVisualizer.vue'
import SvgVision from '@/components/SvgVision.vue'
import SvgReferee from '@/components/SvgReferee.vue'
import SvgTracked from '@/components/SvgTracked.vue'
import SourceSelector from '@/components/SourceSelector.vue'
import { computed, ref, watch } from 'vue'
import { useVisionDetection, useVisionGeometry } from '@/providers/vision'
import { useReferee } from '@/providers/referee'
import { useTrackedFrame, useTrackedSources } from '@/providers/tracker'

const { field } = useVisionGeometry()
const { detectionFrame } = useVisionDetection()
const { referee } = useReferee()
const { trackedFrame, updateSourceId } = useTrackedFrame()
const { trackerSources } = useTrackedSources()

const sources = computed(() => {
  return {
    vision: 'vision',
    ...trackerSources.value,
  }
})

const activeSource = ref('vision')

watch(activeSource, () => {
  if (activeSource.value !== 'vision') {
    updateSourceId(activeSource.value)
  } else {
    updateSourceId('')
  }
})
</script>

<template>
  <FieldVisualizer :field="field">
    <SvgVision v-if="activeSource === 'vision'" :detection-frame="detectionFrame" />
    <SvgReferee :field="field" :referee="referee" />
    <SvgTracked :tracked-frame="trackedFrame" />
  </FieldVisualizer>
  <source-selector :sources="sources" v-model="activeSource" />
</template>
