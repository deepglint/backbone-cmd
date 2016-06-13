var path = require('path');

module.exports = {
    entry: {
        app: "./app/app/app.js",
        login: "./app/app/login.js",
        lib: "./app/app/lib.js"
    },

    output: {
        path: path.resolve('./public/js'),
        filename: "[name].js"
    },

    resolve: {
        alias: {
            "app": path.resolve('./app/app'),
            "lib": path.resolve('./public/lib'),
            "components": path.resolve('./public/components'),
            "bower_components": path.resolve('./public/bower_components'),
            "jquery": "bower_components/jquery/jquery.js",
            // for amd; bower_components/eonasdan-bootstrap-datetimepicker/build/js/bootstrap-datetimepicker.min.js
            "moment": "bower_components/moment/moment.js",
            "randomColor": "bower_components/randomColor/randomColor.js"
        }
    },

    module: {
        loaders: [{
            test: /\.vue$/,
            loader: "vue-loader"
        }]
    }
};