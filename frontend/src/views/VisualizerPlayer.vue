<script setup lang="ts">
import FieldVisualizer from '@/components/FieldVisualizer.vue'
import SvgVision from '@/components/SvgVision.vue'
import SvgReferee from '@/components/SvgReferee.vue'
import SvgTracked from '@/components/SvgTracked.vue'
import SourceSelector from '@/components/SourceSelector.vue'
import LogUrlInput from '@/components/LogUrlInput.vue'
import { useLogUrl } from '@/composables/logUrl.ts'
import { useLogManifest } from '@/composables/logManifest.ts'
import { useLogPlayerSources } from '@/composables/logPlayerSources.ts'
import { useLogPlayerIndices } from '@/composables/logPlayerIndices.ts'
import { useLogPlayerMessages } from '@/composables/logPlayerMessages.ts'
import { useLogPlayerPosition } from '@/composables/logPlayerPosition.ts'
import { refThrottled, onKeyStroke } from '@vueuse/core'

const { logUrlInput, urlError, validLogUrl } = useLogUrl()
const manifest = useLogManifest(validLogUrl)
const { activeSource, sources } = useLogPlayerSources(manifest.manifestEntriesOfType('tracker'))

const { visionDetectionIndices, primaryIndex, globalMinTimestamp, globalMaxTimestamp } =
  useLogPlayerIndices(validLogUrl, manifest, activeSource)

const {
  currentTimestamp,
  sliderValue,
  sliderMax,
  stepForward,
  stepBackward,
  skipTime,
  playing,
  togglePlay,
} = useLogPlayerPosition(globalMinTimestamp, globalMaxTimestamp, primaryIndex, validLogUrl)

onKeyStroke(' ', (e) => {
  e.preventDefault()
  togglePlay()
})

onKeyStroke('ArrowLeft', (e) => {
  if (e.ctrlKey || e.metaKey) {
    e.preventDefault()
    stepBackward()
  } else if (e.shiftKey) {
    e.preventDefault()
    skipTime(-30)
  } else {
    e.preventDefault()
    skipTime(-10)
  }
})

onKeyStroke('ArrowRight', (e) => {
  if (e.ctrlKey || e.metaKey) {
    e.preventDefault()
    stepForward()
  } else if (e.shiftKey) {
    e.preventDefault()
    skipTime(30)
  } else {
    e.preventDefault()
    skipTime(10)
  }
})

// Debounce timestamp for data fetching to avoid flooding the server
// with concurrent HTTP range requests when sliding fast
const debouncedTimestamp = refThrottled(currentTimestamp, 50)

const { detectionFrames, field, referee, trackedFrame } = useLogPlayerMessages(
  validLogUrl,
  manifest,
  visionDetectionIndices,
  debouncedTimestamp,
  activeSource,
)
</script>

<template>
  <div id="container">
    <div id="control">
      <LogUrlInput v-model="logUrlInput" :url-error="urlError" />
      <div v-if="validLogUrl && sliderMax === 0">Loading log file...</div>
    </div>
    <div id="content">
      <FieldVisualizer :field="field">
        <SvgVision
          v-for="(frame, i) in detectionFrames"
          :key="'cam-' + i"
          :detection-frame="frame"
        />
        <SvgReferee v-if="referee" :field="field" :referee="referee" />
        <SvgTracked v-if="trackedFrame" :tracked-frame="trackedFrame" />
      </FieldVisualizer>
      <SourceSelector :sources="sources" v-model="activeSource" />
    </div>
    <div v-if="sliderMax > 0" id="slider">
      <div id="slider-row">
        <input
          type="range"
          min="0"
          :max="sliderMax"
          :value="sliderValue"
          @input="sliderValue = Number(($event.target as HTMLInputElement).value)"
        />
        <div class="skip-buttons">
          <button title="Back 30s (Shift+Left)" @click="skipTime(-30)">&laquo;30s</button>
          <button title="Back 10s (Left)" @click="skipTime(-10)">&lsaquo;10s</button>
          <button title="Previous frame (Ctrl+Left)" @click="stepBackward">&lsaquo;|</button>
          <button title="Play/Pause (Space)" @click="togglePlay">{{ playing ? '⏸' : '▶' }}</button>
          <button title="Next frame (Ctrl+Right)" @click="stepForward">|&rsaquo;</button>
          <button title="Forward 10s (Right)" @click="skipTime(10)">10s&rsaquo;</button>
          <button title="Forward 30s (Shift+Right)" @click="skipTime(30)">30s&raquo;</button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
#container {
  display: flex;
  flex-direction: column;
  width: 100vw;
  height: 100vh;
  overflow: hidden;
  padding: 8px;
  box-sizing: border-box;
}

#control {
  width: 100%;
  flex-shrink: 0;
  margin: 6px 0;
}

#content {
  width: 100%;
  flex: 1;
  overflow: hidden;
  min-height: 0;
}

#slider {
  width: 100%;
  flex-shrink: 0;
  margin: 10px 0;
}

#slider-row {
  display: flex;
  align-items: center;
  gap: 8px;
}

#slider-row input[type='range'] {
  flex: 1;
  margin: 0 4px;
}

.skip-buttons {
  display: flex;
  gap: 4px;
  flex-shrink: 0;
}

.skip-buttons button {
  padding: 2px 6px;
  font-size: 0.8em;
  cursor: pointer;
  white-space: nowrap;
}
</style>
