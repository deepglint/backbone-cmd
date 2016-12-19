<script type="text/javascript">
/**
 * 导航条
 * author: 
 * date: 
 */
module.exports = {
    props: ['main', 'username'],
    data: function() {
        return {
            version: "default",
            item: '',
            path: 'public/img'
        }
    },
    methods: {
        onClick: function(e) {
            e.preventDefault();
            this.main = 'change_password_view';
        },
        getVersion: function() {
            var self = this;
            $.get("version", function(d, s) {
                self.version = window.version = (d.msg || "default").trim()
            })
        }
    },
    ready: function() {
        this.getVersion()
    }

}
</script>
<template>
<nav class="navbar navbar-fixed-top">
        <div class="navbar-header">
            <img class="logo" height="37" v-bind="{src: path + '/logo_title.png'}" />
        </div>
        <div class=" navbar-collapse" id="navbar-collapse-1">
            <ul class="nav navbar-nav navbar-right">
                <li class="version">版本：{{version}}</li>
                <!--<li class="dropdown">
                    <a href="#" class="dropdown-toggle bkb-user" data-toggle="dropdown" role="button" aria-expanded="true">
                        <span class="user"></span> {{username}} <span class="dropdown-ico"></span>
                    </a>
                    <ul class="dropdown-menu" role="menu">
                        <li><a href="#" v-on:click='onClick($event)'>修改密码</a></li>
                        <li class="divider"></li>
                        <li><a href="logout">退出</a></li>
                    </ul>
                </li>-->
            </ul>
        </div>
</nav>
</template>
<style>
/* navbar navigation */
.navbar {
  border-radius: 0px;
  z-index: 1000;
  background-color: #182126;
  margin-bottom: 0;
}
.navbar .btn-sidebar{
  float: left;
  margin-left: 15px;
}
.navbar .navbar-brand{
  float: left;
}
.navbar-nav > li > a {
    color:#d2d2d2;
}
.navbar-nav > li > a:focus, 
.navbar-nav > li > a:hover{
    color: #ffffff;
    background-color: #182126;
}
.navbar-nav .open > a:focus, 
.navbar-nav .open > a, 
.nav .open > a:hover{
    background-color: transparent;
    color: #fff;
}
.navbar-nav .open .dropdown-menu {
    background-color: rgb(255, 255, 255);
}
.navbar-toggle .icon-bar{
    background-color: #ffffff;
}
.navbar-nav > li > a > span {
    margin: 0 5px;
}
.navbar .nav > li > a {
    text-shadow: 0 0 0;
    white-space: nowrap;
}
.navbar .nav > li > a:focus, 
.navbar .nav > li > a:hover {
	color: #fff;
}
.navbar .nav > li.open > a {
    color: #fff;
}
.navbar .version {
    color: #d2d2d2;
    padding: 15px;
}
.navbar .navbar-collapse {
    float: right;
}
.navbar-header {
    float: left;
    padding-left: 10px;
}
.navbar-right > li {
    float: left;
}
.navbar .navbar-nav {
    margin: 0;
}
.navbar-header .logo {
    margin: 7px 0 6px;
}
/*
@media (min-width: 546px) {
  .navbar .navbar-brand.brand-lg{
    display: block;
    
  }
  .navbar .navbar-brand.brand-sm{
    display: none;
  }
}
@media (max-width: 545px) {
  .navbar .navbar-brand.brand-sm{
    display: block;
    width: 202px;
    margin-left: -15px;
  }
  .navbar .navbar-brand.brand-lg{
    display: none;
  }
}
@media (max-width: 319px) {
  .navbar .navbar-brand.brand-sm{
    display: none;
  }
  .navbar .navbar-brand.brand-lg{
    display: none;
  }
}*/


</style>