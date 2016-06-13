
var app=require("./login_test.vue")
var Vue = require('vue')
var vm=new Vue(app)
var VueValidator = require('vue-validator')
Vue.use(VueValidator)
console.log(vm)