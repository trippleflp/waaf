import nodeResolve from "@rollup/plugin-node-resolve";
import commonjs from "@rollup/plugin-commonjs";
import glob from "glob";
import { babel, getBabelOutputPlugin } from "@rollup/plugin-babel";
import nodePolyfills from "rollup-plugin-node-polyfills";
import globals from "rollup-plugin-node-globals";
import regexUnicode from "@babel/plugin-transform-unicode-regex";

const filesToBundle = glob.sync("./js/*");

const myExample = () => {
  return {
    name: "exportDefaults", // this name will show up in warnings and errors
    renderChunk: async (code, chunk) =>
      //   `${code.replace("export", "export default")}`,
      `${code
        .replace(`export { ${chunk.exports.join(", ")} };`, "")
        .replace(
          `export{${chunk.exports.join(",")}}`,
          ""
        )}export default { ${chunk.exports.join(", ")} };`,
  };
};
export default filesToBundle.map((file) => ({
  input: file,
  output: {
    // export: "named",
    file: file.replace("./js/", "./lib/"),
    format: "es",
  },
  plugins: [
    // globals(),
    // nodePolyfills(),
    nodeResolve({
      //   preferBuiltins: true,
      mainFields: ["jsnext:main", "module"],
      moduleDirectories: ["node_modules"],
    }),
    commonjs(),
    // babel(),
    getBabelOutputPlugin({
      presets: ["@babel/preset-env"],
      plugins: [
        // "@babel/plugin-transform-unicode-regex",
        // "@babel/plugin-transform-sticky-regex",
      ],
    }),
    myExample(),
  ],
}));

// export default {
//   plugins: [exportDefaults()],
// };
