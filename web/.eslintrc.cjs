const path = require('path');
const createAliasSetting = require('@vue/eslint-config-airbnb/createAliasSetting');

module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: [
    'plugin:vue/vue3-essential',
    '@vue/airbnb',
  ],
  parserOptions: {
    parser: '@babel/eslint-parser',
    requireConfigFile: false,
  },
  settings: {
    ...createAliasSetting({ '@': path.resolve(__dirname, 'src') }),
    'import/extensions': ['.js', '.jsx', '.mjs', '.vue'],
    'import/core-modules': ['pinia'],
  },
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    // Buefy component props that look like ARIA attributes to the plugin
    'vuejs-accessibility/aria-props': 'off',
    'vuejs-accessibility/click-events-have-key-events': 'off',
    'vuejs-accessibility/anchor-has-content': 'off',
    'vue/require-default-prop': 'off',
    'vue/multi-word-component-names': 'off',
    'vue/no-mutating-props': 'off',
  },
};
