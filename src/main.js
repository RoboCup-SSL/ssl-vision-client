import Vue from "vue";
import App from "./App.vue";
import store from "./store";
import VueNativeSock from 'vue-native-websocket'

Vue.config.productionTip = false;

const urlParams = new URLSearchParams(window.location.search);
const sourceId = urlParams.get('sourceId');

let wsAddress;
if (process.env.NODE_ENV === 'development') {
    // use the default backend port
    wsAddress = 'ws://localhost:8082/api/vision';
} else {
    // UI and backend are served on the same host+port on production builds
    let protocol;
    if (window.location.protocol === 'http:') {
        protocol = 'ws:'
    } else {
        protocol = 'wss:'
    }
    wsAddress = protocol + '//' + window.location.hostname + ':' + window.location.port + '/api/vision';
}

if (sourceId) {
    wsAddress += `?sourceId=${sourceId}`
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
