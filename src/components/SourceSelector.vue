<template>
    <div id="source-selector">
        Sources:
        <template v-for="(sourceName, sourceId) in sources">
            <input
                v-bind:key="'input-' + sourceId"
                type="radio"
                :id="sourceId"
                name="source"
                :value="sourceId"
                :checked="sourceId === activeSourceId"
                @click="updateSource(sourceId)"
            />
            <label v-bind:key="'label-' + sourceId" :for="sourceId">{{ sourceName }}</label>
        </template>
    </div>
</template>

<script>
export default {
    name: "SourceSelector",
    computed: {
        sources() {
            return this.field.sources;
        },
        activeSourceId() {
            return this.field.activeSourceId;
        },
        field() {
            return this.$store.state.field;
        }
    },
    methods: {
        updateSource(newSource) {
            this.$socket.sendObj({activeSourceId: newSource});

            if (history.pushState) {
                const newUrl = window.location.protocol + "//" + window.location.host + window.location.pathname + '?sourceId=' + newSource;
                window.history.pushState({path:newUrl},'',newUrl);
            }
        }
    }
}
</script>

<style scoped>
#source-selector {
    align-content: center;
    width: 100%;
    bottom: 0.1em;
    position: absolute;
}

</style>
