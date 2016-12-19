module.exports = {
  entry: "./changePwd.js",
  output: {
    path: './',
    filename: "changePwd_test.js"
  },

  module: {
    loaders: [
      { test: /\.vue$/, loader: "vue-loader" }
    ]
  }
}