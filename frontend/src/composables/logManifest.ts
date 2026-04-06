import { computed, toValue } from 'vue'
import type { MaybeRefOrGetter } from 'vue'
import { useLogFileManifest } from '@/composables/logfile.ts'
import type { ManifestEntry } from '@/composables/logfile.ts'
import { getLogBase, getLogsBaseUrl } from '@/composables/logUrl.ts'

export const useLogManifest = (logUrl: MaybeRefOrGetter<string>) => {
  const { manifest } = useLogFileManifest(
    computed(() => {
      const url = toValue(logUrl)
      const base = getLogBase(url)
      return base ? `${base}.log.manifest.json` : ''
    }),
  )

  const findManifestEntry = (type: string, source?: string): ManifestEntry | undefined => {
    return manifest.value.indices.find(
      (e) => e.type === type && (source === undefined || e.source === source),
    )
  }

  const manifestIndexUrl = (type: string, source?: string): string => {
    const entry = findManifestEntry(type, source)
    if (!entry) return ''
    return `${getLogsBaseUrl(toValue(logUrl))}${entry.path}`
  }

  const manifestEntriesOfType = (type: string) =>
    computed(() => manifest.value.indices.filter((e) => e.type === type))

  return {
    manifest,
    findManifestEntry,
    manifestIndexUrl,
    manifestEntriesOfType,
  }
}
