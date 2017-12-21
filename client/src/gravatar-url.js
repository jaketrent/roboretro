// adapted from https://github.com/sindresorhus/gravatar-url/blob/master/index.js

const crypto = require('crypto')

const BASE_URL = 'https://gravatar.com/avatar/'

module.exports = (email, opts) =>
  BASE_URL +
  crypto
    .createHash('md5')
    .update(email.toLowerCase().trim(), 'utf8')
    .digest('hex')
