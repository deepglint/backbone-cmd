module.exports = {
  entry: "./helloworld.js",
  output: {
    path: './',
    filename: "helloworld_test.js"
  },

  module: {
    loaders: [
      { test: /\.vue$/, loader: "vue-loader" }
    ]
  }
}