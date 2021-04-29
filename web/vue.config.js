module.exports = {
  publicPath: '/',
  productionSourceMap: false,
  configureWebpack: {
    module: {
      rules: [
        {
          test: /^config\.js$/,
          use: [
            {
              loader: 'file-loader',
              options: {
                name: 'config.js',
              },
            },
          ],
        },
      ],
    },
  },
};
