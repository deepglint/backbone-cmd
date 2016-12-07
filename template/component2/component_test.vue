<script>
import test from "./{{component}}.vue"

module.exports = {
    data: function(){
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