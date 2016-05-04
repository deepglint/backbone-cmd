require("./board.js");

var InterestingArea = function(options){
    this.initialize(options);
}
InterestingArea.prototype = {
    initialize: function(options){
        this.board = new Board(options);
        this.bindEvents()
        this.area = [];
    },
    resize: function(width, height, zoom){
        this.clearAll()
        this.board.update(width, height);
        this.board.zoomTo(zoom);
        this.board.clearAll()
    },
    draw: function(){
        this.board.switchTo('rect');
    },
    disabledDraw: function () {
        this.board.removePlugin()
    },
    clearAll: function(){
        var mask = this.mask;
        mask && mask.remove();
        this.board.clearAll()
        this.area = []
    },
    __drawMask: function(item){
        var board = this.board;
        var rect = item.el;
        var x = rect.attr("x");
        var y = rect.attr("y");
        var width = rect.attr("width");
        var height = rect.attr("height");
        var bWidth = board.width;
        var bHeight = board.height;

        var mask = this.mask;

        mask && mask.remove();
        mask = board.layers.add([{
            type: "rect",
            x: 0,
            y: 0,
            width: x,
            height: bHeight
        }, {
            type: "rect",
            x: x + width,
            y: 0,
            width: bWidth - x - width,
            height: bHeight
        }, {
            type: "rect",
            x: x,
            y: 0,
            width: width,
            height: y
        }, {
            type: "rect",
            x: x,
            y: y + height,
            width: width,
            height: bHeight - y - height
        }]);

        mask.attr({
            "fill": "rgba(0,0,0,0.3)",
            "stroke-width": 0
        })

        this.mask = mask;

        this.area = [x, y, width, height];
    },
    bindEvents: function(){
        var self = this;
        var board = this.board;
        
        $(board).on('board:add', function(e, item) {
            self.__drawMask(item);
        })

        $(board).on('board.drawstart', function(e) {
            self.clearAll()
        })
    }
}


module.exports = InterestingArea;