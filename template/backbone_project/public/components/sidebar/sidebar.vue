<script type="text/javascript">
/**
 * 左侧导航
 * author: zhongjiewang
 * date: 2015-12-09
 */
var config = require('app/config.js');
var isMobile = (/mobile/).test(navigator.userAgent.toLowerCase());
    module.exports = {
        props: ['currentView', 'main'],
        data: function() {
            return {
                config: $.map(config, function(item){
                    return {
                        "name": item["name"],
                        "name_CN": item["name_CN"],
                        "param": item["param"],
                        "PCOnly": item["PCOnly"],
                        "view": item['param'] ? (item['name'] + "@" + item['param']) : item['name']
                    }
                }),
                // API 列表数据
                apiData: []
            }
        },
        events: {
            "apiGuide:data": function(data){
                this.apiData = data;
            }
        },
        watch: {
            "apiData": function(nVal, oVal){
                if(this.main != "api_guide_view")return
                if(nVal.length == 0){
                    this.currentView = "api_guide_view"
                }else{
                    this.currentView = "api_guide_view@index:0"
                }

            },
            "main": function(nVal, oVal){
                this.apiData = [];
            }
        },
        methods: {
            viewTo: function(e, view, PCOnly){
                e.preventDefault();
                if(PCOnly && isMobile){
                    this.currentView = "pc_only_view";
                    return
                }
                this.currentView = view;
            }
        }
    }
</script>
<template>
<div class="sidebar" id="sidebar-nav">
    <ul class="nav nav-sidebar">
        <li v-for="c in config"
            :class="{menutree: c.name == 'api_guide_view' ,active: currentView == c.view}">
            <a href="#" v-on:click="viewTo($event, c.view, c.PCOnly)">{{c.name_CN}}</a>
            <ul v-if="c.name == 'api_guide_view' && apiData.length > 0">
                <li 
                    v-for="a in apiData" 
                    :class="{active: currentView == (c.name + '@index:' + $index)}">
                    <a href="#" v-on:click='viewTo($event, c.name + "@index:" +$index)'>{{a.Name}}</a>
                </li>
            </ul>
        </li>
    </ul>
</div>
</template>
<style>
    /* Sidebar navigation */
    /* 超小屏幕（手机，小于 768px） */
    /* 没有任何媒体查询相关的代码，因为这在 Bootstrap 中是默认的（还记得 Bootstrap 是移动设备优先的吗？） */
    /* 小屏幕（平板，大于等于 768px） */
    /*@media (min-width: @screen-sm-min) { ... }*/
    /* 中等屏幕（桌面显示器，大于等于 992px） */
    /*@media (min-width: @screen-md-min) { ... }*/
    /* 大屏幕（大桌面显示器，大于等于 1200px） */
    /*@media (min-width: @screen-lg-min) { ... }*/
    /*@media (min-width: 768px) {
      .sidebar {
        position: fixed;
        top: 51px;
        bottom: 0;
        left: 0;
        z-index: 1000;
        display: block;
        padding: 20px;
        overflow-x: hidden;
        overflow-y: auto;
        background-color: #3b4b58;
        border-right: 1px solid #eee;
      }
    }
    @media (max-width: 767px) {
      .sidebar {
        height: auto!important;
        position: fixed;
        left: -230px;
        top: 51px;
        bottom: 0;
        z-index: 1000;
        padding: 20px;
        overflow-x: hidden;
        overflow-y: auto; 
        background-color: #3b4b58;
        border-right: 1px solid #eee;
        width: 230px;

        -webkit-transition: left .25s ease;
        -moz-transition: left .25s ease;
        -o-transition: left .25s ease;
        -ms-transition: left .25s ease;
        transition: left .25s ease;
      } 
      .sidebar.in {
        position: fixed;
        left: 0;
        display: block;
      }
    }*/
    
    /*.nav-sidebar {
      margin-right: -21px;
      margin-bottom: 20px;
      margin-left: -20px;
      margin-top: 20px;
    }
    .nav-sidebar > hr{
      margin-top: 25px;
      margin-bottom: 15px;
      margin-left: 20px;
      margin-right: 20px;
      border-top:1px solid #d8d8d8;
    }
    */

</style>