<template>
    <svg id="field-canvas"
         ref="canvas"
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
                  :transform="textTransform(t.p)"
                  :style="[defStyle, defFontStyle, t]">
                {{t.text}}
            </text>
        </g>
    </svg>
</template>

<script>


    export default {
        name: "Field",
        props: {
            defStyle: {
                type: Object,
                default: function () {
                    return {
                        strokeWidth: 10,
                        stroke: 'white',
                        fillOpacity: 1,
                    }
                }
            },
            defFontStyle: {
                type: Object,
                default: function () {
                    return {
                        strokeWidth: 0,
                        textAnchor: "middle",
                        dominantBaseline: "central",
                        font: "bold 7em sans-serif",
                    }
                }
            },
        },
        data() {
            return {
                canvasWidth: 0,
                canvasHeight: 0,
                zoom: 1.0,
            }
        },
        computed: {
            rotateField() {
                return this.fieldRatio < this.viewPortRatio;
            },
            fieldRatio() {
                let wl = this.field.fieldWidth / this.field.fieldLength;
                let lw = this.field.fieldLength / this.field.fieldWidth;
                return (wl + lw) / 2;
            },
            viewPortRatio() {
                return this.canvasHeight / this.canvasWidth;
            },
            getFieldTransformation() {
                let scale = 'scale(' + this.zoom + ')';
                if (this.rotateField) {
                    return 'rotate(90) ' + scale;
                }
                return scale;
            },
            viewBox() {
                if (this.rotateField) {
                    return (-(this.field.fieldWidth / 2 + this.field.boundaryWidth))
                        + ' ' + (-(this.field.fieldLength / 2 + this.field.boundaryWidth))
                        + ' ' + (this.field.fieldWidth + this.field.boundaryWidth * 2)
                        + ' ' + (this.field.fieldLength + this.field.boundaryWidth * 2);
                }
                return (-(this.field.fieldLength / 2 + this.field.boundaryWidth))
                    + ' ' + (-(this.field.fieldWidth / 2 + this.field.boundaryWidth))
                    + ' ' + (this.field.fieldLength + this.field.boundaryWidth * 2)
                    + ' ' + (this.field.fieldWidth + this.field.boundaryWidth * 2);
            },
            field() {
                return this.$store.state.field;
            }
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
            },
            updateCanvasWidth() {
                this.canvasWidth = this.$refs.canvas.scrollWidth;
            },
            updateCanvasHeight() {
                this.canvasHeight = this.$refs.canvas.scrollHeight;
            },
            textTransform(p) {
                if (this.rotateField) {
                    return 'rotate(-90,' + p.x + ',' + p.y + ')'
                }
                return '';
            },
            onScroll(event) {
                let newZoom = this.zoom - event.deltaY / 500;
                if (newZoom < 1) {
                    this.zoom = 1;
                } else {
                    this.zoom = newZoom;
                }
            }
        },
        mounted() {
            this.$nextTick(function () {
                window.addEventListener('resize', this.updateCanvasWidth);
                window.addEventListener('resize', this.updateCanvasHeight);
                document.getElementById("field-canvas").addEventListener("wheel", this.onScroll);

                //Init
                this.updateCanvasWidth();
                this.updateCanvasHeight();
            })

        },
        beforeDestroy() {
            window.removeEventListener('resize', this.updateCanvasWidth);
            window.removeEventListener('resize', this.updateCanvasHeight);
            document.getElementById("field-canvas").removeEventListener("wheel", this.onScroll);
        },
    }
</script>

<style scoped>
    #field-canvas {
        width: 100%;
        height: 100%;
        display: table-row;
    }
</style>