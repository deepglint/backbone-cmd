var path = require('path');

module.exports = {
    entry: {
        app: "./app.js",
        login: "./login.js",
        lib: "./lib.js"
    },

    output: {
        path: '../../public/js/',
        filename: "[name].js"
    },

    resolve: {
        alias: {
            "app": path.resolve('../../', './app/app'),
            "lib": path.resolve('../../', './public/lib'),
            "components": path.resolve('../../', './public/components'),
            "bower_components": path.resolve('../../', './public/bower_components'),
            // for amd; bower_components/eonasdan-bootstrap-datetimepicker/build/js/bootstrap-datetimepicker.min.js
            "moment": "bower_components/moment/moment.js"
        }
    },

    module: {
        loaders: [{
            test: /\.vue$/,
            loader: "vue-loader"
        }]
    }
};