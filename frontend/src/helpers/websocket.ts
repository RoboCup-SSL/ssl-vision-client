export const determineWebSocketAddress = (path: string) => {
  const protocol = window.location.protocol === 'http:' ? 'ws:' : 'wss:'
  return protocol + '//' + window.location.hostname + ':' + window.location.port + path
}
