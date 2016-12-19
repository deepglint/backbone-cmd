<script type="text/javascript">
/**
 * markdown阅读器
 * author: zhongjiewang
 * date: 2015-12-09
 */

// markdown parse
var marked = require('lib/marked.js');
function markdown(md){
    return marked(md, {
        gfm: true,
        tables: true,
        breaks: false,
        pedantic: false,
        sanitize: false,
        smartLists: true,
        smartypants: false
    })
};
module.exports = {
    props: ['param'],
    data: function() {
        return {
            detail: ''
        }
    },
    ready: function() {
        this.parse();
    },
    watch: {
        param: function(){
            this.parse();
        }
    },
    methods: {
        parse: function() {
            var self = this;
            var url = this.param;
            if(url != ""){
            $.ajax({
                type: 'get',
                url: url,
                cache: true,
                success: function(msg) {
                    self.detail = markdown(msg);
                }
            });
            }
        }
    }
}
</script>
<template>
<div class="markdown">{{{detail}}}</div>
</template>
<style type="text/css"></style>