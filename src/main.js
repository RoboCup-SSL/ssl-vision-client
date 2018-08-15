import Vue from "vue";
import App from "./App.vue";
import store from "./store";
import VueNativeSock from 'vue-native-websocket'

Vue.config.productionTip = false;

let wsAddress;
if (process.env.NODE_ENV === 'development') {
    // use the default backend port
    wsAddress = 'ws://localhost:8082/api/vision';
} else {
    // UI and backend are served on the same host+port on production builds
    wsAddress = 'ws://' + window.location.hostname + ':' + window.location.port + '/api/vision';
}

Vue.use(VueNativeSock, wsAddress, {
    reconnection: true,
    format: 'json',
    store: store,
});

new Vue({
    render: h => h(App),
    store,
}).$mount("#app");
