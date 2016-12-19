<style>
    /*Main container*/
    body {
        background-color: #f8f8f8;
    }
    .main {
        background-color: #f8f8f8;
        padding-top: 15px;
        padding-bottom: 15px;
    }
    .main-window {
        margin-top: 50px;
    }
    .main .content{
      padding: 20px 20px;
    }
    a:focus, a:hover{
        outline: none;
    }
    a {
        text-decoration: none;
    }
</style>

<script>
var config = require('app/config.js');

// 摘取设置文件中的组件
var components = {};
$.each(config, function(i, item){
    components[item.name] = item.vue;
});

function getViewName(item){
    return item.param ? (item.name + "@" + item.param) : item.name;
}
// main
module.exports = {
    replace: false,
    el: '#MainBody',
    data: {
        username: '',
        main: '',
        param: ''
    },
    computed: {
        currentView: {
            set: function(viewName){
                var i = viewName.indexOf('@');
                var name = i == -1 ? viewName: viewName.substring(0, i);
                var param = i == -1 ? "" : viewName.substring(i+1);
                this.main = name;
                this.param = param;
                bkb.scrollToTop();
            },
            get: function(){
                if(this.param){
                    return this.main + "@" + this.param
                }
                return this.main
            }
        }
    },
    compiled: function () {
        var self = this;
        this.username = window.username;
        this.currentView = getViewName(config[0]);
    },
    ready:function(){
        var self=this
    },
    events: {
        "apiGuide:index": function(index){
            this.$broadcast("apiGuide:index", index);
        },
        "apiGuide:data": function(data){
            this.$broadcast("apiGuide:data", data);
        },
        "apiGuide:clear": function(){
            this.$broadcast("apiGuide:clear");
        }
    },
    components: $.extend({
        // 头部导航
        'navbarview':               require('components/navbar/navbar.vue')
        // 左侧导航
        ,'sidebarview':             require('components/sidebar/sidebar.vue')
        // 修改密码
        ,'change_password_view':    require('components/changePwd/changePwd.vue')
        

    }, components)
    // {
    //     // 头部导航
    //     'navbarview':               require('components/navbar/navbar.vue')
    //     // 左侧导航
    //     ,'sidebarview':             require('components/sidebar/sidebar.vue')
    //     // 修改密码
    //     ,'change_password_view':    require('components/changePwd/changePwd.vue')
    //     // markdown通用模块
    //     ,'markdown_view':           require('components/markdown/markdown.vue')
    //     // api查询
    //     ,'api_guide_view':          require('components/apiGuide/apiGuide.vue')
    //     // 客户端设置
    //     ,'config_view':             require('components/localConfig/localConfig.vue')
    //     // 服务器设置(server)
    //     ,'server_config_view':      require('components/serverConfig/serverConfig.vue')
    //     // 二次开发设置(server)
    //     ,'dev_config_view':         require('components/developerSetting/developerSetting.vue')
    //     // 高级/服务设置(server)
    //     ,'advanced_config_view':    require('components/advancedConfig/advancedConfig.vue')
    //     // 存储设置(server)
    //     ,'storage_config_view':     require('components/storageConfig/storageConfig.vue')
    //     // PC only
    //     ,'pc_only_view':            require('components/PCOnly/PCOnly.vue')
    // }
    
}
</script>

<template>
    <navbarview 
        :main.sync="currentView"
        :username="username"
        >
    </navbarview>
    <aside class="bkb-sidebar">
        <sidebarview
                :current-view.sync="currentView"
                :main.sync="main"
                :param.sync="param">
        </sidebarview>
    </aside>
    <div class="bkb-container">
        <component :is="main"
            :main.sync="main"
            :param.sync="param">
        </component>
    </div>
</template>


