module.exports = {
  entry: "./login.js",
  output: {
    path: './',
    filename: "login_test.js"
  },

  module: {
    loaders: [
      { test: /\.vue$/, loader: "vue-loader" }
    ]
  }
}