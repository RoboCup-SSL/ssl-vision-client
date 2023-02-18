<script setup lang="ts">
import FieldCanvas from "@/components/FieldCanvas.vue";
import SourceSelector from "@/components/SourceSelector.vue";
import {inject, ref} from "vue";
import {defaultField} from "@/field";
import type {VisionApi} from "@/providers/backend/VisionApi";
import type {Field} from "@/field";

const visionApi = inject<VisionApi>('vision-api')
const field = ref(defaultField)
visionApi?.RegisterConsumer((f: Field) => field.value = f)
</script>

<template>
  <SourceSelector :sources="field.sources" :active-source="field.activeSourceId"/>
  <FieldCanvas :field="field"/>
</template>
