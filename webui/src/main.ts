/**
 * main.ts
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Components
import App from './App.vue'
import { createPinia } from 'pinia'

// Composables
import { createApp } from 'vue'

import VNetworkGraph from "v-network-graph"
import "v-network-graph/lib/style.css"

// Router
import router from "./router";

// Plugins
import { registerPlugins } from '@/plugins'

const pinia = createPinia()
const app = createApp(App)

app.use(pinia)
app.use(VNetworkGraph)

registerPlugins(app)

app.mount('#app')
