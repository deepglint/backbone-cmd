var config = [
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
    },
    /**
     * apiGuide api 查看
     */
    // {
    //     "vue": require('components/apiGuide/apiGuide.vue'),
    //     "name": "api_guide_view",
    //     "name_CN": "API 使用手册",
    //     "PCOnly": false,
    //     "param": ""
    // },
    /**
     * serverConfig 服务器设置
     */
    {
        "vue": require('components/serverConfig/serverConfig.vue'),
        "name": "server_config_view",
        "name_CN": "前端设备设置",
        "PCOnly": true,
        "param": ""
    },
    /**
     * 高级设置
     */
    {
        "vue": require('components/serverPortal/serverPortal.vue'),
        "name": "advanced_config_view",
        "name_CN": "服务器配置",
        "PCOnly": true,
        "param": ""
    },
    /**
     * 二次开发设置
     */
    {
        "vue": require('components/developerSetting/developerSetting.vue'),
        "name": "dev_config_view",
        "name_CN": "二次开发配置",
        "PCOnly": true,
        "param": ""
    }
]

module.exports = config;