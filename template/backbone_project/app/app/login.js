window.bkb = require('./bkb.js');
var Vue = require('vue');
var login = require('components/login/login.vue');
var VueValidator = require('vue-validator')
Vue.use(VueValidator)
var app = new Vue(login);
app.$mount('#MainBody');
