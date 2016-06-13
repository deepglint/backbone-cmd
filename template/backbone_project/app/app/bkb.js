var bkb = window.bkb || {};

var message = {
    loading: "<div class='loading-wait col-sm-offset-6 col-md-offset-6'></div>",
    error: "<div class=\"alert alert-danger\" role=\"alert\"><strong>Sensor 网络连接错误! </strong>请检查网络连接或刷新重试.</div>",
    login: "<div class=\"alert alert-danger\" role=\"alert\"><strong> 登录超时! </strong> 请重新登录。</div>"
};

var Confirm = require('./Confirm.js');

window.sensorid = "";

var urlReg = /\/backend\/bumble\/api\/(get|set|del)\/([\w\/]+)?([\?]*.*)/i;

// /backend/bumble/api/(get or post)/version->
// /api/maintain/bumble/version
// for jquery

var ajaxArray = [];


bkb.ajax = function(options) {
    var def = $.Deferred();


    // console.log(options.url, window.sensorid, 'sensorid')
    if (window.sensorid) {

        var url = options.url;

        var m = url.match(urlReg);

        // ["/backend/bumble/api/get/version/sfdf?", "get", "version"]

        if (m && m.length) {
            url = "/api/maintain/bumble/" + (m[2] || "") + (m[3] || "")
        }

        url += url.indexOf('?') == -1 ? "?" : "&";
        url += 'sensorid=' + window.sensorid;

        options.url = url;

    }

    var suc = options.success;
    var err = options.error;

    var successCallback = function(data, textStatus, xhr) {
        if (data == undefined) {
            data = "success";
        }
        // console.log(data, typeof data, 'teststststst')
        try {
            // 如果是正常的格式
            if (data.code == 0 && data.msg) {
                // msg如果包含错误码
                var arr = data.msg.match(/^(\d+)\: (.+)?/i);
                if (arr && arr.length > 0) {
                    data.code = arr[1];
                }
            }
            // 正常格式，包含code
            if (data.code !== undefined && data.code !== 0) {
                // 拦截错误码，错误码处理
                if (data.code == 401) {
                    console.log('401', options.url, 'redirect login')
                    location.href = "login"
                }
                err && err.call(this, xhr, textStatus, data);
                def.reject.call(this, xhr, textStatus, data);
            } else {

                // 如果格式不匹配，转换格式
                if (data.code === undefined) {
                    if (typeof data == "object") {
                        data = JSON.stringify(data);
                    }
                    data = {
                        code: 0,
                        msg: data,
                        data: data,
                        redirect: ""
                    }
                }

                suc && suc.call(this, data, textStatus, xhr);
                def.resolve.call(this, data, textStatus, xhr);
            }
        } catch (e) {
            err && err.call(this, xhr, textStatus, data);
            def.reject.call(this, xhr, textStatus, data);
        }
    };
    
    var id = bkb.uniqueId('ajax');

    var opts = $.extend(true, {
        cache: false,
        timeout: 30000
        // dataType: "json"
    }, options, {
        success: successCallback,
        error: function(xhr, textStatus, error) {
            if(textStatus == 'abort')return;
            err && err.call(this, xhr, textStatus, error);
            def.reject.call(this, xhr, textStatus, error);
        }
    });

    var ajax = $.ajax(opts).always(function () {
        deleteAjaxObjectById(id);
    });
    
    ajaxArray.push({
        id: id,
        ajax: ajax
    });

    // $.extend(def, ajax);

    return def;
};

function deleteAjaxObjectById(id) {
    for(var i = 0, j = ajaxArray.length; i<j;i++){
        if(id == ajaxArray[i].id){
            ajaxArray.splice(i, 1);
            break;
        }
    }
}

window.ajaxArray = ajaxArray;

bkb.ajaxClear = function () {
    if(ajaxArray.length > 0 ){
        ajaxArray.forEach(function (item, index) {
            item.ajax && item.ajax.abort && item.ajax.abort();
        })
    }
    return ajaxArray;
}

bkb.tip = function(msg, type) {
    msg = msg || "";
    type = type || "default";
    var main = $('<div class="bkb-tip">').appendTo('body');
    var content = $('<div class="bkb-tip-content">').appendTo(main);
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

bkb.uniqueId = function() {
    var idCounter = 0;
    return function(prefix) {
        var id = new Date().getTime().toString(16) + "_" + ++idCounter + "";
        return prefix ? prefix + id : id;
    }
}();

bkb.confirm = function(msg, options) {
    return new Confirm(msg, options)
}

bkb.scrollToTop = function() {
    document.body.scrollTop = 0;
    document.documentElement.scrollTop = 0;// for IE
}

module.exports = bkb;