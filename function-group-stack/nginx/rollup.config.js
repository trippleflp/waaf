import nodeResolve from "@rollup/plugin-node-resolve";
import commonjs from "@rollup/plugin-commonjs";
import glob from "glob";
import { getBabelOutputPlugin } from "@rollup/plugin-babel";

const filesToBundle = glob.sync("./js/*");

const defaultExporter = () => {
  return {
    name: "exportDefaults", // this name will show up in warnings and errors
    renderChunk: async (code, chunk) =>
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
    file: file.replace("./js/", "./lib/"),
    format: "es",
  },
  plugins: [
    nodeResolve({
      mainFields: ["jsnext:main", "module"],
      moduleDirectories: ["node_modules"],
    }),
    commonjs(),
    getBabelOutputPlugin({
      presets: ["@babel/preset-env"],
      plugins: [],
    }),
    defaultExporter(),
  ],
}));
