var merge = require('webpack-merge')
var prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  // STORE_ENV: '"mock"'
  STORE_ENV: '"production"'
})
