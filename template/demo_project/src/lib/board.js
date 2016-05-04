/******/ (function(modules) { // webpackBootstrap
/******/    // The module cache
/******/    var installedModules = {};

/******/    // The require function
/******/    function __webpack_require__(moduleId) {

/******/        // Check if module is in cache
/******/        if(installedModules[moduleId])
/******/            return installedModules[moduleId].exports;

/******/        // Create a new module (and put it into the cache)
/******/        var module = installedModules[moduleId] = {
/******/            exports: {},
/******/            id: moduleId,
/******/            loaded: false
/******/        };

/******/        // Execute the module function
/******/        modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);

/******/        // Flag the module as loaded
/******/        module.loaded = true;

/******/        // Return the exports of the module
/******/        return module.exports;
/******/    }


/******/    // expose the modules object (__webpack_modules__)
/******/    __webpack_require__.m = modules;

/******/    // expose the module cache
/******/    __webpack_require__.c = installedModules;

/******/    // __webpack_public_path__
/******/    __webpack_require__.p = "";

/******/    // Load entry module and return exports
/******/    return __webpack_require__(0);
/******/ })
/************************************************************************/
/******/ ([
/* 0 */
/***/ function(module, exports, __webpack_require__) {

    // var Raphael = require('lib/raphael.js');

    var klass = __webpack_require__(2);

    var acc = window.acc = __webpack_require__(3);
    var util = __webpack_require__(5);

    var Board = klass({
        initialize: function(options){
            var container = this.container = options.container;
            var width = this.width = options.width;
            var height = this.height = options.height;

            this.initProps();

            var main = this.main = $('<div class="dgt-board">').appendTo(container);

            var layers = this.layers = Raphael(main.get(0), width, height);
            layers.setViewBox(0, 0, width, height);

            this.bindEvents()

            this.zoomTo(options.zoom || 1);

        },
        initProps: function(){
            this.pluginName = "";
            this.plugin = null;

            this.list = [];

            this.foreground = "#000";
            this.background = "#000";

            this.minZoom = 0.1;
            this.maxZoom = 5;

            this.zoom = 1;
        },
        set: function(name, value){
            this[name] = value;
            eve('canvas.' + name, this, value);
        },
        get: function(name){
            return name.substr(0, 1) != "_" ? this[name]: "";
        },
        switchTo: function(name, options){
            var pln;
            if(pln = Board.plugins[name]){
                this._switchTo(name, pln, options);
                return pln
            }else{
                console.warn(name + " plugin 不存在");
                return false;
            }
        },
        // 注入插件
        _switchTo: function(name, pln, options){
            this.removePlugin();
            this.pluginName = name;
            this.plugin = new pln(this, options);
        },
        removePlugin: function(){
            if(this.pluginName !== ""){
                this.plugin.destroy();
            }
        },
        add: function(item){
            this.list.push(item);
            $(this).trigger('board:add', item);
        },
        remove: function(id){
            var list = this.list;
            var index = 0;
            for(var i = 0, j = list.length; i<j;i++){
                if(list[i].uid == id){
                    index = i;
                    break
                }
            }
            var aa = this.list.splice(index, 1);
            aa[0].el.remove();
        },
        clearAll: function(){
            this.list.forEach(function(item){
                item.el.remove();
            })
            this.list = [];
        },
        data: function(){
            var plugin, newItem;
            var data = $.map(this.list, function(item){
                plugin = Board.plugins[item.datatype];
                newItem = {
                    datatype: item.datatype
                };
                
                if(plugin && plugin.getData){
                    newItem.data = plugin.getData(item);
                }
                return newItem;
            });
            return JSON.stringify(data);
        },
        readData: function(data){
            var self = this;
            
            $.each(data, function(i, item){
                self._createShape(item);
            });
        },
        _createShape: function(item){
            var el, plugin, data;
            var svg = this.layers;
            var zoom = this.zoom;

            plugin = Board.plugins[item.datatype];
            if(plugin && plugin.draw){
                el = plugin.draw(item, svg, zoom);
                data = $.extend(true, {uid: util.uniqueId(item.datatype+"_")}, item, {el: el, datatype: item.datatype});
                this.add(data);
            }
        },
        zoomTo: function(zoom){
            if (!zoom || zoom < this.minZoom || zoom > this.maxZoom) {
                return false;
            }
            var width = Math.floor(this.width * zoom);
            var height = Math.floor(this.height * zoom);
            this.layers.setSize(width, height);
            this.zoom = zoom;

            eve("board.zoom", this, zoom);
        },
        // 修改边线宽度
        fixSize: function(){
            var plugin, newItem;
            var zoom = this.zoom;
            $.each(this.list, function(i, item){
                plugin = Board.plugins[item.datatype];
                newItem = {
                    datatype: item.datatype
                };
                if(plugin && plugin.fixSize){
                    plugin.fixSize(item, zoom);
                }
            });
        },
        update: function(width, height){
            this.width = width;
            this.height = height;
            this.layers.setSize(width, height).setViewBox(0,0,width, height);
        },
        // 事件绑定
        bindEvents: function(){
            $(this.main).on('contextmenu selectstart', function(e) {
                e.preventDefault();
            });
            eve.on('board.zoom', this.fixSize)
        }
    })
    .statics('plugins', {})
    .statics('addPulgin', function(name, plugin){
        this['plugins'][name] = plugin;
    });

    // 矩形
    Board.addPulgin('rect', __webpack_require__(4));

    window.Board = Board;

/***/ },
/* 1 */,
/* 2 */
/***/ function(module, exports, __webpack_require__) {

    var __WEBPACK_AMD_DEFINE_FACTORY__, __WEBPACK_AMD_DEFINE_RESULT__;/*!
      * klass: a classical JS OOP façade
      * https://github.com/ded/klass
      * License MIT (c) Dustin Diaz 2014
      */

    !function (name, context, definition) {
      if (true) !(__WEBPACK_AMD_DEFINE_FACTORY__ = (definition), __WEBPACK_AMD_DEFINE_RESULT__ = (typeof __WEBPACK_AMD_DEFINE_FACTORY__ === 'function' ? (__WEBPACK_AMD_DEFINE_FACTORY__.call(exports, __webpack_require__, exports, module)) : __WEBPACK_AMD_DEFINE_FACTORY__), __WEBPACK_AMD_DEFINE_RESULT__ !== undefined && (module.exports = __WEBPACK_AMD_DEFINE_RESULT__))
      else if (typeof module != 'undefined') module.exports = definition()
      else context[name] = definition()
    }('klass', this, function () {
      var context = this
        , f = 'function'
        , fnTest = /xyz/.test(function () {xyz}) ? /\bsupr\b/ : /.*/
        , proto = 'prototype'

      function klass(o) {
        return extend.call(isFn(o) ? o : function () {}, o, 1)
      }

      function isFn(o) {
        return typeof o === f
      }

      function wrap(k, fn, supr) {
        return function () {
          var tmp = this.supr
          this.supr = supr[proto][k]
          var undef = {}.fabricatedUndefined
          var ret = undef
          try {
            ret = fn.apply(this, arguments)
          } finally {
            this.supr = tmp
          }
          return ret
        }
      }

      function process(what, o, supr) {
        for (var k in o) {
          if (o.hasOwnProperty(k)) {
            what[k] = isFn(o[k])
              && isFn(supr[proto][k])
              && fnTest.test(o[k])
              ? wrap(k, o[k], supr) : o[k]
          }
        }
      }

      function extend(o, fromSub) {
        // must redefine noop each time so it doesn't inherit from previous arbitrary classes
        function noop() {}
        noop[proto] = this[proto]
        var supr = this
          , prototype = new noop()
          , isFunction = isFn(o)
          , _constructor = isFunction ? o : this
          , _methods = isFunction ? {} : o
        function fn() {
          if (this.initialize) this.initialize.apply(this, arguments)
          else {
            fromSub || isFunction && supr.apply(this, arguments)
            _constructor.apply(this, arguments)
          }
        }

        fn.methods = function (o) {
          process(prototype, o, supr)
          fn[proto] = prototype
          return this
        }

        fn.methods.call(fn, _methods).prototype.constructor = fn

        fn.extend = arguments.callee
        fn[proto].implement = fn.statics = function (o, optFn) {
          o = typeof o == 'string' ? (function () {
            var obj = {}
            obj[o] = optFn
            return obj
          }()) : o
          process(this, o, supr)
          return this
        }

        return fn
      }

      return klass
    });

/***/ },
/* 3 */
/***/ function(module, exports) {

    function decimalLength(num){
        var str = num.toString();
        var index = str.indexOf('.');
        return index == -1 ? 0 : str.substr(index+1).length;
    }

    function suffixInteger(num, length){
        var str = num.toString();
        var decimalLen = decimalLength(num);
        str += Math.pow(10, length - decimalLen).toString().substr(1);
        return Number( str.replace('.', '') );
    }

    // 加法
    function accAdd(num1, num2){
        var r1 = decimalLength(num1);
        var r2 = decimalLength(num2);

        var max = Math.max(r1, r2);

        var n1 = suffixInteger(num1, max);
        var n2 = suffixInteger(num2, max);

        return Number( ( (n1 + n2)/Math.pow(10, max) ).toFixed(max) );
    }

    // 减法
    function accSubtr(num1, num2){
        var r1 = decimalLength(num1);
        var r2 = decimalLength(num2);

        var max = Math.max(r1, r2);

        var n1 = suffixInteger(num1, max);
        var n2 = suffixInteger(num2, max);

        return Number( ( (n1 - n2)/Math.pow(10, max) ).toFixed(max) );
    }
    // 乘法
    function accMul(num1, num2){
        var r1 = decimalLength(num1);
        var r2 = decimalLength(num2);

        var max = Math.max(r1, r2);

        var n1 = suffixInteger(num1, max);
        var n2 = suffixInteger(num2, max);

        return n1*n2/Math.pow(10, max*2);

    }
    // 除法
    function accDiv(num1, num2){
        var r1 = decimalLength(num1);
        var r2 = decimalLength(num2);

        var max = Math.max(r1, r2);

        var n1 = suffixInteger(num1, max);
        var n2 = suffixInteger(num2, max);

        return n1/n2;
    }

    // var i = 0;
    // console.time('add')
    // while(i++ < 10000){
    //     accAdd(0.2222, 0.8);
    // }
    // console.timeEnd('add')

    // console.time('subtr')
    // while(i++ < 20000){
    //     accSubtr(0.2222, 0.8);
    // }
    // console.timeEnd('subtr')

    // console.time('mul')
    // while(i++ < 30000){
    //     accMul(0.2222, 0.8);
    // }
    // console.timeEnd('mul')

    // console.time('div')
    // while(i++ < 40000){
    //     accDiv(0.2222, 0.8);
    // }
    // console.timeEnd('div')

    module.exports = {
        div: accDiv,
        mul: accMul,
        add: accAdd,
        subtr: accSubtr
    }

/***/ },
/* 4 */
/***/ function(module, exports, __webpack_require__) {

    var Parent = __webpack_require__(6);

    var util = __webpack_require__(5);

    var pathWidth = 1;
    var fontSize = 14;

    var Rect = Parent.extend({
        type: "rect",

        bindEvent: function(){
            var self = this;
            $(this.canvas.main)
                .on('mousedown.Rect', $.proxy(this.mousedown, this))
                .on('mousemove.Rect', $.proxy(this.mousemove, this))
                .on('mouseup.Rect', $.proxy(this.mouseup, this))

            eve.on('board.zoom', function(zoom){
                self.fixSize();
            })
        },
        initProp: function(){
            this.isDrawing = false;
            this.isMoving = false;

            this.startX = 0;
            this.startY = 0;

            this.x = 0;
            this.y = 0;
        },
        mousedown: function(e){
            e.preventDefault();
            if (e.button != 0 || this.isMoving) return false;
            var canDraw = $(this.canvas).triggerHandler("board.drawstart", this);
            if(canDraw === false)return false;
            var offset = this.getRealOffset(e);
            this.x = this.getZoomSize(offset.offsetX);
            this.y = this.getZoomSize(offset.offsetY);

            this.startX = e.pageX;
            this.startY = e.pageY;

            this.isDrawing = true;
            this.isMoving = false;
        },
        mousemove: function(e){
            if (this.isDrawing) {
                var x = this.x;
                var y = this.y;
                if(!this.isMoving){
                    this.isMoving = true;
                    this.drawPoint(x, y);
                }
                var offset = this.getRealOffset(e);
                var x2 = this.getZoomSize(offset.offsetX);
                var y2 = this.getZoomSize(offset.offsetY);
                this.setAttr(x, y, x2, y2)
            }
        },
        mouseup: function(e){
            if(!this.isMoving){
                return false;
            }
            this.isDrawing = false;
            this.isMoving = false;
            this.drawRect();
        },
        mouseleave: function(e){
            if(!this.isMoving){
                return false;
            }
            this.isDrawing = false;
            this.isMoving = false;
            this.drawRect();
        },
        drawPoint: function(x, y) {
            var canvas = this.canvas;
            var el = this.el = this.layers.rect(x, y, 0, 0);
            el.attr({
                "fill": "transparent",
                "stroke-width": this.getFixSize(pathWidth),
                "stroke": canvas.get("foreground")
            })
        },
        setAttr: function(x, y, x2, y2) {
            var width = x2 - x;
            var height = y2 - y;
            var el = this.el;
            if(width<0){
                width *= -1;
                x -= width;
            }
            if(height< 0){
                height *= -1;
                y -= height;
            }
            el.attr('x', x);
            el.attr('y', y);
            el.attr('width', width);
            el.attr('height', height);
        },
        fixSize: function(){
            this.el && this.el.attr('stroke-width', this.getFixSize(pathWidth))
        },
        drawRect: function(){
            this.addShape(this.el);
        },
        destroy: function(){
            if(this.isDrawing){
                this.el.remove()
            }
            $(this.canvas.main).off('.Rect');
        }
    })

    Rect.draw = function(data, svg, zoom){
        var ele = data.data;
        var el = svg.rect(ele[0], ele[1], ele[2], ele[3]);
        el.attr({
            "fill": "transparent",
            "stroke-width": util.fixSize(pathWidth, zoom),
            "stroke": data.color || "#000"
        })
        return el;
    }

    Rect.getData = function(data){
        var el = data.el;
        var rect = el;
        var trans = Raphael.parseTransformString(rect.transform())[0] || []
        var tx = trans[1] || 0;
        var ty = trans[2] || 0;
        return [rect.attr('x') + tx, rect.attr('y') + ty, rect.attr('width'), rect.attr('height')]
        return {
            x: rect.attr('x'),
            y: rect.attr('y'),
            width: rect.attr('width'),
            height: rect.attr('height')
        }
    }

    Rect.fixSize = function(item, zoom){
        item.el.attr({
            'stroke-width': util.fixSize(pathWidth, zoom)
        })
    }

    module.exports = Rect;

/***/ },
/* 5 */
/***/ function(module, exports) {

    function realSize(size, zoom){
        return Math.max(Math.floor(size / zoom), 0);
    }

    function fixSize(size, zoom){
        return Math.max(size / zoom, 1);
    }

    var uniqueId = (function(){
        var types = {};
        return function(type){
            type = type || "prefix"
            if(typeof types[type] == "undefined"){
                types[type] = 0;
            }
            types[type] += 1;
            return type + types[type];
        }
    })()

    module.exports = {
        uniqueId: uniqueId,
        realSize: realSize,
        fixSize: fixSize
    }

/***/ },
/* 6 */
/***/ function(module, exports, __webpack_require__) {

    var klass = __webpack_require__(2);

    var util = __webpack_require__(5);

    var types = {};

    var ParentPlugin = klass({

        initialize: function(parent, options){
            this.parent = parent;
            this.canvas = parent;
            this.layers = this.canvas.layers;

            this.options = $.extend(true, {}, options || {});

            this.container = this.canvas.main[0];
            this.initProp();
            this.bindEvent();
        },

        initProp: function(){

        },

        bindEvent: function(){

        },

        stop: function(){

        },

        addShape: function(el){
            var data = $.extend(true, {uid: util.uniqueId(this.type+"_")}, this.options, {el: el, datatype: this.type});
            this.parent.add(data);

            this.el = null;
        },

        getZoomSize: function(size){
            var zoom = this.canvas.zoom || 1;
            return util.realSize(size, zoom)
        },

        getFixSize: function(size){
            var zoom = this.canvas.zoom || 1;
            return util.fixSize(size, zoom)
        },

        getRealOffset: function(e){
            // 最外层的位置
            var mainOffset = $(this.canvas.main).offset();
            // 鼠标点击在页面上的位置
            var pageSet = {
                x: e.pageX,
                y: e.pageY
            };
            // // canvas svg的位置，padding
            // var canvasOffset = $(this.canvas.canvasInner).position();
            // // canvas的滚动位置，scrollLeft, scrollTop
            // var canvas2Offset = $(this.canvas.canvas).position();

            // var offsetX = pageSet.x - mainOffset.left - canvasOffset.left - canvas2Offset.left;
            // var offsetY = pageSet.y - mainOffset.top - canvasOffset.top - canvas2Offset.top;
            // 
            var offsetX = pageSet.x - mainOffset.left;
            var offsetY = pageSet.y - mainOffset.top;

            // console.log(mainOffset, pageSet, canvasOffset,canvas2Offset,  offsetX, offsetY)
            if(offsetX < 0){
                offsetX = Math.min(0, offsetX);
            }

            if(offsetY < 0){
                offsetY = Math.min(0, offsetY);
            }

            return {
                offsetX: offsetX,
                offsetY: offsetY
            }
        },

        destroy: function(){
            this.stop();
            $(this.canvas.main).off('.' + this.type);
        }
    })

    module.exports = ParentPlugin;

/***/ }
/******/ ]);