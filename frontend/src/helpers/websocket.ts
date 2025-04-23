export const determineWebSocketAddress = (path: string) => {
  const protocol = window.location.protocol === 'http:' ? 'ws:' : 'wss:'
  return protocol + '//' + window.location.hostname + ':' + window.location.port + path
}

export class ReconnectingWebSocket {
  private readonly apiPath: string
  private ws?: WebSocket
  private readonly bytesConsumer: ((data: Uint8Array) => void)[] = []
  private readonly textConsumer: ((data: string) => void)[] = []

  constructor(path: string) {
    this.apiPath = path
  }

  public registerBytesConsumer(cb: (data: Uint8Array) => void) {
    this.bytesConsumer.push(cb)
  }

  public registerTextConsumer(cb: (data: string) => void) {
    this.textConsumer.push(cb)
  }

  public SendJSON(data: object) {
    if (this.ws) {
      this.ws.send(JSON.stringify(data))
    }
  }

  public connect() {
    const ws = new WebSocket(determineWebSocketAddress(this.apiPath))

    ws.onmessage = async (e) => {
      if (this.bytesConsumer.length > 0) {
        const blob = await e.data.arrayBuffer()
        const data = new Uint8Array(blob)
        for (const callback of this.bytesConsumer) {
          callback(data)
        }
      }
      if (this.textConsumer.length > 0) {
        const data = e.data as string
        for (const callback of this.textConsumer) {
          callback(data)
        }
      }
    }

    ws.onclose = () => {
      this.ws = undefined
      setTimeout(() => {
        this.connect()
      }, 1000)
    }

    ws.onerror = () => {
      ws.close()
    }

    this.ws = ws
  }

  public disconnect() {
    if (this.ws) {
      this.ws.close()
    }
  }
}
