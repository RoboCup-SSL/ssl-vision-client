import { shallowRef } from 'vue'
import { fromBinary } from '@bufbuild/protobuf'
import { ReconnectingWebSocket } from '@/helpers/websocket.ts'
import {
  type TrackedFrame,
  TrackedFrameSchema,
} from '@/proto/tracked/ssl_vision_detection_tracked_pb.ts'
import { Team } from '@/proto/gc/ssl_gc_common_pb.ts'

const defaultTrackedFrame: TrackedFrame = {
  $typeName: 'TrackedFrame',
  frameNumber: 0,
  timestamp: 0,
  balls: [],
  robots: [],
  capabilities: [],
}

if (import.meta.env.MODE === 'development') {
  defaultTrackedFrame.balls = [
    {
      $typeName: 'TrackedBall',
      pos: {
        $typeName: 'Vector3',
        x: 1,
        y: 2,
        z: 0,
      },
      vel: {
        $typeName: 'Vector3',
        x: 0,
        y: 0,
        z: 0,
      },
      visibility: 1,
    },
  ]

  defaultTrackedFrame.robots = [
    {
      $typeName: 'TrackedRobot',
      robotId: {
        $typeName: 'RobotId',
        id: 0,
        team: Team.BLUE,
      },
      pos: {
        $typeName: 'Vector2',
        x: -1,
        y: -2,
      },
      vel: {
        $typeName: 'Vector2',
        x: 0,
        y: 0,
      },
      orientation: 1.57,
      velAngular: 0,
      visibility: 1,
    },
    {
      $typeName: 'TrackedRobot',
      robotId: {
        $typeName: 'RobotId',
        id: 15,
        team: Team.YELLOW,
      },
      pos: {
        $typeName: 'Vector2',
        x: 5.9,
        y: 0,
      },
      vel: {
        $typeName: 'Vector2',
        x: 0,
        y: 0,
      },
      orientation: 0,
      velAngular: 0,
      visibility: 1,
    },
  ]
}

export const useTrackedFrame = () => {
  const ws = new ReconnectingWebSocket('/api/tracker')

  const trackedFrame = shallowRef<TrackedFrame>(defaultTrackedFrame)
  ws.registerBytesConsumer((data: Uint8Array) => {
    trackedFrame.value = fromBinary(TrackedFrameSchema, data)
  })

  const updateSourceId = (sourceId: string) => {
    ws.SendJSON({ tracker_source: sourceId })
  }

  return { trackedFrame, updateSourceId }
}

export const useTrackedSources = () => {
  const ws = new ReconnectingWebSocket('/api/tracker/sources')

  const trackerSources = shallowRef<{ [id: string]: string }>({})
  ws.registerTextConsumer((data: string) => {
    const controlMessage = JSON.parse(data)
    trackerSources.value = controlMessage.tracker_sources
  })

  return { trackerSources }
}
