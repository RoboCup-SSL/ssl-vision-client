import axios from 'axios'
import { computed, toValue } from 'vue'
import type { MaybeRefOrGetter } from 'vue'
import { computedAsync } from '@vueuse/core'
import { fromBinary } from '@bufbuild/protobuf'
import type { DescMessage, MessageShape } from '@bufbuild/protobuf'

export type LogFileMessage = {
  timestamp: bigint
  messageType: number
  size: number
  data: ArrayBuffer
}

export type IndexEntry = {
  timestamp: bigint
  offset: bigint
}

export type ManifestEntry = {
  type: 'refbox' | 'vision_detection' | 'vision_geometry' | 'tracker'
  path: string
  source?: string
}

export type LogManifest = {
  indices: ManifestEntry[]
}

export const parseProtobuf = <Desc extends DescMessage>(
  message: LogFileMessage | undefined,
  schema: Desc,
  label: string,
): MessageShape<Desc> | undefined => {
  if (!message) return undefined
  try {
    return fromBinary(schema, new Uint8Array(message.data))
  } catch (e) {
    console.error(`Failed to parse ${label}:`, e)
    return undefined
  }
}

const queryLogFileData = async (
  url: string,
  startByte: bigint,
  endByte: bigint,
  signal?: AbortSignal,
): Promise<ArrayBuffer> => {
  const rangeHeader = `bytes=${startByte}-${endByte}`
  const response = await axios.get(url, {
    responseType: 'arraybuffer',
    headers: { range: rangeHeader },
    signal,
  })
  return response.data
}

const loadBinaryIndex = async (url: string): Promise<IndexEntry[]> => {
  if (!url) {
    return []
  }

  const response = await axios.get(url, {
    responseType: 'arraybuffer',
  })
  const buffer = response.data

  // Validate buffer size
  if (!buffer || buffer.byteLength === 0) {
    return []
  }

  // Each entry is 16 bytes (8 bytes timestamp + 8 bytes offset)
  if (buffer.byteLength % 16 !== 0) {
    console.error(`Invalid index file size: ${buffer.byteLength} bytes (not divisible by 16)`)
    return []
  }

  const dataView = new DataView(buffer)
  const numEntries = Math.floor(buffer.byteLength / 16)

  // Validate numEntries is reasonable (prevent creating huge arrays)
  if (numEntries < 0 || numEntries > 10_000_000) {
    console.error(`Invalid number of entries: ${numEntries}`)
    return []
  }

  const entries: IndexEntry[] = new Array(numEntries)

  for (let i = 0; i < numEntries; i++) {
    const offset = i * 16
    const timestamp = dataView.getBigInt64(offset, false) // big endian
    const fileOffset = dataView.getBigInt64(offset + 8, false) // big endian
    entries[i] = { timestamp, offset: fileOffset }
  }

  return entries
}

const queryLogfileMessage = async (
  url: string,
  offset: bigint,
  signal?: AbortSignal,
): Promise<LogFileMessage> => {
  // Read header first to get the size
  const headerData = await queryLogFileData(url, offset, offset + BigInt(15), signal)
  const headerView = new DataView(headerData)
  const timestamp = headerView.getBigInt64(0)
  const messageType = headerView.getInt32(8)
  const size = headerView.getInt32(12)

  // Now read the full message including payload
  const fullData = await queryLogFileData(url, offset, offset + BigInt(16 + size - 1), signal)
  const payload = fullData.slice(16, 16 + size)

  return {
    timestamp,
    messageType,
    size,
    data: payload,
  }
}

export const useLogFileIndex = (indexUrl: MaybeRefOrGetter<string>) => {
  const index = computedAsync(async () => {
    try {
      const url = toValue(indexUrl)
      if (!url) {
        return []
      }
      return await loadBinaryIndex(url)
    } catch (error) {
      console.error('Failed to load index file:', error)
      return []
    }
  }, [])

  return { index }
}

const loadManifest = async (url: string): Promise<LogManifest> => {
  if (!url) return { indices: [] }
  const response = await axios.get<LogManifest>(url)
  return response.data
}

export const useLogFileManifest = (manifestUrl: MaybeRefOrGetter<string>) => {
  const manifest = computedAsync(
    async () => {
      try {
        const url = toValue(manifestUrl)
        if (!url) return { indices: [] } as LogManifest
        return await loadManifest(url)
      } catch (error) {
        console.error('Failed to load manifest:', error)
        return { indices: [] } as LogManifest
      }
    },
    { indices: [] } as LogManifest,
  )

  return { manifest }
}

export const findClosestIndex = (index: IndexEntry[], timestamp: bigint): number => {
  if (index.length === 0) return 0
  if (timestamp < index[0]!.timestamp) return 0

  let lo = 0
  let hi = index.length - 1
  while (lo < hi) {
    const mid = lo + Math.ceil((hi - lo) / 2)
    if (index[mid]!.timestamp <= timestamp) {
      lo = mid
    } else {
      hi = mid - 1
    }
  }
  return lo
}

export const findClosestEntry = (
  index: IndexEntry[],
  timestamp: bigint,
): IndexEntry | undefined => {
  if (index.length === 0 || timestamp < index[0]!.timestamp) return undefined
  return index[findClosestIndex(index, timestamp)]
}

export const useLogFileMessageAtTimestamp = (
  logUrl: MaybeRefOrGetter<string>,
  index: MaybeRefOrGetter<IndexEntry[] | undefined>,
  targetTimestamp: MaybeRefOrGetter<bigint>,
) => {
  const message = computedAsync(async (onCancel) => {
    const indexValue = toValue(index)
    const timestamp = toValue(targetTimestamp)
    const url = toValue(logUrl)

    if (!indexValue || indexValue.length === 0) {
      return undefined
    }

    const closestEntry = findClosestEntry(indexValue, timestamp)
    if (!closestEntry) {
      return undefined
    }

    const abort = new AbortController()
    onCancel(() => abort.abort())

    try {
      return await queryLogfileMessage(url, closestEntry.offset, abort.signal)
    } catch (e) {
      if (!abort.signal.aborted) {
        console.error('Failed to fetch message:', e)
      }
      return undefined
    }
  }, undefined)

  return { message }
}

export const useLogFileMultiIndex = (
  baseUrl: MaybeRefOrGetter<string>,
  manifestEntries: MaybeRefOrGetter<ManifestEntry[]>,
) => {
  const indices = computedAsync(async () => {
    const entries = toValue(manifestEntries)
    const base = toValue(baseUrl)
    if (!base || entries.length === 0) return new Map<string, IndexEntry[]>()
    const result = new Map<string, IndexEntry[]>()
    for (const entry of entries) {
      const url = `${base}${entry.path}`
      try {
        result.set(entry.source ?? entry.path, await loadBinaryIndex(url))
      } catch (e) {
        console.error('Failed to load index:', entry.source, e)
      }
    }
    return result
  }, new Map<string, IndexEntry[]>())

  const primaryIndex = computed(() => {
    const map = indices.value
    if (map.size === 0) return undefined
    const arrays = Array.from(map.values())
    const merged: IndexEntry[] = arrays.flat()
    merged.sort((a, b) => (a.timestamp < b.timestamp ? -1 : a.timestamp > b.timestamp ? 1 : 0))
    return merged
  })

  return { indices, primaryIndex }
}

export const useLogFileMessagesAtTimestamp = (
  logUrl: MaybeRefOrGetter<string>,
  indices: MaybeRefOrGetter<Map<string, IndexEntry[]>>,
  targetTimestamp: MaybeRefOrGetter<bigint>,
) => {
  const messages = computedAsync(async (onCancel) => {
    const map = toValue(indices)
    const timestamp = toValue(targetTimestamp)
    const url = toValue(logUrl)

    if (!url || map.size === 0) return []

    const abort = new AbortController()
    onCancel(() => abort.abort())

    const results: LogFileMessage[] = []
    for (const [, index] of map) {
      if (!index || index.length === 0) continue
      const entry = findClosestEntry(index, timestamp)
      if (!entry) continue
      try {
        results.push(await queryLogfileMessage(url, entry.offset, abort.signal))
      } catch (e) {
        if (!abort.signal.aborted) {
          console.error('Failed to fetch message for source:', e)
        }
      }
    }
    return results
  }, [])

  return { messages }
}
