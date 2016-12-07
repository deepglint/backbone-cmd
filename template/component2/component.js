var app=require("./{{component}}_test.vue")
var Vue = require('vue')
var vm=new Vue(app)
vm.$mount('#show');
console.log(vm)
