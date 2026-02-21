<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { fromBinary } from '@bufbuild/protobuf'
import { RefereeSchema } from '@/proto/gc/ssl_gc_referee_message_pb.ts'
import { SSL_WrapperPacketSchema } from '@/proto/vision/ssl_vision_wrapper_pb.ts'
import { TrackerWrapperPacketSchema } from '@/proto/tracked/ssl_vision_wrapper_tracked_pb.ts'
import { useLogFileIndex, useLogFileMessageAtTimestamp } from '@/composables/logfile.ts'
import FieldVisualizer from '@/components/FieldVisualizer.vue'
import SvgVision from '@/components/SvgVision.vue'
import SvgReferee from '@/components/SvgReferee.vue'
import SvgTracked from '@/components/SvgTracked.vue'
import { defaultField } from '@/composables/vision.ts'
import type { SSL_GeometryFieldSize } from '@/proto/vision/ssl_vision_geometry_pb.ts'

const logBase = '/logs/2025-03-13_14-00_GROUP_PHASE_ER-Force-vs-Immortals'
const logUrl = `${logBase}.log`

const { index: visionIndex } = useLogFileIndex(`${logBase}.log.vision.idx`)
const { index: refboxIndex } = useLogFileIndex(`${logBase}.log.refbox.idx`)
const { index: trackerIndex } = useLogFileIndex(`${logBase}.log.tracker.idx`)

const currentIndexPosition = ref(0)

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

const trackedFrame = computed(() => {
  if (!trackerMessage.value) return undefined
  try {
    const wrapper = fromBinary(
      TrackerWrapperPacketSchema,
      new Uint8Array(trackerMessage.value.data),
    )
    return wrapper.trackedFrame
  } catch {
    return undefined
  }
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
      <div v-if="visionIndex && visionIndex.length > 0" class="slider-container">
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
      <div v-else>Loading index...</div>
    </div>
    <div id="content">
      <FieldVisualizer :field="field">
        <SvgVision v-if="detectionFrame" :detection-frame="detectionFrame" />
        <SvgReferee v-if="referee" :field="field" :referee="referee" />
        <SvgTracked v-if="trackedFrame" :tracked-frame="trackedFrame" />
      </FieldVisualizer>
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
