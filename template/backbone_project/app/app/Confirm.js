(function (global, factory) {
    typeof exports === 'object' && typeof module !== 'undefined' ? module.exports = factory() :
    typeof define === 'function' && define.amd ? define(factory) :
    global.moment = factory()
}(this, function () { 
    'use strict';

    /**
     * options.confirm      确认后的回调函数
     * options.cancel       取消后的回调函数
     * options.render       渲染后的回调函数
     * options.confirmText  确认按钮的文字
     * options.cancelText   取消按钮的文字
     */
    

    /*
        new Confirm('save confirm', {
            confirm: function(){
                console.log('confirm config')
            },
            cancel: function(){
                console.log('cancle confirm')
            },
            render: function(){
                console.log('render confirm')
            }
        })
    */

    function Confirm(msg, options){
        this.msg = msg || "确认";
        this.options = $.extend(true, {
            container: "body",
            confirmText: "确认",
            cancelText: "取消"
        }, options || {});
        this.initialize();
    }

    Confirm.prototype = {
        initialize: function(){
            this.render();
            this.bindEvent()
        },
        render: function(){
            var msg = '<div class="page-footer-bar">' +
                '<label>'+this.msg+'</label>' +
                '<button type="button" class="btn btn-default" data-role="confirm">'+this.options.confirmText+'</button>' +
                '<button type="button" class="btn btn-default" data-role="cancel">'+this.options.cancelText+'</button>' +
            '</div>';

            this.content = $(msg).appendTo(this.options.container);

            var renderCallback = this.options.render;

            if($.type(renderCallback) == 'function'){
                renderCallback();
            }

        },
        bindEvent: function(){
            var self = this;
            this.content.on('click', '[data-role]', function(){
                var name = $(this).attr('data-role');
                self.destroy();
                self[name]();
            })
        },
        confirm: function(){
            var callback = this.options.confirm;

            if($.type(callback) == 'function'){
                callback();
            }
        },
        cancel: function(){
            var callback = this.options.cancel;

            if($.type(callback) == 'function'){
                callback();
            }
        },
        destroy: function(){
            this.content.off('click').remove();
        }
    }


    return Confirm


}));