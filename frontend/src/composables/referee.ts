import { type Referee, RefereeSchema } from '@/proto/gc/ssl_gc_referee_message_pb.ts'
import { useWebSocketProtobuf } from '@/composables/websocket.ts'

export const useReferee = () => {
  const { message: referee } = useWebSocketProtobuf<Referee>('/api/referee', RefereeSchema)
  return { referee }
}
