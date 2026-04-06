import { describe, it, expect, vi, beforeEach } from 'vitest'
import { ref, nextTick } from 'vue'
import {
  findClosestEntry,
  findClosestIndex,
  useLogFileMultiIndex,
  useLogFileMessagesAtTimestamp,
} from '../logfile'
import type { IndexEntry, ManifestEntry } from '../logfile'

// Mock axios
vi.mock('axios', () => ({
  default: {
    get: vi.fn(),
  },
}))

import axios from 'axios'
const mockedAxios = vi.mocked(axios)

function createIndexBuffer(entries: { timestamp: bigint; offset: bigint }[]): ArrayBuffer {
  const buffer = new ArrayBuffer(entries.length * 16)
  const view = new DataView(buffer)
  for (let i = 0; i < entries.length; i++) {
    view.setBigInt64(i * 16, entries[i]!.timestamp, false)
    view.setBigInt64(i * 16 + 8, entries[i]!.offset, false)
  }
  return buffer
}

function createMessageBuffer(
  timestamp: bigint,
  messageType: number,
  payload: Uint8Array,
): ArrayBuffer {
  const buffer = new ArrayBuffer(16 + payload.length)
  const view = new DataView(buffer)
  view.setBigInt64(0, timestamp)
  view.setInt32(8, messageType)
  view.setInt32(12, payload.length)
  new Uint8Array(buffer).set(payload, 16)
  return buffer
}

describe('findClosestEntry', () => {
  const entries: IndexEntry[] = [
    { timestamp: 100n, offset: 0n },
    { timestamp: 200n, offset: 100n },
    { timestamp: 300n, offset: 200n },
    { timestamp: 400n, offset: 300n },
  ]

  it('returns undefined for empty index', () => {
    expect(findClosestEntry([], 100n)).toBeUndefined()
  })

  it('returns undefined when timestamp is before all entries', () => {
    expect(findClosestEntry(entries, 50n)).toBeUndefined()
  })

  it('returns exact match', () => {
    expect(findClosestEntry(entries, 200n)).toEqual({ timestamp: 200n, offset: 100n })
  })

  it('returns closest entry not in the future', () => {
    expect(findClosestEntry(entries, 250n)).toEqual({ timestamp: 200n, offset: 100n })
  })

  it('returns last entry when timestamp is beyond all entries', () => {
    expect(findClosestEntry(entries, 500n)).toEqual({ timestamp: 400n, offset: 300n })
  })

  it('returns first entry when timestamp equals first entry', () => {
    expect(findClosestEntry(entries, 100n)).toEqual({ timestamp: 100n, offset: 0n })
  })
})

describe('findClosestIndex', () => {
  const entries: IndexEntry[] = [
    { timestamp: 100n, offset: 0n },
    { timestamp: 200n, offset: 100n },
    { timestamp: 300n, offset: 200n },
    { timestamp: 400n, offset: 300n },
  ]

  it('returns 0 for empty index', () => {
    expect(findClosestIndex([], 100n)).toBe(0)
  })

  it('returns 0 when timestamp is before all entries', () => {
    expect(findClosestIndex(entries, 50n)).toBe(0)
  })

  it('returns exact match index', () => {
    expect(findClosestIndex(entries, 200n)).toBe(1)
  })

  it('returns floor index (last entry not in the future)', () => {
    expect(findClosestIndex(entries, 250n)).toBe(1)
  })

  it('returns last index when timestamp is beyond all entries', () => {
    expect(findClosestIndex(entries, 500n)).toBe(3)
  })

  it('returns 0 when timestamp equals first entry', () => {
    expect(findClosestIndex(entries, 100n)).toBe(0)
  })
})

describe('useLogFileMultiIndex', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('returns empty map and undefined primaryIndex for empty entries', async () => {
    const baseUrl = ref('')
    const entries = ref<ManifestEntry[]>([])
    const { indices, primaryIndex } = useLogFileMultiIndex(baseUrl, entries)

    await nextTick()
    await nextTick()

    expect(indices.value.size).toBe(0)
    expect(primaryIndex.value).toBeUndefined()
  })

  it('loads multiple camera indices', async () => {
    const cam0Index = createIndexBuffer([
      { timestamp: 100n, offset: 0n },
      { timestamp: 200n, offset: 50n },
    ])
    const cam1Index = createIndexBuffer([
      { timestamp: 110n, offset: 0n },
      { timestamp: 210n, offset: 60n },
    ])

    mockedAxios.get.mockImplementation((url: string) => {
      if (url.includes('cam0')) return Promise.resolve({ data: cam0Index })
      if (url.includes('cam1')) return Promise.resolve({ data: cam1Index })
      return Promise.reject(new Error('unknown url'))
    })

    const baseUrl = ref('http://example.com/logs/')
    const entries = ref<ManifestEntry[]>([
      { type: 'vision_detection', path: 'test.cam0.idx', source: 'cam0' },
      { type: 'vision_detection', path: 'test.cam1.idx', source: 'cam1' },
    ])

    const { indices, primaryIndex } = useLogFileMultiIndex(baseUrl, entries)

    // Wait for computedAsync to resolve
    await vi.waitFor(() => {
      expect(indices.value.size).toBe(2)
    })

    expect(indices.value.has('cam0')).toBe(true)
    expect(indices.value.has('cam1')).toBe(true)
    expect(indices.value.get('cam0')).toHaveLength(2)
    expect(indices.value.get('cam1')).toHaveLength(2)
    expect(primaryIndex.value).toHaveLength(4)
    expect(primaryIndex.value![0]!.timestamp).toBe(100n)
    expect(primaryIndex.value![1]!.timestamp).toBe(110n)
    expect(primaryIndex.value![2]!.timestamp).toBe(200n)
    expect(primaryIndex.value![3]!.timestamp).toBe(210n)
  })

  it('handles failed index loads gracefully', async () => {
    const errorSpy = vi.spyOn(console, 'error').mockImplementation(() => {})
    const cam0Index = createIndexBuffer([{ timestamp: 100n, offset: 0n }])

    mockedAxios.get.mockImplementation((url: string) => {
      if (url.includes('cam0')) return Promise.resolve({ data: cam0Index })
      return Promise.reject(new Error('network error'))
    })

    const baseUrl = ref('http://example.com/logs/')
    const entries = ref<ManifestEntry[]>([
      { type: 'vision_detection', path: 'test.cam0.idx', source: 'cam0' },
      { type: 'vision_detection', path: 'test.cam1.idx', source: 'cam1' },
    ])

    const { indices } = useLogFileMultiIndex(baseUrl, entries)

    await vi.waitFor(() => {
      expect(indices.value.size).toBe(1)
    })

    expect(indices.value.has('cam0')).toBe(true)
    expect(indices.value.has('cam1')).toBe(false)
    expect(errorSpy).toHaveBeenCalledWith('Failed to load index:', 'cam1', expect.any(Error))
    errorSpy.mockRestore()
  })
})

describe('useLogFileMessagesAtTimestamp', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('returns empty array for empty indices', async () => {
    const logUrl = ref('http://example.com/test.log')
    const indices = ref(new Map<string, IndexEntry[]>())
    const timestamp = ref(100n)

    const { messages } = useLogFileMessagesAtTimestamp(logUrl, indices, timestamp)

    await nextTick()
    await nextTick()

    expect(messages.value).toEqual([])
  })

  it('fetches messages from multiple camera indices', async () => {
    const payload0 = new Uint8Array([1, 2, 3])
    const payload1 = new Uint8Array([4, 5, 6])
    const msg0 = createMessageBuffer(100n, 1, payload0)
    const msg1 = createMessageBuffer(110n, 1, payload1)

    mockedAxios.get.mockImplementation(
      (_url: string, config?: import('axios').AxiosRequestConfig) => {
        const range = config?.headers?.range as string
        if (!range) return Promise.reject(new Error('no range'))

        // Parse the range header to determine which request this is
        const match = range.match(/bytes=(\d+)-(\d+)/)
        if (!match) return Promise.reject(new Error('bad range'))
        const start = Number(match[1])
        const end = Number(match[2])

        // Header request (16 bytes) vs full request
        if (end - start === 15) {
          // Header-only request - return based on offset
          if (start === 0) return Promise.resolve({ data: msg0.slice(0, 16) })
          if (start === 100) return Promise.resolve({ data: msg1.slice(0, 16) })
        } else {
          // Full message request
          if (start === 0) return Promise.resolve({ data: msg0 })
          if (start === 100) return Promise.resolve({ data: msg1 })
        }
        return Promise.reject(new Error('unknown range'))
      },
    )

    const cam0Index: IndexEntry[] = [{ timestamp: 100n, offset: 0n }]
    const cam1Index: IndexEntry[] = [{ timestamp: 110n, offset: 100n }]

    const logUrl = ref('http://example.com/test.log')
    const indices = ref(
      new Map<string, IndexEntry[]>([
        ['cam0', cam0Index],
        ['cam1', cam1Index],
      ]),
    )
    const timestamp = ref(200n)

    const { messages } = useLogFileMessagesAtTimestamp(logUrl, indices, timestamp)

    await vi.waitFor(() => {
      expect(messages.value).toHaveLength(2)
    })

    expect(messages.value[0]!.data.byteLength).toBe(3)
    expect(messages.value[1]!.data.byteLength).toBe(3)
  })

  it('skips cameras with no matching entry for timestamp', async () => {
    const payload = new Uint8Array([1, 2, 3])
    const msg = createMessageBuffer(200n, 1, payload)

    mockedAxios.get.mockImplementation(
      (_url: string, config?: import('axios').AxiosRequestConfig) => {
        const range = config?.headers?.range as string
        const match = range.match(/bytes=(\d+)-(\d+)/)
        if (!match) return Promise.reject(new Error('bad range'))
        const end = Number(match[1])
        const endB = Number(match[2])
        if (endB - end === 15) return Promise.resolve({ data: msg.slice(0, 16) })
        return Promise.resolve({ data: msg })
      },
    )

    const cam0Index: IndexEntry[] = [{ timestamp: 200n, offset: 0n }]
    const cam1Index: IndexEntry[] = [
      { timestamp: 300n, offset: 100n }, // all entries in the future
    ]

    const logUrl = ref('http://example.com/test.log')
    const indices = ref(
      new Map<string, IndexEntry[]>([
        ['cam0', cam0Index],
        ['cam1', cam1Index],
      ]),
    )
    const timestamp = ref(250n)

    const { messages } = useLogFileMessagesAtTimestamp(logUrl, indices, timestamp)

    await vi.waitFor(() => {
      expect(messages.value).toHaveLength(1)
    })
  })
})
