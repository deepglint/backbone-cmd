<script>
import test from "./{{component}}.vue"
import Vuex from "vuex"
import storeConfig from "../../store"
const store = new Vuex.Store(storeConfig)

export default {
    store,
    data(){
        return {

        }
    },
    components: {
        test
    }
}
</script>

<template>
<div class="test-container">
    <test></test>
</div>
</template>

<style lang="sass">
    .test-container {
        width: 200px;
        border: solid black 2px;
    }
</style>