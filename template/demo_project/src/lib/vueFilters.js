/**
 * DetectionFaceType = 101
 * DetectionCarType  = 201
 */
Vue.filter('parsetype', function(val){
    var map = {
        201 : "卡口/电警",
        101 : "海关/车站"
    }
    return map[val] || "未知";
})