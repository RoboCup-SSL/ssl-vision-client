import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { defineComponent, ref, nextTick } from 'vue'
import type { MaybeRefOrGetter } from 'vue'
import { mount } from '@vue/test-utils'
import type { IndexEntry } from '../logfile'

const mockRoute = { query: {} as Record<string, string> }

vi.mock('vue-router', () => ({
  useRoute: () => mockRoute,
}))

const mockStorage = new Map<string, string>()
vi.stubGlobal('localStorage', {
  getItem: (key: string) => mockStorage.get(key) ?? null,
  setItem: (key: string, value: string) => mockStorage.set(key, value),
  removeItem: (key: string) => mockStorage.delete(key),
})

import { useLogPlayerPosition } from '../logPlayerPosition'

function withSetup(
  globalMinTimestamp: MaybeRefOrGetter<bigint>,
  globalMaxTimestamp: MaybeRefOrGetter<bigint>,
  primaryIndex: MaybeRefOrGetter<IndexEntry[] | undefined>,
  logUrl: MaybeRefOrGetter<string>,
) {
  let result!: ReturnType<typeof useLogPlayerPosition>
  const wrapper = mount(
    defineComponent({
      setup() {
        result = useLogPlayerPosition(globalMinTimestamp, globalMaxTimestamp, primaryIndex, logUrl)
        return () => {}
      },
    }),
  )
  return { result, wrapper }
}

describe('useLogPlayerPosition', () => {
  let wrapper: ReturnType<typeof mount> | undefined

  beforeEach(() => {
    mockRoute.query = {}
    mockStorage.clear()
  })

  afterEach(() => {
    wrapper?.unmount()
    wrapper = undefined
  })

  const makeIndex = (timestamps: bigint[]): IndexEntry[] =>
    timestamps.map((ts, i) => ({ timestamp: ts, offset: BigInt(i * 100) }))

  it('starts with zero timestamp', () => {
    const { result, wrapper: w } = withSetup(
      ref(BigInt(0)),
      ref(BigInt(0)),
      ref(undefined),
      ref('http://example.com/test.log'),
    )
    wrapper = w
    expect(result.currentTimestamp.value).toBe(BigInt(0))
    expect(result.sliderValue.value).toBe(0)
    expect(result.sliderMax.value).toBe(0)
  })

  it('computes sliderMax from global timestamp range', () => {
    const { result, wrapper: w } = withSetup(
      ref(100n),
      ref(500n),
      ref(undefined),
      ref('http://example.com/test.log'),
    )
    wrapper = w
    expect(result.sliderMax.value).toBe(400)
  })

  it('maps sliderValue to currentTimestamp', () => {
    const { result, wrapper: w } = withSetup(
      ref(1000n),
      ref(5000n),
      ref(undefined),
      ref('http://example.com/test.log'),
    )
    wrapper = w

    result.sliderValue.value = 2000
    expect(result.currentTimestamp.value).toBe(3000n)
  })

  it('maps currentTimestamp to sliderValue', () => {
    const { result, wrapper: w } = withSetup(
      ref(1000n),
      ref(5000n),
      ref(undefined),
      ref('http://example.com/test.log'),
    )
    wrapper = w

    result.currentTimestamp.value = 3000n
    expect(result.sliderValue.value).toBe(2000)
  })

  it('clamps sliderValue to valid range', () => {
    const { result, wrapper: w } = withSetup(
      ref(100n),
      ref(500n),
      ref(undefined),
      ref('http://example.com/test.log'),
    )
    wrapper = w

    // Set timestamp beyond max
    result.currentTimestamp.value = 9999n
    expect(result.sliderValue.value).toBe(400)

    // Set timestamp below min
    result.currentTimestamp.value = 0n
    expect(result.sliderValue.value).toBe(0)
  })

  it('restores timestamp from localStorage', async () => {
    mockStorage.set('ssl-vision-timestamp', '300')
    const globalMax = ref(BigInt(0))
    const { result, wrapper: w } = withSetup(
      ref(100n),
      globalMax,
      ref(undefined),
      ref('http://example.com/test.log'),
    )
    wrapper = w

    // Simulate indices loading
    globalMax.value = 500n
    await nextTick()
    expect(result.currentTimestamp.value).toBe(300n)
  })

  it('restores timestamp from query param', async () => {
    mockRoute.query = { ts: '250' }
    const globalMax = ref(BigInt(0))
    const { result, wrapper: w } = withSetup(
      ref(100n),
      globalMax,
      ref(undefined),
      ref('http://example.com/test.log'),
    )
    wrapper = w

    globalMax.value = 500n
    await nextTick()
    expect(result.currentTimestamp.value).toBe(250n)
  })

  it('does not crash on invalid BigInt in query param', async () => {
    mockRoute.query = { ts: 'not-a-number' }
    const globalMax = ref(BigInt(0))
    const { result, wrapper: w } = withSetup(
      ref(100n),
      globalMax,
      ref(undefined),
      ref('http://example.com/test.log'),
    )
    wrapper = w

    globalMax.value = 500n
    await nextTick()
    expect(result.currentTimestamp.value).toBe(BigInt(0))
  })

  it('persists timestamp to localStorage', async () => {
    const { result, wrapper: w } = withSetup(
      ref(100n),
      ref(500n),
      ref(undefined),
      ref('http://example.com/test.log'),
    )
    wrapper = w

    result.currentTimestamp.value = 300n
    await nextTick()
    expect(mockStorage.get('ssl-vision-timestamp')).toBe('300')
  })

  it('resets timestamp when logUrl changes', async () => {
    const logUrl = ref('http://example.com/test.log')
    const { result, wrapper: w } = withSetup(ref(100n), ref(500n), ref(undefined), logUrl)
    wrapper = w

    result.currentTimestamp.value = 300n
    await nextTick()

    logUrl.value = 'http://example.com/other.log'
    await nextTick()
    expect(result.currentTimestamp.value).toBe(BigInt(0))
  })

  it('steps forward through primaryIndex frames', () => {
    const index = ref<IndexEntry[] | undefined>(makeIndex([100n, 200n, 300n, 400n]))
    const { result, wrapper: w } = withSetup(
      ref(100n),
      ref(400n),
      index,
      ref('http://example.com/test.log'),
    )
    wrapper = w

    result.currentTimestamp.value = 200n
    result.stepForward()
    expect(result.currentTimestamp.value).toBe(300n)

    result.stepForward()
    expect(result.currentTimestamp.value).toBe(400n)

    // At end, stays at last
    result.stepForward()
    expect(result.currentTimestamp.value).toBe(400n)
  })

  it('steps backward through primaryIndex frames', () => {
    const index = ref<IndexEntry[] | undefined>(makeIndex([100n, 200n, 300n, 400n]))
    const { result, wrapper: w } = withSetup(
      ref(100n),
      ref(400n),
      index,
      ref('http://example.com/test.log'),
    )
    wrapper = w

    result.currentTimestamp.value = 300n
    result.stepBackward()
    expect(result.currentTimestamp.value).toBe(200n)

    result.stepBackward()
    expect(result.currentTimestamp.value).toBe(100n)

    // At start, stays at first
    result.stepBackward()
    expect(result.currentTimestamp.value).toBe(100n)
  })

  it('steps to nearest frame when between frames', () => {
    const index = ref<IndexEntry[] | undefined>(makeIndex([100n, 200n, 300n]))
    const { result, wrapper: w } = withSetup(
      ref(100n),
      ref(300n),
      index,
      ref('http://example.com/test.log'),
    )
    wrapper = w

    // Between 100n and 200n, floor semantics → at index 0 (100n), step to index 1 (200n)
    result.currentTimestamp.value = 150n
    result.stepForward()
    expect(result.currentTimestamp.value).toBe(200n)
  })
})
