module.exports = {
  root: true,
  env: {
    node: true
  },
  'extends': [
    'plugin:vue/strongly-recommended',
    'eslint:recommended'
  ],
  rules: {
    'no-console': 'off',
    'no-debugger': 'off',
    'vue/max-attributes-per-line': [2, {
      "singleline": 5,
      "multiline": {
        "max": 1,
        "allowFirstLine": false
      }
    }]
  },

   parserOptions: {
    parser: 'babel-eslint'
  }
}
