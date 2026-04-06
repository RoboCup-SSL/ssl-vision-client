import { computed, onUnmounted, ref, toValue, watch } from 'vue'
import type { MaybeRefOrGetter } from 'vue'
import { useRoute } from 'vue-router'
import { findClosestIndex } from '@/composables/logfile.ts'
import type { IndexEntry } from '@/composables/logfile.ts'

const STORAGE_TIMESTAMP_KEY = 'ssl-vision-timestamp'

const safeParseBigInt = (value: string): bigint | null => {
  try {
    return BigInt(value)
  } catch {
    return null
  }
}

export const useLogPlayerPosition = (
  globalMinTimestamp: MaybeRefOrGetter<bigint>,
  globalMaxTimestamp: MaybeRefOrGetter<bigint>,
  primaryIndex: MaybeRefOrGetter<IndexEntry[] | undefined>,
  logUrl: MaybeRefOrGetter<string>,
) => {
  const route = useRoute()

  const currentTimestamp = ref(BigInt(0))
  const initialPositionSet = ref(false)

  // Reset when log file changes
  watch(
    () => toValue(logUrl),
    () => {
      currentTimestamp.value = BigInt(0)
      initialPositionSet.value = false
    },
  )

  // Restore initial timestamp from query param or localStorage
  const initialTimestampNs = computed<bigint | null>(() => {
    const ts = typeof route.query.ts === 'string' ? route.query.ts : ''
    if (ts) return safeParseBigInt(ts)
    const saved = localStorage.getItem(STORAGE_TIMESTAMP_KEY)
    if (saved) return safeParseBigInt(saved)
    return null
  })

  watch([() => toValue(globalMaxTimestamp), initialTimestampNs], ([maxTs, targetNs]) => {
    if (initialPositionSet.value || maxTs === BigInt(0)) return
    initialPositionSet.value = true
    if (targetNs !== null) {
      currentTimestamp.value = targetNs
    }
  })

  // Slider mapping: linear over global timestamp range
  const sliderMax = computed(() => {
    const min = toValue(globalMinTimestamp)
    const max = toValue(globalMaxTimestamp)
    if (max <= min) return 0
    return Number(max - min)
  })

  const sliderValue = computed({
    get: () => {
      const min = toValue(globalMinTimestamp)
      if (sliderMax.value === 0) return 0
      const val = currentTimestamp.value - min
      return Number(
        val < BigInt(0) ? BigInt(0) : val > BigInt(sliderMax.value) ? BigInt(sliderMax.value) : val,
      )
    },
    set: (val: number) => {
      currentTimestamp.value = toValue(globalMinTimestamp) + BigInt(val)
    },
  })

  // Step through frames of the active source's index
  const stepForward = () => {
    const index = toValue(primaryIndex)
    if (!index || index.length === 0) return
    const currentIdx = findClosestIndex(index, currentTimestamp.value)
    const nextIdx = Math.min(currentIdx + 1, index.length - 1)
    currentTimestamp.value = index[nextIdx]!.timestamp
  }

  const stepBackward = () => {
    const index = toValue(primaryIndex)
    if (!index || index.length === 0) return
    const currentIdx = findClosestIndex(index, currentTimestamp.value)
    const prevIdx = Math.max(currentIdx - 1, 0)
    currentTimestamp.value = index[prevIdx]!.timestamp
  }

  const skipTime = (seconds: number) => {
    const deltaNs = BigInt(seconds) * BigInt(1_000_000_000)
    const min = toValue(globalMinTimestamp)
    const max = toValue(globalMaxTimestamp)
    let newTs = currentTimestamp.value + deltaNs
    if (newTs < min) newTs = min
    if (newTs > max) newTs = max
    currentTimestamp.value = newTs
  }

  // Play mode: advance timestamp in real time
  const playing = ref(false)
  let playTimer: ReturnType<typeof setInterval> | null = null
  const PLAY_INTERVAL_MS = 25

  const startPlaying = () => {
    if (playTimer) return
    playTimer = setInterval(() => {
      const deltaNs = BigInt(PLAY_INTERVAL_MS) * BigInt(1_000_000)
      const max = toValue(globalMaxTimestamp)
      const newTs = currentTimestamp.value + deltaNs
      if (newTs >= max) {
        currentTimestamp.value = max
        stopPlaying()
      } else {
        currentTimestamp.value = newTs
      }
    }, PLAY_INTERVAL_MS)
  }

  const stopPlaying = () => {
    if (playTimer) {
      clearInterval(playTimer)
      playTimer = null
    }
    playing.value = false
  }

  const togglePlay = () => {
    playing.value = !playing.value
    if (playing.value) {
      startPlaying()
    } else {
      stopPlaying()
    }
  }

  onUnmounted(() => {
    if (playTimer) clearInterval(playTimer)
  })

  // Persist current timestamp to localStorage for session restore
  watch(currentTimestamp, (timestamp) => {
    if (timestamp === BigInt(0)) {
      localStorage.removeItem(STORAGE_TIMESTAMP_KEY)
    } else {
      localStorage.setItem(STORAGE_TIMESTAMP_KEY, timestamp.toString())
    }
  })

  return {
    currentTimestamp,
    sliderValue,
    sliderMax,
    stepForward,
    stepBackward,
    skipTime,
    playing,
    togglePlay,
  }
}
