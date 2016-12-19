module.exports = {
  entry: "./navbar.js",
  output: {
    path: './',
    filename: "navbar_test.js"
  },

  module: {
    loaders: [
      { test: /\.vue$/, loader: "vue-loader" }
    ]
  }
}