window.bkb = require('./bkb.js');
var Vue = window.Vue = require('vue');

Vue.config.debug = true

var VueValidator = require('vue-validator');
Vue.use(VueValidator);
var app = require("./app.vue");
var d = require("vue-datetime-picker/src/vue-datetime-picker.js");
var vm = new Vue(app);

var staggered = {

    beforeEnter: function(el) {
        var h = $(el).height();
        var p = $(el).parent().height();
        $(el).parent().css('height', Math.max(h, p))
    },
    enter: function(el) {
    },
    afterEnter: function(el) {
        $(el).parent().css('height', 'auto');
    },
    enterCancelled: function(el) {
        // handle cancellation
    },

    beforeLeave: function(el) {
        var h = $(el).height();
        var p = $(el).parent().height();
        $(el).parent().css('height', Math.max(h, p))
    },
    leave: function(el) {
    },
    afterLeave: function(el) {
        $(el).parent().css('height', 'auto');
    },
    leaveCancelled: function(el) {
        // handle cancellation
    }
};

Vue.transition('staggered', staggered)
Vue.transition('staggered2', staggered)

// 设备状态
Vue.filter("deviceStatus", function(value){
    switch(value){
        case -1:
        case false:
            return '<span class="bkb-device-status err"></span>';
            break;
        case 0:
        case true:
            return '<span class="bkb-device-status ok"></span>';
            break;
        case 1:
            return '<span class="bkb-device-status warn"></span>';
            break;
    }
})


Vue.filter('unixTimestampSecondToMoment', function(value){
    return value*=1000;
})
Vue.filter('momentToString', function(value, format){
    return moment(value).format(format);
})