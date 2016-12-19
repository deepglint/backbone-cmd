var config = [
    // {name: "bgmodel_config_view", name_CN: "背景", vue: require('components/bgmodelset/bgmodelset.vue')},
    /**
     * markdow 解析
     * 参数作为加载路径使用
     */
    {
        "vue": require('components/markdown/markdown.vue'),
        "name": "markdown_view",
        "name_CN": "如何开始",
        "PCOnly": false,
        "param": "public/md/getstarted.md"
    },
    /**
     * markdow 解析
     * 参数作为加载路径使用
     */
    {
        "vue": require('components/markdown/markdown.vue'),
        "name": "markdown_view",
        "name_CN": "数据结构定义",
        "PCOnly": false,
        "param": "public/md/data_structure.md"
    }

]

module.exports = config;