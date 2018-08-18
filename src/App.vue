<template>
    <div id="app">
        <Field/>
    </div>
</template>

<script>
    import Field from "./components/Field.vue";

    export default {
        name: "app",
        components: {
            Field
        },
        computed: {
            field() {
                return this.$store.state.field;
            }
        }
    };
</script>

<style>
    #app {
        font-family: "Avenir", Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
        height: 100vh;
    }
</style>
