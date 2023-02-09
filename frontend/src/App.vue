<script setup lang="ts">
import FieldCanvas from "@/components/FieldCanvas.vue";
import {defaultField} from "@/field";
import {ref} from "vue";
import {connect} from "@/websocket";

const field = ref(defaultField)
connect((f: any) => field.value = f)
</script>

<template>
  <FieldCanvas :field="field"/>
</template>
