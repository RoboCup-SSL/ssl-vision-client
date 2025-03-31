<script setup lang="ts">
import { inject } from 'vue'
import type { VisionApi } from '@/providers/backend/VisionApi'

const props = defineProps<{
  sources: { string: string }
  activeSource: string
}>()

const visionApi = inject<VisionApi>('vision-api')
function updateSource(newSource: string) {
  visionApi?.Send({ activeSourceId: newSource })

  const newUrl =
    window.location.protocol +
    '//' +
    window.location.host +
    window.location.pathname +
    '?sourceId=' +
    newSource
  window.history.pushState({ path: newUrl }, '', newUrl)
}
</script>

<template>
  <div id="source-selector">
    Sources:
    <template v-for="(sourceName, sourceId) in props.sources" :key="'input-' + sourceId">
      <input
        type="radio"
        :id="sourceId"
        name="source"
        :value="sourceId"
        :checked="sourceId === props.activeSource"
        @click="updateSource(sourceId)"
      />
      <label :for="sourceId">
        {{ sourceName }}
      </label>
    </template>
  </div>
</template>

<style scoped>
#source-selector {
  align-content: center;
  width: 100%;
  bottom: 0.1em;
  position: absolute;
}
</style>
