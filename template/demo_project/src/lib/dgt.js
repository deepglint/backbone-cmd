var dgt = window.dgt || {};

dgt.ajax = function(options) {
    var def = $.Deferred();

    var suc = options.success;
    var err = options.error;

    var successCallback = function(data, textStatus, xhr) {
        try {
            // if($.type(data) == 'object'){
            //     var tmp = {};
            //     $.each(data, function(key, val){
            //         tmp[key.toLowerCase()] = val;
            //     })
            //     data = tmp;
            // }
            // 正常格式，包含code
            if (data.code !== undefined && data.code !== 0) {
                // 拦截错误码，错误码处理
                err && err.call(this, xhr, textStatus, data);
                def.reject.call(this, xhr, textStatus, data);
            } else {

                suc && suc.call(this, data, textStatus, xhr);
                def.resolve.call(this, data, textStatus, xhr);
            }
        } catch (e) {
            err && err.call(this, xhr, e.message, data);
            def.reject.call(this, xhr, e.message, data);
        }
    };

    var opts = $.extend(true, 
                            {
                                cache: false, // mock数据时，关闭
                                dataType: "json",
                                timeout: 30000
                            }, 
                            options, 
                            {
                                success: successCallback,
                                error: function(xhr, textStatus, error) {
                                    err && err.call(this, xhr, textStatus, error);
                                    def.reject.call(this, xhr, textStatus, error);
                                }
                            });

    var ajax = $.ajax(opts);


    return def;
};

dgt.tip = function(msg, type) {
    msg = msg || "";
    type = type || "default";
    var main = $('<div class="dgt-tip">').appendTo('body');
    var content = $('<div class="dgt-tip-content">').appendTo(main);
    var html = msg;
    content.html(msg);
    main.addClass('fade-in-down');
    setTimeout(function() {
        main.one('animationEnd webkitAnimationEnd', function() {
            main.remove();
        }).addClass('fade-out-down');
    }, 2000);
    if (type == 'error') {
        console.error(msg)
    } else {
        console.info(msg)
    }
};

dgt.uniqueId = function() {
    var idCounter = 0;
    return function(prefix) {
        var id = new Date().getTime().toString(16) + "_" + ++idCounter + "";
        return prefix ? prefix + id : id;
    }
}();

function capitalizeFirstLetter(key){
    return key.replace(/^(\w{1})/, function(a){
        return a.toUpperCase()
    })
}

function arrayFormat(body){
    return $.map(body, function(item){
        switch($.type(item)){
            case "array":
                return arrayFormat(item);
                break;
            case "object":
                return objectFormat(item);
                break;
            default:
                return item
        }
    })
}

function objectFormat(body){
    var tmp = {};
    $.each(body, function(key, val){
        key = capitalizeFirstLetter(key);
        switch($.type(val)){
            case "array":
                tmp[key] = arrayFormat(val);
                break;
            case "object":
                tmp[key] = objectFormat(val);
                break;
            default:
                tmp[key] = val
        }
    })
    return tmp;
}


dgt.formatBody = function(body){
    var tmp;
    switch($.type(body)){
        case "object":
            return JSON.stringify( objectFormat(body) );
            break;
        case "array":
            return JSON.stringify( arrayFormat(body) );
            break;
        default:
            return body
    }
}

module.exports = dgt;