import { computed, toValue } from 'vue'
import type { MaybeRefOrGetter } from 'vue'
import { useLogFileMultiIndex } from '@/composables/logfile.ts'
import type { IndexEntry } from '@/composables/logfile.ts'
import type { useLogManifest } from '@/composables/logManifest.ts'
import { getLogsBaseUrl } from '@/composables/logUrl.ts'
import { SOURCE_VISION } from '@/composables/vision.ts'

const indexBounds = (index: IndexEntry[] | undefined): { min: bigint; max: bigint } | undefined => {
  if (!index || index.length === 0) return undefined
  return { min: index[0]!.timestamp, max: index[index.length - 1]!.timestamp }
}

export const useLogPlayerIndices = (
  logUrl: MaybeRefOrGetter<string>,
  manifest: ReturnType<typeof useLogManifest>,
  activeSource: MaybeRefOrGetter<string>,
) => {
  const logsBaseUrl = computed(() => getLogsBaseUrl(toValue(logUrl)))

  const { indices: visionDetectionIndices, primaryIndex: visionDetectionPrimaryIndex } =
    useLogFileMultiIndex(logsBaseUrl, manifest.manifestEntriesOfType('vision_detection'))

  // Load ALL tracker indices unconditionally for global timeline bounds
  const { primaryIndex: trackerPrimaryIndex } = useLogFileMultiIndex(
    logsBaseUrl,
    manifest.manifestEntriesOfType('tracker'),
  )

  // Primary index switches based on active source (used for step mode)
  const primaryIndex = computed(() => {
    if (toValue(activeSource) === SOURCE_VISION) {
      return visionDetectionPrimaryIndex.value
    }
    return trackerPrimaryIndex.value
  })

  // Global timestamp bounds across ALL loaded indices
  const globalMinTimestamp = computed(() => {
    const bounds = [
      indexBounds(visionDetectionPrimaryIndex.value),
      indexBounds(trackerPrimaryIndex.value),
    ].filter((b) => b !== undefined)
    if (bounds.length === 0) return BigInt(0)
    return bounds.reduce((a, b) => (a.min < b.min ? a : b)).min
  })

  const globalMaxTimestamp = computed(() => {
    const bounds = [
      indexBounds(visionDetectionPrimaryIndex.value),
      indexBounds(trackerPrimaryIndex.value),
    ].filter((b) => b !== undefined)
    if (bounds.length === 0) return BigInt(0)
    return bounds.reduce((a, b) => (a.max > b.max ? a : b)).max
  })

  return {
    visionDetectionIndices,
    primaryIndex,
    globalMinTimestamp,
    globalMaxTimestamp,
  }
}
