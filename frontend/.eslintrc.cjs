module.exports = {
  parser: '@typescript-eslint/parser',
  extends: ['eslint:recommended', 'plugin:@typescript-eslint/recommended', 'plugin:@typescript-eslint/recommended-requiring-type-checking'],
  parserOptions: {
    ecmaVersion: 2020,
    sourceType: 'module',
    tsconfigRootDir: __dirname,
    project: ['./tsconfig.json'],
    extraFileExtensions: ['.svelte'],
  },
  env: {
    es6: true,
    browser: true,
  },
  overrides: [
    {
      files: ['*.svelte'],
      parser: 'svelte-eslint-parser',
    },
  ],
  settings: {
    'svelte3/typescript': require('typescript'),
    'svelte3/ignore-styles': () => true,
  },
  plugins: ['svelte3', '@typescript-eslint', 'prettier'],
  ignorePatterns: ['node_modules'],
}
