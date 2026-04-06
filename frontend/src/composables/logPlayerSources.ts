import { computed, ref, toValue, watch } from 'vue'
import type { MaybeRefOrGetter } from 'vue'
import type { ManifestEntry } from '@/composables/logfile.ts'
import { SOURCE_VISION } from '@/composables/vision.ts'

export const useLogPlayerSources = (trackerManifestEntries: MaybeRefOrGetter<ManifestEntry[]>) => {
  const activeSource = ref(SOURCE_VISION)

  watch(
    () => toValue(trackerManifestEntries),
    (entries) => {
      if (entries.length > 0 && activeSource.value === SOURCE_VISION) {
        activeSource.value = entries[0]!.source ?? SOURCE_VISION
      }
    },
  )

  const sources = computed(() => {
    const s: Record<string, string> = { [SOURCE_VISION]: SOURCE_VISION }
    for (const entry of toValue(trackerManifestEntries)) {
      if (entry.source) {
        s[entry.source] = entry.source
      }
    }
    return s
  })

  return {
    activeSource,
    sources,
  }
}
