module.exports = {
  entry: "./hello.js",
  output: {
    path: './',
    filename: "hello_test.js"
  },

  module: {
    loaders: [
      { test: /\.vue$/, loader: "vue-loader" }
    ]
  }
}