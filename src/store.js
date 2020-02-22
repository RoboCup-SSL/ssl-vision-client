import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

let fieldWidth = 9000;
let fieldLength = 12000;
let centerCircleRadius = 500;

let defaultField = {
    fieldWidth: fieldWidth,
    fieldLength: fieldLength,
    boundaryWidth: 300,
    penAreaWidth: 1000,
    penAreaDepth: 500,
    goalWidth: 600,
    goalDepth: 180,
    centerCircleRadius: centerCircleRadius,
    ballRadius: 21.5,
    shapes: [
        {
            line: {
                p1: {x: -fieldLength / 2, y: -fieldWidth / 2},
                p2: {x: -fieldLength / 2, y: fieldWidth / 2}
            },
        },
        {

            line: {
                p1: {x: -fieldLength / 2, y: -fieldWidth / 2},
                p2: {x: -fieldLength / 2, y: fieldWidth / 2}
            }
        },
        {
            line: {
                p1: {x: -fieldLength / 2, y: fieldWidth / 2},
                p2: {x: fieldLength / 2, y: fieldWidth / 2}
            }
        },
        {
            line: {
                p1: {x: fieldLength / 2, y: fieldWidth / 2},
                p2: {x: fieldLength / 2, y: -fieldWidth / 2}
            }
        },
        {
            line: {
                p1: {x: fieldLength / 2, y: -fieldWidth / 2},
                p2: {x: -fieldLength / 2, y: -fieldWidth / 2}
            }
        },
        {
            line: {
                p1: {x: 0, y: fieldWidth / 2},
                p2: {x: 0, y: -fieldWidth / 2}
            }
        },
        {
            circle: {
                center: {x: 0, y: 0},
                radius: centerCircleRadius,
                stroke: 'white',
                strokeWidth: 10,
                fill: '',
                fillOpacity: 0
            }
        },
        {
            path: {
                d: [
                    {
                        type: 'M',
                        args: [925, -550]
                    },
                    {
                        type: 'A',
                        args: [90, 90, 0, 1, 1, 925, -450]
                    },
                    {
                        type: 'L',
                        args: [925, -550]
                    }
                ],
                stroke: 'black',
                strokeWidth: '10',
                fill: 'yellow',
                fillOpacity: 1
            }
        },
        {
            text: {
                text: '1',
                p: {x: 990, y: -510},
                fill: 'black',
            }
        }
    ],
};

export default new Vuex.Store({
    state: {
        field: defaultField
    },
    mutations: {
        SOCKET_ONOPEN() {
        },
        SOCKET_ONCLOSE() {
        },
        SOCKET_ONERROR() {
        },
        SOCKET_ONMESSAGE(state, message) {
            state.field = message;
        },
        SOCKET_RECONNECT() {
        },
        SOCKET_RECONNECT_ERROR() {
        },
    }
});
