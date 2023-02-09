function determineWebSocketAddress() {
    const protocol = window.location.protocol === 'http:' ? 'ws:' : 'wss:'
    const urlParams = new URLSearchParams(window.location.search)
    const wsAddress = protocol + '//' + window.location.hostname + ':' + window.location.port + '/api/vision'

    const sourceId = urlParams.get('sourceId')
    if (sourceId) {
        return `${wsAddress}?sourceId=${sourceId}`
    }
    return wsAddress
}

export function connect(callback: Function) {
    const ws = new WebSocket(determineWebSocketAddress());

    ws.onmessage = (e) => {
        callback(JSON.parse(e.data))
    };

    ws.onclose = () => {
        setTimeout(() => {
            connect(callback)
        }, 1000);
    };

    ws.onerror = () => {
        ws.close()
    };
}
