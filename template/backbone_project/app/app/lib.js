// jquery
require('jquery');
// bootstrap
// require('lib/bootstrap.min.js');
require('bower_components/bootstrap/dist/js/bootstrap.js');
// swf create
// window.swfobject = require('../../public/lib/swfobject.js');
// underscore
// window._ = require('bower_components/underscore/underscore-min.js');
// moment
// window.moment = require('moment');

// moment 扩展，包含moment
window.moment = require('bower_components/moment-timezone/builds/moment-timezone-with-data.js');

require('bower_components/eonasdan-bootstrap-datetimepicker/src/js/bootstrap-datetimepicker.js');

window.bootbox = require('bootbox');
bootbox.setDefaults("locale", "zh_CN");
bootbox.setDefaults("backdrop", false);