<script setup lang="ts">
const props = defineProps<{
  sources: { [key: string]: string }
}>()

const activeSource = defineModel()


const updateActiveSource = (sourceId: string) => {
  activeSource.value = sourceId
}
</script>

<template>
  <div id="source-selector">
    Sources:
    <template v-for="[sourceId, sourceName] of Object.entries(props.sources)" :key="'input-' + sourceId">
      <input
        type="radio"
        :id="sourceId"
        :value="sourceId"
        name="source-selector"
        :checked="sourceId === activeSource"
        @click="updateActiveSource(sourceId)"
      />
      <label :for="sourceId">
        {{ sourceName }}
      </label>
    </template>
  </div>
</template>

<style scoped>
#source-selector {
  align-content: center;
  width: 100%;
  bottom: 0.1em;
  position: absolute;
}
</style>
