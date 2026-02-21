import axios from 'axios'
import { toValue } from 'vue'
import type { MaybeRefOrGetter } from 'vue'
import { computedAsync } from '@vueuse/core'

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

const queryLogFileData = async (
  url: string,
  startByte: bigint,
  endByte: bigint,
): Promise<ArrayBuffer> => {
  const rangeHeader = `bytes=${startByte}-${endByte}`
  const response = await axios.get(url, {
    responseType: 'arraybuffer',
    headers: { range: rangeHeader },
  })
  return response.data
}

const loadBinaryIndex = async (url: string): Promise<IndexEntry[]> => {
  const response = await axios.get(url, {
    responseType: 'arraybuffer',
  })
  const buffer = response.data
  const dataView = new DataView(buffer)

  const numEntries = buffer.byteLength / 16
  const entries: IndexEntry[] = new Array(numEntries)

  for (let i = 0; i < numEntries; i++) {
    const offset = i * 16
    const timestamp = dataView.getBigInt64(offset, false) // big endian
    const fileOffset = dataView.getBigInt64(offset + 8, false) // big endian
    entries[i] = { timestamp, offset: fileOffset }
  }

  return entries
}

const queryLogfileMessage = async (url: string, offset: bigint): Promise<LogFileMessage> => {
  // Read header first to get the size
  const headerData = await queryLogFileData(url, offset, offset + BigInt(15))
  const headerView = new DataView(headerData)
  const timestamp = headerView.getBigInt64(0)
  const messageType = headerView.getInt32(8)
  const size = headerView.getInt32(12)

  // Now read the full message including payload
  const fullData = await queryLogFileData(url, offset, offset + BigInt(16 + size - 1))
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
    return await loadBinaryIndex(toValue(indexUrl))
  }, [])

  return { index }
}

export const useLogFileMessageAtTimestamp = (
  logUrl: MaybeRefOrGetter<string>,
  index: MaybeRefOrGetter<IndexEntry[] | undefined>,
  targetTimestamp: MaybeRefOrGetter<bigint>,
) => {
  const message = computedAsync(async () => {
    const indexValue = toValue(index)
    const timestamp = toValue(targetTimestamp)
    const url = toValue(logUrl)

    if (!indexValue || indexValue.length === 0) {
      return undefined
    }

    // Find closest entry
    const firstEntry = indexValue[0]
    if (!firstEntry) {
      return undefined
    }

    let closestEntry = firstEntry
    let minDiff = timestamp > closestEntry.timestamp
      ? timestamp - closestEntry.timestamp
      : closestEntry.timestamp - timestamp

    for (const entry of indexValue) {
      const diff = timestamp > entry.timestamp
        ? timestamp - entry.timestamp
        : entry.timestamp - timestamp

      if (diff < minDiff) {
        minDiff = diff
        closestEntry = entry
      }
    }

    return await queryLogfileMessage(url, closestEntry.offset)
  }, undefined)

  return { message }
}



