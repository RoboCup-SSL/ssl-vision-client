<script setup lang="ts">
defineProps<{
  modelValue: string
  urlError: string
}>()

const emit = defineEmits<{
  'update:modelValue': [value: string]
}>()
</script>

<template>
  <div class="url-input-container">
    <input
      :value="modelValue"
      @input="emit('update:modelValue', ($event.target as HTMLInputElement).value)"
      type="text"
      placeholder="Enter log file URL (must end with .log)"
      class="url-input"
      :class="{ 'url-input-error': urlError }"
    />
    <span v-if="urlError" class="url-error">{{ urlError }}</span>
  </div>
</template>

<style scoped>
.url-input-container {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.url-input {
  width: 100%;
  padding: 6px 10px;
  font-size: 14px;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
}

.url-input-error {
  border-color: #e53e3e;
}

.url-error {
  font-size: 12px;
  color: #e53e3e;
}
</style>
