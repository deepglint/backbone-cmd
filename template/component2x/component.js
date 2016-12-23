window.dgt = require("../../utils/dgt").default
window.config = {fakeRequest:true}
var Vue = require('vue')
var Vuex = require('vuex')
Vue.use(Vuex);
var app=require("./{{component}}_test.vue")
var vm=new Vue(app)
vm.$mount('#show');
console.log(vm)
