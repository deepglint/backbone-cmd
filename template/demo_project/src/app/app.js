

Vue.transition('staggered', {

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
});