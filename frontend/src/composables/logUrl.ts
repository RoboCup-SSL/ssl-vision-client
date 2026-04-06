import { computed, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const STORAGE_URL_KEY = 'ssl-vision-log-url'

export const getLogBase = (url: string): string => {
  if (!url) return ''
  return url.slice(0, -'.log'.length)
}

export const getLogsBaseUrl = (url: string): string => {
  if (!url) return ''
  const lastSlash = url.lastIndexOf('/')
  return lastSlash >= 0 ? url.substring(0, lastSlash + 1) : ''
}

export const useLogUrl = () => {
  const route = useRoute()
  const router = useRouter()

  const logUrlInput = ref('')

  const urlError = computed(() => {
    const url = logUrlInput.value.trim()
    if (url && !url.endsWith('.log')) return 'URL must end with .log'
    return ''
  })

  watch(
    () => route.query.url,
    (urlParam) => {
      if (typeof urlParam === 'string' && urlParam && !logUrlInput.value) {
        logUrlInput.value = urlParam
      } else if (!urlParam && !logUrlInput.value) {
        const savedUrl = localStorage.getItem(STORAGE_URL_KEY)
        if (savedUrl) {
          logUrlInput.value = savedUrl
          router.replace({ query: { ...route.query, url: savedUrl } })
        }
      }
    },
    { immediate: true },
  )

  watch(logUrlInput, (url) => {
    router.replace({ query: { ...route.query, url: url || undefined } })
    if (url && url.trim()) {
      localStorage.setItem(STORAGE_URL_KEY, url)
    } else {
      localStorage.removeItem(STORAGE_URL_KEY)
    }
  })

  const validLogUrl = computed(() => {
    const url = logUrlInput.value.trim()
    if (!url || !url.endsWith('.log')) return ''
    return url
  })

  return {
    logUrlInput,
    urlError,
    validLogUrl,
  }
}
