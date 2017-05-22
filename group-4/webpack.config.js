var webpack = require("webpack")
var path = require("path")
 
module.exports = {
  resolve: {
    extensions: ['.js']
  },
  entry: __dirname + "/assets/js/app.js",
  output: {
    path: __dirname + "/assets/js",
    filename: 'bundle.js',
    libraryTarget: "umd"
  },
  module: {
    loaders: [
      {
        test: /\.js?$/,
        loader: "babel-loader",
        exclude: "/node_modules/",
        query: {
          presets: ["es2015"]
        }
      }
    ]
  },
  plugins: [
    new webpack.DefinePlugin({ "global.GENTLY": false }),
  ]
};