import type {Field} from "@/field";

export interface Request {
    activeSourceId: string,
}

export class VisionApi {
    private readonly apiPath = '/api/vision'
    private ws ?: WebSocket
    private readonly consumer: ((message: Field) => any)[] = []
    private latestField ?: Field

    constructor() {
        this.connect()
    }

    public Send(request: Request) {
        const ws = this.ws
        if (ws) {
            const json = JSON.stringify(request)
            ws.send(json)
        } else {
            console.warn("No WebSocket connection. Dropping ", request)
        }
    }

    public RegisterConsumer(cb: ((output: Field) => void)) {
        this.consumer.push(cb)
        if (this.latestField) {
            cb(this.latestField)
        }
    }

    private determineWebSocketAddress() {
        const protocol = window.location.protocol === 'http:' ? 'ws:' : 'wss:'
        const urlParams = new URLSearchParams(window.location.search)
        const wsAddress = protocol + '//' + window.location.hostname + ':' + window.location.port + this.apiPath

        const sourceId = urlParams.get('sourceId')
        if (sourceId) {
            return `${wsAddress}?sourceId=${sourceId}`
        }
        return wsAddress
    }

    private connect() {
        const ws = new WebSocket(this.determineWebSocketAddress());

        ws.onmessage = (e) => {
            this.latestField = JSON.parse(e.data)
            for (const callback of this.consumer) {
                callback(this.latestField!)
            }
        };

        ws.onclose = () => {
            this.ws = undefined
            setTimeout(() => {
                this.connect()
            }, 1000);
        };

        ws.onerror = () => {
            ws.close()
        };

        this.ws = ws;
    }
}
