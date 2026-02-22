<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { fromBinary } from '@bufbuild/protobuf'
import { RefereeSchema } from '@/proto/gc/ssl_gc_referee_message_pb.ts'
import { SSL_WrapperPacketSchema } from '@/proto/vision/ssl_vision_wrapper_pb.ts'
import { TrackerWrapperPacketSchema } from '@/proto/tracked/ssl_vision_wrapper_tracked_pb.ts'
import { useLogFileIndex, useLogFileMessageAtTimestamp } from '@/composables/logfile.ts'
import FieldVisualizer from '@/components/FieldVisualizer.vue'
import SvgVision from '@/components/SvgVision.vue'
import SvgReferee from '@/components/SvgReferee.vue'
import SvgTracked from '@/components/SvgTracked.vue'
import SourceSelector from '@/components/SourceSelector.vue'
import { defaultField } from '@/composables/vision.ts'
import type { SSL_GeometryFieldSize } from '@/proto/vision/ssl_vision_geometry_pb.ts'

const route = useRoute()
const router = useRouter()

const logUrlInput = ref('')
const urlError = computed(() => {
  const url = logUrlInput.value.trim()
  if (url && !url.endsWith('.log')) return 'URL must end with .log'
  return ''
})

watch(
  () => route.query.url,
  (urlParam) => {
    if (typeof urlParam === 'string' && urlParam && !logUrlInput.value) {
      logUrlInput.value = urlParam
    }
  },
  { immediate: true },
)

watch(logUrlInput, (url) => {
  router.replace({ query: { ...route.query, url: url || undefined } })
})

const validLogUrl = computed(() => {
  const url = logUrlInput.value.trim()
  if (!url || !url.endsWith('.log')) return ''
  return url
})

const logBase = computed(() => {
  const url = validLogUrl.value
  if (!url) return ''
  return url.slice(0, -'.log'.length)
})

const logUrl = computed(() => validLogUrl.value)

const { index: visionIndex } = useLogFileIndex(computed(() => logBase.value ? `${logBase.value}.log.vision.idx` : ''))
const { index: refboxIndex } = useLogFileIndex(computed(() => logBase.value ? `${logBase.value}.log.refbox.idx` : ''))
const { index: trackerIndex } = useLogFileIndex(computed(() => logBase.value ? `${logBase.value}.log.tracker.idx` : ''))

const currentIndexPosition = ref(0)
let initialPositionSet = false

watch(logBase, () => {
  currentIndexPosition.value = 0
  initialPositionSet = false
})

const queryTimestampNs = computed<bigint | null>(() => {
  const ts = typeof route.query.timestamp === 'string' ? route.query.timestamp : ''
  if (!ts) return null
  const ms = new Date(ts).getTime()
  if (isNaN(ms)) return null
  return BigInt(ms) * BigInt(1_000_000)
})
watch([visionIndex, queryTimestampNs], ([index, targetNs]) => {
  if (initialPositionSet || !index || index.length === 0 || targetNs === null) return
  initialPositionSet = true
  let closest = 0
  let minDiff =
    targetNs > index[0]!.timestamp ? targetNs - index[0]!.timestamp : index[0]!.timestamp - targetNs
  for (let i = 1; i < index.length; i++) {
    const diff =
      targetNs > index[i]!.timestamp ? targetNs - index[i]!.timestamp : index[i]!.timestamp - targetNs
    if (diff < minDiff) {
      minDiff = diff
      closest = i
    }
  }
  currentIndexPosition.value = closest
})

const currentTimestamp = computed(() => {
  if (!visionIndex.value || visionIndex.value.length === 0) return BigInt(0)
  return visionIndex.value[currentIndexPosition.value]?.timestamp ?? BigInt(0)
})

const minTimestamp = computed(() => {
  if (!visionIndex.value || visionIndex.value.length === 0) return BigInt(0)
  return visionIndex.value[0]?.timestamp ?? BigInt(0)
})

const maxTimestamp = computed(() => {
  if (!visionIndex.value || visionIndex.value.length === 0) return BigInt(0)
  return visionIndex.value[visionIndex.value.length - 1]?.timestamp ?? BigInt(0)
})

const maxIndexPosition = computed(() => {
  if (!visionIndex.value) return 0
  return Math.max(0, visionIndex.value.length - 1)
})

const { message: visionMessage } = useLogFileMessageAtTimestamp(
  logUrl,
  visionIndex,
  currentTimestamp,
)
const { message: refboxMessage } = useLogFileMessageAtTimestamp(
  logUrl,
  refboxIndex,
  currentTimestamp,
)
const { message: trackerMessage } = useLogFileMessageAtTimestamp(
  logUrl,
  trackerIndex,
  currentTimestamp,
)

const latestField = ref<SSL_GeometryFieldSize>(defaultField)

const visionWrapper = computed(() => {
  if (!visionMessage.value) return undefined
  try {
    return fromBinary(SSL_WrapperPacketSchema, new Uint8Array(visionMessage.value.data))
  } catch {
    return undefined
  }
})

watch(visionWrapper, (wrapper) => {
  if (wrapper?.geometry?.field) {
    latestField.value = wrapper.geometry.field
  }
})

const detectionFrame = computed(() => visionWrapper.value?.detection)

const field = computed(() => latestField.value)

const referee = computed(() => {
  if (!refboxMessage.value) return undefined
  try {
    return fromBinary(RefereeSchema, new Uint8Array(refboxMessage.value.data))
  } catch {
    return undefined
  }
})

const trackerWrapper = computed(() => {
  if (!trackerMessage.value) return undefined
  try {
    return fromBinary(TrackerWrapperPacketSchema, new Uint8Array(trackerMessage.value.data))
  } catch {
    return undefined
  }
})

const trackerSources = ref<{ [uuid: string]: string }>({})

watch(trackerWrapper, (wrapper) => {
  if (wrapper?.uuid) {
    trackerSources.value = { ...trackerSources.value, [wrapper.uuid]: wrapper.sourceName }
  }
})

const activeSource = ref('vision')

const sources = computed(() => {
  return { vision: 'vision', ...trackerSources.value }
})

const trackedFrame = computed(() => {
  if (!trackerWrapper.value) return undefined
  if (activeSource.value !== trackerWrapper.value.uuid) return undefined
  return trackerWrapper.value.trackedFrame
})

const handleSliderChange = (event: Event) => {
  const target = event.target as HTMLInputElement
  currentIndexPosition.value = parseInt(target.value)
}

const formatTimestamp = (ts: bigint): string => {
  return new Date(Number(ts) / 1000 / 1000).toISOString()
}
</script>

<template>
  <div id="container">
    <div id="control">
      <div class="url-input-container">
        <input
          v-model="logUrlInput"
          type="text"
          placeholder="Enter log file URL (must end with .log)"
          class="url-input"
          :class="{ 'url-input-error': urlError }"
        />
        <span v-if="urlError" class="url-error">{{ urlError }}</span>
      </div>
      <div v-if="validLogUrl && visionIndex && visionIndex.length > 0" class="slider-container">
        <label>
          Timestamp: {{ formatTimestamp(currentTimestamp) }} ({{ currentIndexPosition + 1 }} /
          {{ visionIndex.length }})
        </label>
        <input
          type="range"
          min="0"
          :max="maxIndexPosition"
          :value="currentIndexPosition"
          @input="handleSliderChange"
        />
        <div class="timestamp-range">
          <span>{{ formatTimestamp(minTimestamp) }}</span>
          <span>{{ formatTimestamp(maxTimestamp) }}</span>
        </div>
      </div>
      <div v-else-if="validLogUrl">Loading index...</div>
    </div>
    <div id="content">
      <FieldVisualizer :field="field">
        <SvgVision
          v-if="activeSource === 'vision' && detectionFrame"
          :detection-frame="detectionFrame"
        />
        <SvgReferee v-if="referee" :field="field" :referee="referee" />
        <SvgTracked v-if="trackedFrame" :tracked-frame="trackedFrame" />
      </FieldVisualizer>
      <SourceSelector :sources="sources" v-model="activeSource" />
    </div>
  </div>
</template>

<style scoped>
#container {
  display: flex;
  flex-direction: column;
  height: 100%;
  padding: 20px;
}

#control {
  width: 100%;
  margin-bottom: 20px;
}

.url-input-container {
  display: flex;
  flex-direction: column;
  gap: 4px;
  margin-bottom: 12px;
}

.url-input {
  width: 100%;
  padding: 6px 10px;
  font-size: 14px;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
}

.url-input-error {
  border-color: #e53e3e;
}

.url-error {
  font-size: 12px;
  color: #e53e3e;
}

.slider-container {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.slider-container label {
  font-weight: bold;
  font-size: 16px;
}

input[type='range'] {
  width: 100%;
}

.timestamp-range {
  display: flex;
  justify-content: space-between;
  font-size: 12px;
  color: #666;
}

#content {
  width: 100%;
  flex: 1;
  overflow: hidden;
}
</style>
