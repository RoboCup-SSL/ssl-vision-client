import { describe, it, expect, vi, beforeEach } from 'vitest'
import { ref } from 'vue'

vi.mock('axios', () => ({
  default: {
    get: vi.fn(),
  },
}))

import axios from 'axios'
import { useLogManifest } from '../logManifest'
import type { LogManifest } from '../logfile'

const mockedAxios = vi.mocked(axios)

describe('useLogManifest', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('finds manifest entry by type', async () => {
    const metadata: LogManifest = {
      indices: [
        { type: 'vision_detection', path: 'det.idx' },
        { type: 'vision_geometry', path: 'geo.idx' },
        { type: 'refbox', path: 'ref.idx' },
      ],
    }

    mockedAxios.get.mockResolvedValue({ data: metadata })

    const logUrl = ref('http://example.com/test.log')
    const manifest = useLogManifest(logUrl)

    await vi.waitFor(() => {
      expect(manifest.manifest.value.indices.length).toBe(3)
    })

    expect(manifest.findManifestEntry('vision_geometry')).toEqual({
      type: 'vision_geometry',
      path: 'geo.idx',
    })
    expect(manifest.findManifestEntry('nonexistent')).toBeUndefined()
  })

  it('constructs manifest index URL', async () => {
    const metadata: LogManifest = {
      indices: [{ type: 'refbox', path: 'ref.idx' }],
    }

    mockedAxios.get.mockResolvedValue({ data: metadata })

    const logUrl = ref('http://example.com/test.log')
    const manifest = useLogManifest(logUrl)

    await vi.waitFor(() => {
      expect(manifest.manifest.value.indices.length).toBe(1)
    })

    expect(manifest.manifestIndexUrl('refbox')).toBe('http://example.com/ref.idx')
    expect(manifest.manifestIndexUrl('nonexistent')).toBe('')
  })

  it('filters manifest entries by type', async () => {
    const metadata: LogManifest = {
      indices: [
        { type: 'vision_detection', path: 'cam0.idx', source: 'cam0' },
        { type: 'vision_detection', path: 'cam1.idx', source: 'cam1' },
        { type: 'refbox', path: 'ref.idx' },
      ],
    }

    mockedAxios.get.mockResolvedValue({ data: metadata })

    const logUrl = ref('http://example.com/test.log')
    const manifest = useLogManifest(logUrl)

    await vi.waitFor(() => {
      expect(manifest.manifest.value.indices.length).toBe(3)
    })

    const detectionEntries = manifest.manifestEntriesOfType('vision_detection')
    expect(detectionEntries.value).toHaveLength(2)
    expect(detectionEntries.value[0]!.source).toBe('cam0')
  })

  it('finds manifest entry by type and source', async () => {
    const metadata: LogManifest = {
      indices: [
        { type: 'tracker', path: 'tracker_a.idx', source: 'tracker_a' },
        { type: 'tracker', path: 'tracker_b.idx', source: 'tracker_b' },
      ],
    }

    mockedAxios.get.mockResolvedValue({ data: metadata })

    const logUrl = ref('http://example.com/test.log')
    const manifest = useLogManifest(logUrl)

    await vi.waitFor(() => {
      expect(manifest.manifest.value.indices.length).toBe(2)
    })

    expect(manifest.findManifestEntry('tracker', 'tracker_b')?.path).toBe('tracker_b.idx')
    expect(manifest.findManifestEntry('tracker', 'nonexistent')).toBeUndefined()
  })
})
