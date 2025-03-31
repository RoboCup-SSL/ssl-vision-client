import { createApp } from 'vue'
import App from './App.vue'

import './assets/main.css'
import { VisionApi } from '@/providers/backend/VisionApi'

createApp(App).provide('vision-api', new VisionApi()).mount('#app')
