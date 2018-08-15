<template>
    <svg class="field-canvas"
         :viewBox="viewBox">

        <!-- rotate field -->
        <g :transform="getFieldTransformation">

            <!-- draw field background -->
            <rect :x="-(field.fieldLength/2+field.boundaryWidth)"
                  :y="-(field.fieldWidth/2+field.boundaryWidth)"
                  :width="field.fieldLength+field.boundaryWidth*2"
                  :height="field.fieldWidth+field.boundaryWidth*2"
                  :style="{fill: 'green', fillOpacity: 1, stroke: 'none'}"></rect>

            <line v-for="(l,i) of field.lines"
                  :key="'line-' + i"
                  :x1="l.p1.x"
                  :y1="l.p1.y"
                  :x2="l.p2.x"
                  :y2="l.p2.y"
                  :style="[defStyle, l]">
            </line>

            <circle v-for="(a,i) of field.circles"
                    :key="'circle-' + i"
                    :cx="a.center.x"
                    :cy="a.center.y"
                    :r="a.radius"
                    :style="[defStyle, a]">
            </circle>

            <path v-for="(p,i) of field.paths"
                  :key="'path-' + i"
                  :d="pathFromD(p.d)"
                  :style="[defStyle, p]"></path>

            <text v-for="(t,i) of field.texts"
                  :key="'text-' + i"
                  :x="t.p.x"
                  :y="t.p.y"
                  :dx="t.d.x"
                  :dy="t.d.y"
                  :textLength="t.textLength"
                  :lengthAdjust="'spacingAndGlyphs'"
                  :style="[defStyle, t]">
                {{t.text}}
            </text>
        </g>
    </svg>
</template>

<script>


    export default {
        name: "Field",
        props: {
            rotateField: {
                type: Boolean,
                default: false
            },
            defStyle: {
                type: Object,
                default: function () {
                    return {
                        strokeWidth: 10,
                        stroke: 'white',
                        fillOpacity: 0
                    }
                }
            },
            field: {
                type: Object
            }
        },
        computed: {
            getFieldTransformation() {
                if (this.rotateField) {
                    return 'rotate(90) scale(' + (this.field.fieldWidth / this.field.fieldLength) + ')';
                }
                return '';
            },
            viewBox() {
                return (-(this.field.fieldLength / 2 + this.field.boundaryWidth))
                    + ' ' + (-(this.field.fieldWidth / 2 + this.field.boundaryWidth))
                    + ' ' + (this.field.fieldLength + this.field.boundaryWidth * 2)
                    + ' ' + (this.field.fieldWidth + this.field.boundaryWidth * 2);
            },
        },
        methods: {
            pathFromD: function (pd) {
                let d = '';
                for (let s of pd) {
                    d += s.type;
                    for (let a of s.args) {
                        d += ' ' + a
                    }
                }
                return d;
            }
        }
    }
</script>

<style scoped>
    .field-canvas {
        width: 100%;
        height: 100%;
        display: table-row;
    }
</style>