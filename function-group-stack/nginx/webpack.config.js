const path = require("path");
const { library } = require("webpack");

module.exports = {
  experiments: {
    outputModule: true,
  },
  entry: "./src/index.ts",
  mode: "production",
  //   target: ["web", "es5"],
  output: {
    // module: true,
    library: {
      type: "module",
    },
    filename: "index.js",
    path: path.resolve(__dirname, "lib"),
  },
  optimization: {
    minimize: false,
  },
  resolve: {
    extensions: [".tsx", ".ts", ".js"],
  },
  node: {
    global: true,
  },
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        use: "ts-loader",
        exclude: /node_modules/,
      },
      // {
      //   test: /\.m?js$$/,
      //   exclude: /(bower_components)/,
      //   use: {
      //     loader: "babel-loader",
      //     options: {
      //       // plugins: ["@babel/plugin-transform-runtime"],
      //       presets: ["@babel/preset-env"],
      //     },
      //   },
      // },
    ],
  },
};
