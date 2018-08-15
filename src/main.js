import Vue from "vue";
import App from "./App.vue";
import store from "./store";
import VueNativeSock from 'vue-native-websocket'

Vue.config.productionTip = false;

Vue.use(VueNativeSock, 'ws://localhost:8081/api/vision', {
    reconnection: true,
    format: 'json',
    store: store,
});

new Vue({
    render: h => h(App),
    store,
}).$mount("#app");
