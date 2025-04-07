import { ReconnectingWebSocket } from '@/helpers/websocket.ts'
import {
  type SSL_DetectionFrame,
  SSL_DetectionFrameSchema,
} from '@/proto/vision/ssl_vision_detection_pb.ts'
import { fromBinary } from '@bufbuild/protobuf'
import { shallowRef } from 'vue'
import {
  SSL_GeometryDataSchema,
  type SSL_GeometryFieldSize,
} from '@/proto/vision/ssl_vision_geometry_pb.ts'

const defaultDetectionFrame: SSL_DetectionFrame = {
  $typeName: 'SSL_DetectionFrame',
  frameNumber: 0,
  cameraId: 0,
  tCapture: 0,
  tSent: 0,
  balls: [],
  robotsBlue: [],
  robotsYellow: [],
}

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

if (import.meta.env.MODE === 'development') {
  defaultDetectionFrame.balls = [
    {
      $typeName: 'SSL_DetectionBall',
      confidence: 0,
      area: 0,
      x: 1000,
      y: 2000,
      z: 0,
      pixelX: 0,
      pixelY: 0,
    },
  ]

  defaultDetectionFrame.robotsBlue = [
    {
      $typeName: 'SSL_DetectionRobot',
      robotId: 0,
      confidence: 0,
      x: -1000,
      y: -2000,
      orientation: 1.57,
      height: 150,
      pixelX: 0,
      pixelY: 0,
    },
  ]
  defaultDetectionFrame.robotsYellow = [
    {
      $typeName: 'SSL_DetectionRobot',
      robotId: 15,
      confidence: 0,
      x: 5900,
      y: 0,
      orientation: 0,
      height: 150,
      pixelX: 0,
      pixelY: 0,
    },
  ]
}

export const useVisionDetection = () => {
  const ws = new ReconnectingWebSocket('/api/vision/detection')

  const detectionFrame = shallowRef<SSL_DetectionFrame>(defaultDetectionFrame)
  ws.registerBytesConsumer((data: Uint8Array) => {
    detectionFrame.value = fromBinary(SSL_DetectionFrameSchema, data)
  })

  return { detectionFrame }
}

export const useVisionGeometry = () => {
  const ws = new ReconnectingWebSocket('/api/vision/geometry')

  const field = shallowRef<SSL_GeometryFieldSize>(defaultField)
  ws.registerBytesConsumer((data: Uint8Array) => {
    const geometry = fromBinary(SSL_GeometryDataSchema, data)
    if (geometry.field) {
      field.value = geometry.field
    }
  })

  return { field }
}
