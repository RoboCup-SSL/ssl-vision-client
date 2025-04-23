import {
  type SSL_DetectionFrame,
  SSL_DetectionFrameSchema,
} from '@/proto/vision/ssl_vision_detection_pb.ts'
import { computed, type MaybeRefOrGetter, toValue } from 'vue'
import {
  type SSL_GeometryData,
  SSL_GeometryDataSchema,
  type SSL_GeometryFieldSize,
} from '@/proto/vision/ssl_vision_geometry_pb.ts'
import { useWebSocketProtobuf } from '@/composables/websocket.ts'
import {
  type TrackedFrame,
  TrackedFrameSchema,
} from '@/proto/tracked/ssl_vision_detection_tracked_pb.ts'
import { useWebSocket } from '@vueuse/core'
import { determineWebSocketAddress } from '@/helpers/websocket.ts'

const defaultField: SSL_GeometryFieldSize = {
  $typeName: 'SSL_GeometryFieldSize',
  fieldLength: 12000,
  fieldWidth: 9000,
  goalWidth: 1800,
  goalDepth: 180,
  boundaryWidth: 300,
  penaltyAreaDepth: 1800,
  penaltyAreaWidth: 3600,
  centerCircleRadius: 500,
  lineThickness: 10,
  goalCenterToPenaltyMark: 8000,
  goalHeight: 155,
  ballRadius: 21.5,
  maxRobotRadius: 90,
  fieldLines: [],
  fieldArcs: [],
}

export const SOURCE_VISION = 'vision'

export const useVisionDetection = (activeSource: MaybeRefOrGetter<string>) => {
  const enabled = computed(() => toValue(activeSource) === SOURCE_VISION)

  const { message: detectionFrame } = useWebSocketProtobuf<SSL_DetectionFrame>(
    '/api/vision/detection',
    SSL_DetectionFrameSchema,
    enabled,
  )

  return { detectionFrame }
}

export const useVisionGeometry = () => {
  const { message } = useWebSocketProtobuf<SSL_GeometryData>(
    '/api/vision/geometry',
    SSL_GeometryDataSchema,
  )

  const field = computed(() => {
    if (message.value && message.value.field) {
      return message.value.field
    }
    return defaultField
  })

  return { field }
}

export const useTrackedFrame = (activeSource: MaybeRefOrGetter<string>) => {
  const enabled = computed(() => toValue(activeSource) !== SOURCE_VISION)
  const path = computed(() => {
    return '/api/tracker?source=' + toValue(activeSource)
  })

  const { message: trackedFrame } = useWebSocketProtobuf<TrackedFrame>(
    path,
    TrackedFrameSchema,
    enabled,
  )

  return { trackedFrame }
}

export const useTrackedSources = () => {
  const { data } = useWebSocket<string>(determineWebSocketAddress('/api/tracker/sources'), {
    autoReconnect: true,
  })

  const trackerSources = computed(() => {
    if (data.value) {
      const controlMessage = JSON.parse(data.value)
      return controlMessage.tracker_sources
    }
    return {}
  })

  return { trackerSources }
}
