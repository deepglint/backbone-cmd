module.exports = {
  entry: "./sidebar.js",
  output: {
    path: './',
    filename: "sidebar_test.js"
  },

  module: {
    loaders: [
      { test: /\.vue$/, loader: "vue-loader" }
    ]
  }
}