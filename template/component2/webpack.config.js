module.exports = {
 entry: "./{{component}}.js",
  output: {
    path: './',
    filename: "{{component}}_test.js"
  },

  devtool: 'inline-source-map',

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
            },
            {
                test: /\.(png|jpg)$/,
　　　　　　      loader: "url-loader?limit=100000&name=build/[hash:8].[name].[ext]"
            }
        ]
    }
} 