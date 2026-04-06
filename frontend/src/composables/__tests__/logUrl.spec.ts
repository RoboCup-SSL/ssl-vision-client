import { describe, it, expect, vi, beforeEach } from 'vitest'

const mockRoute = { query: {} as Record<string, string> }
const mockReplace = vi.fn()

vi.mock('vue-router', () => ({
  useRoute: () => mockRoute,
  useRouter: () => ({ replace: mockReplace }),
}))

const mockStorage = new Map<string, string>()
vi.stubGlobal('localStorage', {
  getItem: (key: string) => mockStorage.get(key) ?? null,
  setItem: (key: string, value: string) => mockStorage.set(key, value),
  removeItem: (key: string) => mockStorage.delete(key),
})

import { useLogUrl, getLogBase, getLogsBaseUrl } from '../logUrl'

describe('useLogUrl', () => {
  beforeEach(() => {
    mockRoute.query = {}
    mockReplace.mockClear()
    mockStorage.clear()
  })

  it('validates URL must end with .log', () => {
    const { logUrlInput, urlError } = useLogUrl()
    logUrlInput.value = 'http://example.com/test.txt'
    expect(urlError.value).toBe('URL must end with .log')
  })

  it('has no error for valid .log URL', () => {
    const { logUrlInput, urlError } = useLogUrl()
    logUrlInput.value = 'http://example.com/test.log'
    expect(urlError.value).toBe('')
  })

  it('has no error for empty input', () => {
    const { urlError } = useLogUrl()
    expect(urlError.value).toBe('')
  })

  it('returns valid URL for .log input', () => {
    const { logUrlInput, validLogUrl } = useLogUrl()
    logUrlInput.value = 'http://example.com/data/test.log'
    expect(validLogUrl.value).toBe('http://example.com/data/test.log')
  })

  it('returns empty validLogUrl for invalid input', () => {
    const { logUrlInput, validLogUrl } = useLogUrl()
    logUrlInput.value = 'http://example.com/test.txt'
    expect(validLogUrl.value).toBe('')
  })
})

describe('getLogBase', () => {
  it('removes .log extension', () => {
    expect(getLogBase('http://example.com/data/test.log')).toBe('http://example.com/data/test')
  })

  it('returns empty for empty input', () => {
    expect(getLogBase('')).toBe('')
  })
})

describe('getLogsBaseUrl', () => {
  it('extracts directory part of URL', () => {
    expect(getLogsBaseUrl('http://example.com/data/test.log')).toBe('http://example.com/data/')
  })

  it('returns empty for empty input', () => {
    expect(getLogsBaseUrl('')).toBe('')
  })

  it('returns empty for URL without slash', () => {
    expect(getLogsBaseUrl('test.log')).toBe('')
  })
})
