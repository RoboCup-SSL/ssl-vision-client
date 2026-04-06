import { describe, it, expect } from 'vitest'
import { ref, nextTick } from 'vue'
import { useLogPlayerSources } from '../logPlayerSources'
import type { ManifestEntry } from '../logfile'
import { SOURCE_VISION } from '../vision'

describe('useLogPlayerSources', () => {
  it('defaults activeSource to vision', () => {
    const entries = ref<ManifestEntry[]>([])
    const { activeSource } = useLogPlayerSources(entries)
    expect(activeSource.value).toBe(SOURCE_VISION)
  })

  it('builds sources map from tracker manifest entries', () => {
    const entries = ref<ManifestEntry[]>([
      { type: 'tracker', path: 'a.idx', source: 'tracker_a' },
      { type: 'tracker', path: 'b.idx', source: 'tracker_b' },
    ])
    const { sources } = useLogPlayerSources(entries)
    expect(sources.value).toEqual({
      [SOURCE_VISION]: SOURCE_VISION,
      tracker_a: 'tracker_a',
      tracker_b: 'tracker_b',
    })
  })

  it('auto-selects first tracker source when entries arrive', async () => {
    const entries = ref<ManifestEntry[]>([])
    const { activeSource } = useLogPlayerSources(entries)
    expect(activeSource.value).toBe(SOURCE_VISION)

    entries.value = [
      { type: 'tracker', path: 'a.idx', source: 'tracker_a' },
      { type: 'tracker', path: 'b.idx', source: 'tracker_b' },
    ]
    await nextTick()
    expect(activeSource.value).toBe('tracker_a')
  })

  it('allows switching between sources', () => {
    const entries = ref<ManifestEntry[]>([{ type: 'tracker', path: 'a.idx', source: 'tracker_a' }])
    const { activeSource } = useLogPlayerSources(entries)

    activeSource.value = 'tracker_a'
    expect(activeSource.value).toBe('tracker_a')

    activeSource.value = SOURCE_VISION
    expect(activeSource.value).toBe(SOURCE_VISION)
  })
})
