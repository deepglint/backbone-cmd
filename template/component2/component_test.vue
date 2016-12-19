<script>
import test from "./{{component}}.vue"
export default {
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
    <test :text_field="test_{{component}}"></test>
</div>
</template>

<style lang="sass">
    .test-container {
        width: 200px;
        border: solid black 2px;
    }
</style>