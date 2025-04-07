import { shallowRef } from 'vue'
import {
  type Referee,
  type RefereeJson,
  RefereeSchema,
} from '@/proto/gc/ssl_gc_referee_message_pb.ts'
import { fromBinary, fromJson } from '@bufbuild/protobuf'
import { ReconnectingWebSocket } from '@/helpers/websocket.ts'

const defaultReferee: RefereeJson = {
  matchType: 'UNKNOWN_MATCH',
  packetTimestamp: '0',
  stage: 'NORMAL_FIRST_HALF_PRE',
  stageTimeLeft: 0,
  command: 'HALT',
  commandCounter: 0,
  commandTimestamp: '0',
  yellow: {
    name: 'Yellow',
    score: 0,
    redCards: 0,
    yellowCards: 0,
    timeouts: 0,
    timeoutTime: 0,
    goalkeeper: 0,
  },
  blue: {
    name: 'Blue',
    score: 0,
    redCards: 0,
    yellowCards: 0,
    timeouts: 0,
    timeoutTime: 0,
    goalkeeper: 0,
  },
  blueTeamOnPositiveHalf: false,
  designatedPosition: {
    x: 6000,
    y: 4500,
  }
}

export const useReferee = () => {
  const ws = new ReconnectingWebSocket('/api/referee')

  const defRef = fromJson(RefereeSchema, defaultReferee)
  const referee = shallowRef<Referee>(defRef)
  ws.registerBytesConsumer((data: Uint8Array) => {
    referee.value = fromBinary(RefereeSchema, data)
  })

  return { referee }
}
