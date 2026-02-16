<script setup lang="ts">
import { computed, ref } from 'vue'
import { fromBinary, toJsonString } from '@bufbuild/protobuf'
import { RefereeSchema } from '@/proto/gc/ssl_gc_referee_message_pb.ts'
import { useLogFileIndex, useLogFileMessageAtTimestamp } from '@/composables/logfile.ts'

const logUrl = '/logs/2025-03-13_14-00_GROUP_PHASE_ER-Force-vs-Immortals.log'
const indexUrl = '/logs/2025-03-13_14-00_GROUP_PHASE_ER-Force-vs-Immortals.log.refbox.idx'

const { index } = useLogFileIndex(indexUrl)

const currentIndexPosition = ref(0)

const currentTimestamp = computed(() => {
  if (!index.value || index.value.length === 0) return BigInt(0)
  return index.value[currentIndexPosition.value]?.timestamp ?? BigInt(0)
})

const minTimestamp = computed(() => {
  if (!index.value || index.value.length === 0) return BigInt(0)
  return index.value[0]?.timestamp ?? BigInt(0)
})

const maxTimestamp = computed(() => {
  if (!index.value || index.value.length === 0) return BigInt(0)
  return index.value[index.value.length - 1]?.timestamp ?? BigInt(0)
})

const { message } = useLogFileMessageAtTimestamp(logUrl, index, currentTimestamp)

const decodedJson = computed(() => {
  if (!message.value) return ''

  try {
    const bytes = new Uint8Array(message.value.data)
    const decoded = fromBinary(RefereeSchema, bytes)
    return toJsonString(RefereeSchema, decoded, { prettySpaces: 2 })
  } catch (error) {
    return `Error decoding message: ${error}`
  }
})

const handleSliderChange = (event: Event) => {
  const target = event.target as HTMLInputElement
  currentIndexPosition.value = parseInt(target.value)
}

const formatTimestamp = (ts: bigint): string => {
  return new Date(Number(ts) / 1000 / 1000).toISOString()
}

const maxIndexPosition = computed(() => {
  if (!index.value) return 0
  return Math.max(0, index.value.length - 1)
})
</script>

<template>
  <div id="container">
    <div id="control">
      <div v-if="index" class="slider-container">
        <label>
          Timestamp: {{ formatTimestamp(currentTimestamp) }}
          ({{ currentIndexPosition + 1 }} / {{ index.length }})
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
      <pre v-if="decodedJson">{{ decodedJson }}</pre>
      <div v-else>No message loaded</div>
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
  height: 100%;
  overflow: auto;
}

pre {
  background-color: #f5f5f5;
  padding: 15px;
  border-radius: 5px;
  overflow: auto;
  font-family: 'Courier New', monospace;
  font-size: 14px;
  line-height: 1.5;
}
</style>
