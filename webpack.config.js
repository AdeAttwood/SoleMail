const c = require("@practically/webpack-config");
const path = require("path");

c.initialize({
  src_path: path.resolve(__dirname, "frontend"),
  dest_path: path.resolve(__dirname, "build", "dist"),
  entry_point: path.resolve(__dirname, "frontend", "index.ts"),
});

c.html(path.resolve(__dirname, "frontend", "index.html"));
c.styles();
c.typescript();

const webpackConfig = c.build();

webpackConfig.resolve.alias["@app"] = path.resolve(__dirname, "frontend");

module.exports = webpackConfig;
