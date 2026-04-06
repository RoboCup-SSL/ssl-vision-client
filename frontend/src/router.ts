import { createRouter, createWebHashHistory } from 'vue-router'

import FieldVisualizerVision from './views/FieldVisualizerVision.vue'
import VisualizerPlayer from './views/VisualizerPlayer.vue'

const routes = [
  { path: '/', component: FieldVisualizerVision },
  { path: '/player', component: VisualizerPlayer },
]

export const router = createRouter({
  history: createWebHashHistory(),
  routes,
})
