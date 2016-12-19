#NetWorkConfig Spec
##综述

NetworkConfig分为3个功能模块

1.设备网络设置

Name:dev-net
{
.title
.ip-form
.button
.tip
}

2.网管服务器设置

Name:.serv-net
{
.title
.ip-form
.button
.tip
}

3.时间服务器设置

Name:.time-serv
{
.title
.input-form
.button
.tip
}
4.视觉设置

Name:.algorithm
{
.title
.introduction
.input-checkbox
.button
.tip
}

基础组件:

Name:.ip-form

外部是div,内部用％来规定宽度

组件的Hooks:

tips变量

onerror函数

function(err){}

err is string


ref:[http://vuejs.org/guide/components.html#Props]
