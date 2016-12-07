module.exports = {
 entry: "./{{component}}.js",
  output: {
    path: './',
    filename: "{{component}}_test.js"
  },

  module:  {
        loaders: [
            {
                test: /\.js$/,
                loader: "babel-loader",
                exclude: /(node_modules|bower_components)/
            },
            {
                test: /\.vue$/,
                loader: "vue-loader"
            }
        ]
    }
} 