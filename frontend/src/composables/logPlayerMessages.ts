import { computed, ref, toValue, watch } from 'vue'
import type { MaybeRefOrGetter } from 'vue'
import {
  useLogFileIndex,
  useLogFileMessageAtTimestamp,
  useLogFileMessagesAtTimestamp,
  parseProtobuf,
} from '@/composables/logfile.ts'
import type { IndexEntry } from '@/composables/logfile.ts'
import type { useLogManifest } from '@/composables/logManifest.ts'
import { SSL_WrapperPacketSchema } from '@/proto/vision/ssl_vision_wrapper_pb.ts'
import { RefereeSchema } from '@/proto/gc/ssl_gc_referee_message_pb.ts'
import { TrackerWrapperPacketSchema } from '@/proto/tracked/ssl_vision_wrapper_tracked_pb.ts'
import { applyFieldDefaults, defaultField, SOURCE_VISION } from '@/composables/vision.ts'
import type { SSL_GeometryFieldSize } from '@/proto/vision/ssl_vision_geometry_pb.ts'

export const useLogPlayerMessages = (
  logUrl: MaybeRefOrGetter<string>,
  manifest: ReturnType<typeof useLogManifest>,
  visionDetectionIndices: MaybeRefOrGetter<Map<string, IndexEntry[]>>,
  currentTimestamp: MaybeRefOrGetter<bigint>,
  activeSource: MaybeRefOrGetter<string>,
) => {
  const isVisionActive = computed(() => toValue(activeSource) === SOURCE_VISION)

  const { index: visionGeometryIndex } = useLogFileIndex(
    computed(() => manifest.manifestIndexUrl('vision_geometry')),
  )
  const { index: refboxIndex } = useLogFileIndex(
    computed(() => manifest.manifestIndexUrl('refbox')),
  )
  const { index: trackerIndex } = useLogFileIndex(
    computed(() => {
      const source = toValue(activeSource)
      if (source === SOURCE_VISION) return ''
      return manifest.manifestIndexUrl('tracker', source)
    }),
  )

  // Only fetch vision detection messages when vision is active
  const { messages: visionDetectionMessages } = useLogFileMessagesAtTimestamp(
    logUrl,
    computed(() => (isVisionActive.value ? toValue(visionDetectionIndices) : new Map())),
    currentTimestamp,
  )

  const { message: visionGeometryMessage } = useLogFileMessageAtTimestamp(
    logUrl,
    visionGeometryIndex,
    currentTimestamp,
  )
  const { message: refboxMessage } = useLogFileMessageAtTimestamp(
    logUrl,
    refboxIndex,
    currentTimestamp,
  )

  // Only fetch tracker message when a tracker source is active
  const { message: trackerMessage } = useLogFileMessageAtTimestamp(
    logUrl,
    computed(() => (isVisionActive.value ? undefined : trackerIndex.value)),
    currentTimestamp,
  )

  const visionDetectionWrappers = computed(() =>
    visionDetectionMessages.value
      .map((msg) => parseProtobuf(msg, SSL_WrapperPacketSchema, 'vision detection wrapper'))
      .filter((w) => w !== undefined),
  )

  const visionGeometryWrapper = computed(() =>
    parseProtobuf(visionGeometryMessage.value, SSL_WrapperPacketSchema, 'vision geometry wrapper'),
  )

  const latestField = ref<SSL_GeometryFieldSize>(defaultField)

  watch(visionGeometryWrapper, (wrapper) => {
    if (wrapper?.geometry?.field) {
      latestField.value = applyFieldDefaults(wrapper.geometry.field)
    }
  })

  const detectionFrames = computed(() => {
    if (!isVisionActive.value) return []
    return visionDetectionWrappers.value.map((w) => w.detection).filter((d) => d !== undefined)
  })

  const field = computed(() => latestField.value)

  const referee = computed(() =>
    parseProtobuf(refboxMessage.value, RefereeSchema, 'referee message'),
  )

  const trackedFrame = computed(() => {
    if (isVisionActive.value) return undefined
    const wrapper = parseProtobuf(
      trackerMessage.value,
      TrackerWrapperPacketSchema,
      'tracker wrapper',
    )
    return wrapper?.trackedFrame
  })

  return {
    detectionFrames,
    field,
    referee,
    trackedFrame,
  }
}
