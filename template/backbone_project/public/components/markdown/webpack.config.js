module.exports = {
  entry: "./markdown.js",
  output: {
    path: './',
    filename: "markdown_test.js"
  },

  module: {
    loaders: [
      { test: /\.vue$/, loader: "vue-loader" }
    ]
  }
}