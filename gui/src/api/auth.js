import axios from 'axios'
import cookies from 'browser-cookies'

// login uses credentials to login an account.
let login = ({credentials, successCallback, errorCallback}) => {
  // Deal with credentials here ...
  // ...

  // Send request
  axios.post('/mogapis/v1/auth/login', {
    data: credentials
  }, {
    headers: {
      'Authorization': 'Bearer ' + cookies.get('at')
    }
  })
  .then((response) => {
    if (typeof successCallback === 'function') {
      successCallback(response.data.data)
    }
  })
  .catch((e) => {
    if (typeof errorCallback === 'function') {
      if (typeof e.response === 'undefined' || e.response.status === 504) {
        return errorCallback({
          status: 500,
          error: {
            code: 'ServerError',
            message: 'Can not connect to server.'
          }
        })
      }
      errorCallback({
        status: e.response.status,
        error: e.response.data.error
      })
    }
  })
}

// fetchAuthInfo gets authentication information of an authenticated account.
let fetchAuthInfo = ({successCallback, errorCallback}) => {
  // Send request
  axios.get('/mogapis/v1/auth/info', {
    headers: {
      'Authorization': 'Bearer ' + cookies.get('at')
    }
  })
  .then((response) => {
    if (typeof successCallback === 'function') {
      successCallback(response.data.data)
    }
  })
  .catch((e) => {
    if (typeof errorCallback === 'function') {
      if (typeof e.response === 'undefined' || e.response.status === 504) {
        return errorCallback({
          status: 500,
          error: {
            code: 'ServerError',
            message: 'Can not connect to server.'
          }
        })
      }
      errorCallback({
        status: e.response.status,
        error: e.response.data.error
      })
    }
  })
}

export default {
  login,
  fetchAuthInfo
}
