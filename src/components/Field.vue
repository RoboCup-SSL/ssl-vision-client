<template>
    <svg id="field-canvas"
         ref="canvas"
         v-on:mousemove="onMouseMove"
         v-on:mousedown="onMouseDown"
         v-on:mouseup="onMouseUp"
         :viewBox="viewBox">

        <!-- rotate field -->
        <g :transform="getFieldTransformation">

            <!-- draw field background -->
            <rect :x="-(field.fieldLength/2+field.boundaryWidth)"
                  :y="-(field.fieldWidth/2+field.boundaryWidth)"
                  :width="field.fieldLength+field.boundaryWidth*2"
                  :height="field.fieldWidth+field.boundaryWidth*2"
                  ref="background"
                  :style="{fill: 'green', fillOpacity: 1, stroke: 'none'}"></rect>

            <template v-for="(s,i) of field.shapes">

                <line v-if="s.line"
                      :key="'shape-' + i"
                      :x1="s.line.p1.x"
                      :y1="s.line.p1.y"
                      :x2="s.line.p2.x"
                      :y2="s.line.p2.y"
                      :style="[defStyle, s.line]">
                </line>

                <circle v-if="s.circle"
                        :key="'shape-' + i"
                        :cx="s.circle.center.x"
                        :cy="s.circle.center.y"
                        :r="s.circle.radius"
                        :style="[defStyle, s.circle]">
                </circle>

                <path v-if="s.path"
                      :key="'shape-' + i"
                      :d="pathFromD(s.path.d)"
                      :style="[defStyle, s.path]"></path>

                <text v-if="s.text"
                      :key="'shape-' + i"
                      :x="s.text.p.x"
                      :y="s.text.p.y"
                      :transform="textTransform(s.text.p)"
                      :style="[defStyle, defFontStyle, s.text]">
                    {{s.text.text}}
                </text>

            </template>
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
                canvasWidth: 1,
                canvasHeight: 1,
                zoom: 1.0,
                translation: {x: 0, y: 0},
                activeTranslation: {x: 0, y: 0},
                mouseDownPoint: null,
            }
        },
        computed: {
            fieldToPixelRatioX() {
                if (this.rotateField) {
                    return (this.field.fieldWidth + this.field.boundaryWidth * 2) / this.canvasWidth;
                }
                return (this.field.fieldLength + this.field.boundaryWidth * 2) / this.canvasWidth;
            },
            fieldToPixelRatioY() {
                if (this.rotateField) {
                    return (this.field.fieldLength + this.field.boundaryWidth * 2) / this.canvasHeight;
                }
                return (this.field.fieldWidth + this.field.boundaryWidth * 2) / this.canvasHeight;
            },
            fieldTranslationX() {
                return (this.translation.x + this.activeTranslation.x) * this.fieldToPixelRatioX;
            },
            fieldTranslationY() {
                return (this.translation.y + this.activeTranslation.y) * this.fieldToPixelRatioY;
            },
            rotateField() {
                return this.meanFieldRatio < this.canvasRatio;
            },
            meanFieldRatio() {
                let wl = this.field.fieldWidth / this.field.fieldLength;
                let lw = this.field.fieldLength / this.field.fieldWidth;
                return (wl + lw) / 2;
            },
            canvasRatio() {
                return this.canvasHeight / this.canvasWidth;
            },
            getFieldTransformation() {
                let scale = 'scale(' + this.zoom + ') ';
                let transform = 'translate(' + this.fieldTranslationX + ',' + this.fieldTranslationY + ') ';
                if (this.rotateField) {
                    return transform + scale + 'rotate(90)';
                }
                return transform + scale;
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
                // Firefox does not support clientWidth for SVG, so fall back to the parent element
                this.canvasWidth = this.$refs.canvas.clientWidth || this.$refs.canvas.parentNode.clientWidth;
            },
            updateCanvasHeight() {
                // Firefox does not support clientWidth for SVG, so fall back to the parent element
                this.canvasHeight = this.$refs.canvas.clientHeight || this.$refs.canvas.parentNode.clientHeight;
            },
            textTransform(p) {
                if (this.rotateField) {
                    return 'rotate(-90,' + p.x + ',' + p.y + ')'
                }
                return '';
            },
            onScroll(event) {
                let newZoom = this.zoom - event.deltaY / 300;
                if (newZoom < 1) {
                    this.zoom = 1;
                } else {
                    this.zoom = newZoom;
                }
            },
            onMouseMove(event) {
                if (this.mouseDownPoint !== null) {
                    this.activeTranslation = {
                        x: event.clientX - this.mouseDownPoint.x,
                        y: event.clientY - this.mouseDownPoint.y
                    };
                }
            },
            onMouseDown(event) {
                this.mouseDownPoint = {x: event.clientX, y: event.clientY};
            },
            onMouseUp() {
                if (this.mouseDownPoint !== null) {
                    this.translation = {
                        x: this.translation.x + this.activeTranslation.x,
                        y: this.translation.y + this.activeTranslation.y,
                    };
                    this.activeTranslation = {x: 0, y: 0};
                    this.mouseDownPoint = null;
                }
            },
            onClick(event) {
                event = event || window.event;
                if (event.key === " ") {
                    this.zoom = 1;
                    this.translation = {x: 0, y: 0};
                }
            }
        },
        mounted() {
            this.$nextTick(function () {
                window.addEventListener('resize', this.updateCanvasWidth);
                window.addEventListener('resize', this.updateCanvasHeight);
                document.getElementById("field-canvas").addEventListener("wheel", this.onScroll);
                document.addEventListener('keydown', this.onClick);

                //Init
                this.updateCanvasWidth();
                this.updateCanvasHeight();
            })

        },
        beforeDestroy() {
            window.removeEventListener('resize', this.updateCanvasWidth);
            window.removeEventListener('resize', this.updateCanvasHeight);
            document.getElementById("field-canvas").removeEventListener("wheel", this.onScroll);
            document.removeEventListener('keydown', this.onClick);
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