module.exports = {
  configureWebpack: {
    watch: true,
    watchOptions: {
      poll: true,
      ignored: /node_modules/,
    },
  },
};
